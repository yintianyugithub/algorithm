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
	Val    int
	Next   *ListNode
	Random *ListNode
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

// IsCycle 判断链表是否有环
func IsCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil || head.Next.Next == nil {
		return false
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}

	return false
}

// TowNumAdd 链表两数相加
func TowNumAdd(l1, l2 *ListNode) *ListNode {
	f := 0
	dummy := &ListNode{}
	pre := dummy
	for l1 != nil || l2 != nil || f > 0 {
		if l1 != nil {
			f += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			f += l2.Val
			l2 = l2.Next
		}

		pre.Next = &ListNode{Val: f % 10}
		f /= 10

		pre = pre.Next
	}

	return dummy.Next
}

// RandomCopy 随机链表的复制
func RandomCopy(root *ListNode) *ListNode {
	for cur := root; cur != nil; cur = cur.Next.Next {
		cur.Next = &ListNode{
			Val:  cur.Val,
			Next: cur.Next,
		}
	}

	for cur := root; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}

	dummy := &ListNode{}
	pre := dummy
	for cur := root; cur != nil; cur, pre = cur.Next, pre.Next {
		pre.Next = cur.Next
		cur.Next = pre.Next.Next
	}

	return dummy.Next
}

// DeleteListN 删除链表倒数第N个节点
func DeleteListN(head *ListNode, num int) *ListNode {
	if head == nil || num == 0 {
		return nil
	}

	dummy := &ListNode{Next: head}
	l, r := dummy, dummy
	for range num {
		r = r.Next
	}
	for r.Next != nil {
		l, r = l.Next, r.Next
	}

	l.Next = r.Next.Next

	return dummy.Next
}

// ReserveList2 反转链表2
func ReserveList2(head *ListNode, l, r int) *ListNode {
	if head == nil || l > r {
		return nil
	}

	dummy := &ListNode{Next: head}
	tail := dummy

	for range l - 1 {
		tail = tail.Next
	}

	var pre, cur *ListNode = nil, tail.Next

	for range r - l + 1 {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	tail.Next.Next = cur
	tail.Next = pre

	return dummy.Next
}
