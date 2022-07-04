package subnet

import (
	"fmt"
	"strings"
)

func loopTreeAndPrint(s *subnetNode, level int) {
	if s == nil {
		return
	}

	switch level {
	case 0:
		//  全树打印
		printNode(s)

	case 1:
		//  只打印未分配
		if !s.isAllocated || s.IsRoot {
			printNode(s)
		}
	case 2:
		//  只打印未分配 且 merge subnet
		//  只打印未分配
		if !s.isAllocated || s.IsRoot {
			printNode(s)
		}
		if !s.isAllocated {
			//  上层 未分配，下层也一定未分配，不必打印 直接返回
			return
		}

	default:
		printNode(s)

	}

	if s.left != nil {
		loopTreeAndPrint(s.left, level)
	}
	if s.right != nil {
		loopTreeAndPrint(s.right, level)
	}

}
func printNode(s *subnetNode) {
	space := []string{}
	for i := 0; i < s.level-1; i++ {
		if i == s.level-2 {
			space = append(space, "| _")
			continue

		}
		space = append(space, "| ")
	}
	space = append(space, fmt.Sprintf("%s %v", s.Net.String(), s.isAllocated))

	if s.Net != nil {
		fmt.Printf("%s\n", strings.Join(space, ""))
	}
}
