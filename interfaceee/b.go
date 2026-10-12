package interfaceee

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

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

// DeleteDuplication 删除链表重复元素
func DeleteDuplication(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	pre := dummy

	for pre.Next != nil && pre.Next.Next != nil {
		v := pre.Next.Val
		if v == pre.Next.Next.Val {
			for pre.Next != nil && pre.Next.Val == v {
				pre.Next = pre.Next.Next
			}
		} else {
			pre = pre.Next
		}
	}

	return dummy.Next
}

// rotateRight 旋转链表，把链表向右移动K个位置
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	n := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}

	if n%k == 0 {
		return head
	}

	dummy := head
	for range n - k - 1 {
		dummy = dummy.Next
	}

	res := dummy.Next
	dummy.Next = cur.Next
	cur.Next = head

	return res
}

// PartitionList 分隔链表 把所有小于X的节点放在比大于等于X节点的左侧：保持原始位置不变
func PartitionList(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	l, r := &ListNode{}, &ListNode{}
	dl, dr := l, r
	cur := head
	for cur != nil {
		if cur.Val >= x {
			dr.Next = cur
			dr = dr.Next
		} else {
			dl.Next = cur
			dr = dl.Next
		}

		cur = cur.Next
	}

	dl.Next = r.Next
	dr.Next = nil

	return l.Next
}

// splitWord 反转字符串中的单词，不允许用split
func splitWord(s string) string {
	if s == "" {
		return s
	}

	s = strings.TrimLeft(s, "#")
	s = strings.TrimRight(s, "#")

	worldSlice := make([]string, 0)
	start := 0

	for i, v := range s {
		if v == '#' {
			worldSlice = append(worldSlice, s[start:i])
			start = i + 1
		}
	}

	worldSlice = append(worldSlice, s[start:])

	res := ""
	for i := len(worldSlice) - 1; i >= 0; i-- {
		if worldSlice[i] == "" {
			continue
		}
		res = fmt.Sprintf("%s#%s", res, worldSlice[i])
	}

	res = strings.TrimLeft(res, "#")

	return res
}

// MostNumber 多数元素
func MostNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k, cnt := nums[0], 1
	for _, v := range nums[1:] {
		if cnt == 0 {
			cnt = v
		}

		if v == k {
			cnt++
		} else {
			cnt--
		}
	}

	return k
}

// summaryRanges 汇总区间
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	res := make([]string, 0, len(nums))
	start := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i]-1 {
			if start == nums[i-1] {
				res = append(res, strconv.Itoa(nums[i-1]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			start = nums[i]
		}
	}

	if start == nums[len(nums)-1] {
		res = append(res, strconv.Itoa(start))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", start, nums[len(nums)-1]))
	}

	return res
}

// insertInterval 插入区间
func insertInterval(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	for _, v := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < v[0] {
			merged = append(merged, v)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], v[1])
		}
	}

	return merged
}

// mergeInterval合并区间
func mergeInterval(intervals [][]int) [][]int {

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	merged := make([][]int, 0)
	for _, v := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < v[0] {
			merged = append(merged, v)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], v[1])
		}
	}

	return merged
}

// latest 最少的箭引爆气球数
func latest(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	sort.Slice(points, func(a, b int) bool {
		return points[a][1] < points[b][1]
	})

	ans, start := 0, math.MinInt64
	for _, v := range points {
		if v[0] > start {
			ans++
			start = v[1]
		}
	}

	return ans
}
