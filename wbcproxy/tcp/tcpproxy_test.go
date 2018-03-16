package tcp

import "testing"

func TestTcpProxyYandex(t *testing.T){
	runTcpProxy("pop.yandex.ru", "110", "8080")
	// proxy is started after running this func on localhost

	// to check use telnet as example:

	// > telnet localhost 8080
	// > USER test
	// > PASS none

	// check log files
}