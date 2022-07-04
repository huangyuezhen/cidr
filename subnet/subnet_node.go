package subnet

import (
	"fmt"
	"net"
)

type subnetNode struct {
	IsRoot      bool
	Net         *net.IPNet
	left        *subnetNode
	right       *subnetNode
	parent      *subnetNode
	isAllocated bool
	level       int
	IPRange     string
}

func build(srcNet *net.IPNet, parent *subnetNode) *subnetNode {
	ones, _ := srcNet.Mask.Size()
	if ones > (27) {
		return &subnetNode{}

	}
	s := &subnetNode{
		isAllocated: false,
		Net:         srcNet,
	}
	ip_start, ip_end := IPAddrRange(srcNet)
	s.IPRange = fmt.Sprintf("%s~%s", ip_start, ip_end)
	if parent == nil {

		s.parent = nil
		s.IsRoot = true
		s.level = 1
	} else {
		s.parent = parent
		s.IsRoot = false
		s.level = parent.level + 1
	}

	left, right := SplitSubnet(srcNet)

	//left child
	s.left = build(left, s)

	// right child
	s.right = build(right, s)

	return s

}

var result []*subnetNode

//  二叉树前序遍历，根结点--》左子树--》右子树
func (s *subnetNode) loop() {
	if s.Net == nil {
		return

	}
	result = append(result, s)

	if s.left != nil {
		s.left.loop()
	}
	if s.right != nil {
		s.right.loop()
	}

}
func (s *subnetNode) MarkAllocated(cidrs []string) {
	for _, cidr := range cidrs {
		_, pNet, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(err)
		}
		s.markAllocated(pNet)

	}

}

// 对一颗已构建好的二叉子网树，根据传入的已使用子网，标记二叉树 的子网使用情况
//  若一个子网已被分配， 则其父+子节点 都会被标记为 已使用
func (s *subnetNode) markAllocated(ipNet *net.IPNet) {

	if !s.Net.Contains(ipNet.IP) {
		// 不属于范围内 跳过
		return
	}
	pSize, _ := s.Net.Mask.Size()
	cSize, _ := ipNet.Mask.Size()

	if pSize == cSize {
		//  标记本身+child+parent 为已分配
		s.markChildAllocated()
		s.markParentAllocated()
		return
	}

	if pSize > cSize {
		// 掩码小于 传入的值, 非包含关系 直接返回
		return
	}

	if s.left != nil && s.left.Net != nil && s.left.Net.Contains(ipNet.IP) {
		//  递归标记
		s.left.markAllocated(ipNet)
	} else if s.right != nil && s.right.Net != nil {
		s.right.markAllocated(ipNet)
	}

}
func (s *subnetNode) markChildAllocated() {
	s.isAllocated = true
	if s.right != nil {
		s.right.markChildAllocated()
	}
	if s.left != nil {
		s.left.markChildAllocated()
	}

}
func (s *subnetNode) markParentAllocated() {
	s.isAllocated = true
	if s.parent != nil {
		s.parent.markParentAllocated()
	}
}
