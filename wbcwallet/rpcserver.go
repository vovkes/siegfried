// Copyright (c) 2013-2015 The btcsuite developers
// The WBC developers. Copyright (c) 2017 
//

package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	xcontext "golang.org/x/net/context"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/wbcoin/wbcwallet/loader"
	"github.com/wbcoin/wbcwallet/rpc/legacyrpc"
	"github.com/wbcoin/wbcwallet/rpc/rpcserver"
	"github.com/wbcoin/wbc/certgen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

// openRPCKeyPair creates or loads the RPC TLS keypair specified by the
// application config.  This function respects the cfg.OneTimeTLSKey setting.
func openRPCKeyPair() (tls.Certificate, error) {
	// Check for existence of the TLS key file.  If one time TLS keys are
	// enabled but a key already exists, this function should error since
	// it's possible that a persistent certificate was copied to a remote
	// machine.  Otherwise, generate a new keypair when the key is missing.
	// When generating new persistent keys, overwriting an existing cert is
	// acceptable if the previous execution used a one time TLS key.
	// Otherwise, both the cert and key should be read from disk.  If the
	// cert is missing, the read error will occur in LoadX509KeyPair.
	_, e := os.Stat(cfg.RPCKey.Value)
	keyExists := !os.IsNotExist(e)
	switch {
	case cfg.OneTimeTLSKey && keyExists:
		err := fmt.Errorf("one time TLS keys are enabled, but TLS key "+
			"`%s` already exists", cfg.RPCKey)
		return tls.Certificate{}, err
	case cfg.OneTimeTLSKey:
		return generateRPCKeyPair(false)
	case !keyExists:
		return generateRPCKeyPair(true)
	default:
		return tls.LoadX509KeyPair(cfg.RPCCert.Value, cfg.RPCKey.Value)
	}
}

// generateRPCKeyPair generates a new RPC TLS keypair and writes the cert and
// possibly also the key in PEM format to the paths specified by the config.  If
// successful, the new keypair is returned.
func generateRPCKeyPair(writeKey bool) (tls.Certificate, error) {
	log.Infof("Generating TLS certificates...")

	// Create directories for cert and key files if they do not yet exist.
	certDir, _ := filepath.Split(cfg.RPCCert.Value)
	keyDir, _ := filepath.Split(cfg.RPCKey.Value)
	err := os.MkdirAll(certDir, 0700)
	if err != nil {
		return tls.Certificate{}, err
	}
	err = os.MkdirAll(keyDir, 0700)
	if err != nil {
		return tls.Certificate{}, err
	}

	// Generate cert pair.
	org := "dcrwallet autogenerated cert"
	validUntil := time.Now().Add(time.Hour * 24 * 365 * 10)
	cert, key, err := certgen.NewTLSCertPair(cfg.TLSCurve.Curve(), org,
		validUntil, nil)
	if err != nil {
		return tls.Certificate{}, err
	}
	keyPair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return tls.Certificate{}, err
	}

	// Write cert and (potentially) the key files.
	err = ioutil.WriteFile(cfg.RPCCert.Value, cert, 0600)
	if err != nil {
		return tls.Certificate{}, err
	}
	if writeKey {
		err = ioutil.WriteFile(cfg.RPCKey.Value, key, 0600)
		if err != nil {
			rmErr := os.Remove(cfg.RPCCert.Value)
			if rmErr != nil {
				log.Warnf("Cannot remove written certificates: %v",
					rmErr)
			}
			return tls.Certificate{}, err
		}
	}

	log.Info("Done generating TLS certificates")
	return keyPair, nil
}

func startRPCServers(walletLoader *loader.Loader) (*grpc.Server, *legacyrpc.Server, error) {
	var jsonrpcAddrNotifier jsonrpcListenerEventServer
	var grpcAddrNotifier grpcListenerEventServer
	if cfg.RPCListenerEvents {
		jsonrpcAddrNotifier = newJSONRPCListenerEventServer(outgoingPipeMessages)
		grpcAddrNotifier = newGRPCListenerEventServer(outgoingPipeMessages)
	}

	var (
		server       *grpc.Server
		legacyServer *legacyrpc.Server
		legacyListen = net.Listen
		keyPair      tls.Certificate
		err          error
	)
	if cfg.DisableServerTLS {
		log.Info("Server TLS is disabled.  Only legacy RPC may be used")
	} else {
		keyPair, err = openRPCKeyPair()
		if err != nil {
			return nil, nil, err
		}

		// Change the standard net.Listen function to the tls one.
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{keyPair},
			MinVersion:   tls.VersionTLS12,
			NextProtos:   []string{"h2"}, // HTTP/2 over TLS
		}
		legacyListen = func(net string, laddr string) (net.Listener, error) {
			return tls.Listen(net, laddr, tlsConfig)
		}

		if len(cfg.GRPCListeners) != 0 {
			listeners := makeListeners(cfg.GRPCListeners, net.Listen)
			if len(listeners) == 0 {
				err := errors.New("failed to create listeners for RPC server")
				return nil, nil, err
			}
			creds := credentials.NewServerTLSFromCert(&keyPair)
			server = grpc.NewServer(
				grpc.Creds(creds),
				grpc.StreamInterceptor(interceptStreaming),
				grpc.UnaryInterceptor(interceptUnary),
			)
			rpcserver.RegisterServices(server)
			rpcserver.StartWalletLoaderService(server, walletLoader, activeNet)
			rpcserver.StartTicketBuyerService(server, walletLoader, &cfg.tbCfg)
			rpcserver.StartAgendaService(server, activeNet.Params)
			rpcserver.StartDecodeMessageService(server, activeNet.Params)
			for _, lis := range listeners {
				lis := lis
				go func() {
					laddr := lis.Addr().String()
					grpcAddrNotifier.notify(laddr)
					log.Infof("gRPC server listening on %s", laddr)
					err := server.Serve(lis)
					log.Tracef("Finished serving gRPC: %v", err)
				}()
			}
		}
	}

	if cfg.Username == "" || cfg.Password == "" {
		log.Info("Legacy RPC server disabled (requires username and password)")
	} else if len(cfg.LegacyRPCListeners) != 0 {
		listeners := makeListeners(cfg.LegacyRPCListeners, legacyListen)
		if len(listeners) == 0 {
			err := errors.New("failed to create listeners for legacy RPC server")
			return nil, nil, err
		}
		opts := legacyrpc.Options{
			Username:            cfg.Username,
			Password:            cfg.Password,
			MaxPOSTClients:      cfg.LegacyRPCMaxClients,
			MaxWebsocketClients: cfg.LegacyRPCMaxWebsockets,
		}
		legacyServer = legacyrpc.NewServer(&opts, activeNet.Params, walletLoader, listeners)
		for _, lis := range listeners {
			jsonrpcAddrNotifier.notify(lis.Addr().String())
		}
	}

	// Error when neither the GRPC nor legacy RPC servers can be started.
	if server == nil && legacyServer == nil {
		return nil, nil, errors.New("no suitable RPC services can be started")
	}

	return server, legacyServer, nil
}

// serviceName returns the package.service segment from the full gRPC method
// name `/package.service/method`.
func serviceName(method string) string {
	// Slice off first /
	method = method[1:]
	// Keep everything before the next /
	return method[:strings.IndexRune(method, '/')]
}

func interceptStreaming(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	p, ok := peer.FromContext(ss.Context())
	if ok {
		grpcLog.Infof("Streaming method %s invoked by %s", info.FullMethod,
			p.Addr.String())
	}
	err := rpcserver.ServiceReady(serviceName(info.FullMethod))
	if err != nil {
		return err
	}
	err = handler(srv, ss)
	if err != nil && ok {
		grpcLog.Errorf("Streaming method %s invoked by %s errored: %v",
			info.FullMethod, p.Addr.String(), err)
	}
	return err
}

func interceptUnary(ctx xcontext.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	p, ok := peer.FromContext(ctx)
	if ok {
		grpcLog.Infof("Unary method %s invoked by %s", info.FullMethod,
			p.Addr.String())
	}
	err = rpcserver.ServiceReady(serviceName(info.FullMethod))
	if err != nil {
		return nil, err
	}
	resp, err = handler(ctx, req)
	if err != nil && ok {
		grpcLog.Errorf("Unary method %s invoked by %s errored: %v",
			info.FullMethod, p.Addr.String(), err)
	}
	return resp, err
}

type listenFunc func(net string, laddr string) (net.Listener, error)

// makeListeners splits the normalized listen addresses into IPv4 and IPv6
// addresses and creates new net.Listeners for each with the passed listen func.
// Invalid addresses are logged and skipped.
func makeListeners(normalizedListenAddrs []string, listen listenFunc) []net.Listener {
	ipv4Addrs := make([]string, 0, len(normalizedListenAddrs)*2)
	ipv6Addrs := make([]string, 0, len(normalizedListenAddrs)*2)
	for _, addr := range normalizedListenAddrs {
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			// Shouldn't happen due to already being normalized.
			log.Errorf("`%s` is not a normalized "+
				"listener address", addr)
			continue
		}

		// Empty host or host of * on plan9 is both IPv4 and IPv6.
		if host == "" || (host == "*" && runtime.GOOS == "plan9") {
			ipv4Addrs = append(ipv4Addrs, addr)
			ipv6Addrs = append(ipv6Addrs, addr)
			continue
		}

		// Remove the IPv6 zone from the host, if present.  The zone
		// prevents ParseIP from correctly parsing the IP address.
		// ResolveIPAddr is intentionally not used here due to the
		// possibility of leaking a DNS query over Tor if the host is a
		// hostname and not an IP address.
		zoneIndex := strings.Index(host, "%")
		if zoneIndex != -1 {
			host = host[:zoneIndex]
		}

		ip := net.ParseIP(host)
		switch {
		case ip == nil:
			log.Warnf("`%s` is not a valid IP address", host)
		case ip.To4() == nil:
			ipv6Addrs = append(ipv6Addrs, addr)
		default:
			ipv4Addrs = append(ipv4Addrs, addr)
		}
	}
	listeners := make([]net.Listener, 0, len(ipv6Addrs)+len(ipv4Addrs))
	for _, addr := range ipv4Addrs {
		listener, err := listen("tcp4", addr)
		if err != nil {
			log.Warnf("Can't listen on %s: %v", addr, err)
			continue
		}
		listeners = append(listeners, listener)
	}
	for _, addr := range ipv6Addrs {
		listener, err := listen("tcp6", addr)
		if err != nil {
			log.Warnf("Can't listen on %s: %v", addr, err)
			continue
		}
		listeners = append(listeners, listener)
	}
	return listeners
}
