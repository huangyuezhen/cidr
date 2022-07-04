package subnet

import (
	"net"
)

var subnets = []string{
	"10.253.0.0/18",
	"10.253.64.0/19",
	"10.253.144.0/20",
	"10.253.96.0/20",
	"10.253.160.0/20",
	"10.253.128.0/21",
	"10.253.136.0/21",
	"10.253.240.0/20",
	"10.253.224.0/20",
	"10.253.208.0/21",
}
var vpcs = map[string]string{
	"10.253.0.0/18":   "vpc-uf6vqld3a289b0hln8vk6",
	"10.253.64.0/19":  "vpc-01c3aa587cc3f0d69",
	"10.253.144.0/20": "vpc-0ebc956a90304be53",
	"10.253.96.0/20":  "vpc-0f53ce842db1c14ed",
	"10.253.160.0/20": "vpc-0c8745d360e04e642",
	"10.253.128.0/21": "vpc-0072f2241e2fd986a",
	"10.253.136.0/21": "vpc-098a351b5dff5da66",
	"10.253.224.0/20": "projects/bytedance-game-project/global/networks/vpc-253",
	"10.253.208.0/21": "projects/bytedance-game-project/global/networks/vpc-253",
	"10.253.240.0/20": "vpc-0072f2241e2fd986a",
}
var pSubnet string = "10.253.0.0/16"

func markForRoot() {
	_, pNet, err := net.ParseCIDR(pSubnet)
	if err != nil {
		panic(err)
	}

	ss := build(pNet, nil)
	ss.MarkAllocated(subnets)

	ss.loop()

	loopTreeAndPrint(ss, 2)
}
func testBuild() {
	markForRoot()
	// loopVPCs()
}
