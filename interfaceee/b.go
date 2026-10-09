package interfaceee

// sliceToBST 升序数组转成平衡二叉树
func sliceToBST(sorted []int) *TreeNode {
	if len(sorted) == 0 {
		return nil
	}

	var divide func(l, r int) *TreeNode
	divide = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := l + (r-l)/2
		root := &TreeNode{
			V: sorted[mid],
		}

		root.L = divide(l, mid-1)
		root.R = divide(mid+1, r)

		return root
	}

	return divide(0, len(sorted)-1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeList 合并两个有序链表
func MergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}

		pre = pre.Next
	}

	if l1 != nil {
		pre.Next = l1
	} else if l2 != nil {
		pre.Next = l2
	}

	return dummy.Next
}
