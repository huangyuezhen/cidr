package subnet

import (
	"math/big"
	"net"
)

var ipv4BitLen = 8
var ipv4BytesNum = 4

// SplitSubnet 增加一个mask 位，把网段一分为2， eg: src: 10.253.0.0/24  return 10.253.0.0/17 ,10.253.128.0/17
// ref: https://github.com/apparentlymart/go-cidr/blob/master/cidr/cidr.go
func SplitSubnet(src *net.IPNet) (*net.IPNet, *net.IPNet) {
	netLens, bits := src.Mask.Size()
	lowIP := src.IP
	highIP := highIP(lowIP, netLens, bits)

	return newNet(lowIP, netLens, bits), newNet(highIP, netLens, bits)
}

func highIP(lIP net.IP, netLens, bits int) net.IP {
	high := big.NewInt(1)
	high.Lsh(high, uint(bits-(netLens+1))) // 1<< hostBit, hostBit= bits- (netLens-1)

	low := &big.Int{}
	low.SetBytes(lIP)
	high.Or(low, high) // eg : low: 10.253.0.0  high: 0.0.128.0, low|high= 10.253.128.0

	return net.IP(high.FillBytes(make([]byte, ipv4BytesNum)))
}
func newNet(ip net.IP, netLens, bits int) *net.IPNet {
	return &net.IPNet{
		IP:   ip,
		Mask: net.CIDRMask(netLens+1, bits),
	}
}

func IPAddrRange(src *net.IPNet) (net.IP, net.IP) {
	start := src.IP
	netLens, bits := src.Mask.Size()

	endInt := big.NewInt(1) // set 1 then shift to proper position

	endInt.Lsh(endInt, uint(bits-netLens))     // nolint:gomnd // ipv4 prefix length is 32, (32-(ones+1))
	endInt = endInt.Sub(endInt, big.NewInt(1)) // eg: 10000-1 =1111

	startInt := &big.Int{}
	startInt.SetBytes(start)
	startInt.Or(startInt, endInt) // start ip int与 host 位都为1  取 或运算，eg: 10.253.0.0  | 0.0.255.255 -> 10.253.255.255

	return start, intToIP(startInt, bits)
}

func intToIP(ipInt *big.Int, bits int) net.IP {
	ipBytes := ipInt.Bytes()
	ret := make([]byte, bits/ipv4BitLen)
	// Pack our IP bytes into the end of the return array,
	// since big.Int.Bytes() removes front zero padding.
	for i := 1; i <= len(ipBytes); i++ {
		ret[len(ret)-i] = ipBytes[len(ipBytes)-i]
	}

	return net.IP(ret)
}
