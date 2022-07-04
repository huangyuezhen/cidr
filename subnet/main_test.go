package subnet

import (
	"fmt"
	"net"
	"testing"
)

func Test_testBuild(t *testing.T) {
	// testBuild()
	_, _net, _ := net.ParseCIDR("10.253.128.1/32")
	start, end := IPAddrRange(_net)
	fmt.Printf("%s~%s\n", start, end)
}
