// The WBC developers. Copyright (c) 2017 
//

package legacyrpc

import "context"

type contextKey string

func withRemoteAddr(parent context.Context, remoteAddr string) context.Context {
	return context.WithValue(parent, contextKey("remote-addr"), remoteAddr)
}

func remoteAddr(ctx context.Context) string {
	v := ctx.Value(contextKey("remote-addr"))
	if v == nil {
		return "<unknown>"
	}
	return v.(string)
}
