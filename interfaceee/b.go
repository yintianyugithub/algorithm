package interfaceee

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
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

// TowSum 两数之和
func TowSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	l, r := 0, len(nums)-1

	for l < r {
		v := nums[l] + nums[r]
		if v == target {
			return []int{l + 1, r + 1}
		} else if v < target {
			r--
		} else {
			l++
		}
	}

	return nil
}

// ThreeSum 三数之和
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			v := nums[l] + nums[r] + nums[i]
			if v == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if v > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

// MaxArea 盛最多水的容器
func MaxArea(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l, r := 0, len(nums)-1
	area := 0
	for l < r {
		if v := (r - l) * min(nums[l], nums[r]); v > area {
			area = v
		}
		if nums[l] < nums[r] {
			l++
		} else {
			r--
		}
	}

	return area
}

// IsHui 验证回文子串
func IsHui(s string) bool {
	tmp := make([]rune, 0, len(s))
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			tmp = append(tmp, unicode.ToLower(v))
		}
	}

	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		if tmp[i] != tmp[j] {
			return false
		}
	}

	return true
}

// SubStr 判断子序列
func SubStr(s, t string) bool {
	if s == "" {
		return true
	}

	i := 0
	for _, v := range t {
		if byte(v) == s[i] {
			i++
		}

		if i == len(s)-1 {
			return true
		}
	}

	return false
}

// MinSubArr 长度最小的子数组
func MinSubArr(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	sum, l, ml := 0, 0, len(nums)+1
	for r, v := range nums {
		sum += v
		for sum >= target {
			if r-l+1 < ml {
				ml = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}

	if ml == len(nums)+1 {
		return 0
	}

	return ml
}

// NoRepeatLongestStr 无重复最长子串
func NoRepeatLongestStr(s string) int {
	if len(s) == 0 {
		return 0
	}

	l, repeat, ml := 0, make(map[rune]int, len(s)), 0
	for r, v := range s {
		if idx, ok := repeat[v]; ok && idx >= l {
			l = idx + 1
		}
		if r-l+1 > ml {
			ml = r - l + 1
		}
		repeat[v] = r
	}

	return ml
}

// CoverWordsSubStr 覆盖所有单词的子串
func CoverWordsSubStr(s string, words []string) []int {
	if s == "" || len(words) == 0 {
		return nil
	}

	wordCnt := len(words)
	wordLen := len(words[0])
	need := make(map[string]int)
	for _, v := range words {
		need[v]++
	}

	res := make([]int, 0)
	for i := range wordLen {
		appearCnt := 0
		l := i
		win := make(map[string]int)
		for r := l; r+wordLen <= len(s); r = r + wordLen {
			word := s[r : r+wordLen]
			if needCnt, ok := need[word]; ok {
				appearCnt++
				win[word]++
				for win[word] > needCnt {
					lWord := s[l : l+wordLen]
					win[lWord]--
					appearCnt--
					l += wordLen
				}

				if appearCnt == wordCnt {
					res = append(res, l)
				}
			} else {
				win = make(map[string]int)
				appearCnt = 0
				l = r + wordLen
			}
		}
	}

	return res
}

// MinCoverSubStr 最小覆盖子串
func MinCoverSubStr(s, t string) string {
	if s == "" || t == "" {
		return ""
	}

	l, mLen, start, need, require := 0, len(s)+1, 0, make(map[rune]int), 0
	for _, v := range t {
		need[v]++
	}

	sn := make(map[rune]int)
	for r, v := range s {
		sn[v]++
		if need[v] > 0 && sn[v] == need[v] {
			require++

			for require == len(need) && l <= r {
				if mLen > r-l+1 {
					mLen = r - l + 1
					start = l
				}
				lv := rune(s[l])
				sn[lv]--
				if need[lv] > 0 && sn[lv] < need[lv] {
					require--
				}
				l++
			}
		}
	}

	if mLen == len(s)+1 {
		return ""
	}

	return s[start : start+mLen]
}

// LongestPalindrome 最长回文子串
func LongestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	ansL, ansR := 0, 0
	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > ansR-ansL {
			ansL, ansR = l+1, r
		}
	}

	return s[ansL:ansR]
}

// minimumTotal 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := range triangle[i] {
			triangle[i][j] += triangle[i+1][j] + triangle[i+1][j+1]
		}
	}

	return triangle[0][0]
}

// SetZeros 矩阵置零
func SetZeros(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	row, col := false, false

	// 首行
	for j := range n {
		if matrix[0][j] == 0 {
			row = true
			break
		}
	}

	// 首列
	for i := range m {
		if matrix[i][0] == 0 {
			col = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0 // 列
				matrix[i][0] = 0 // 行
			}
		}
	}
	// 改行
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// 改列
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	// 首行
	if row {
		for j := range n {
			matrix[0][j] = 0
		}
	}
	// 首列
	if col {
		for i := range m {
			matrix[i][0] = 0
		}
	}
}

// MaxSquare 最大正方形
func MaxSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	ans := 0
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := range m {
		for j := range n {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i][j], dp[i+1][j]) + 1
				ans = max(ans, dp[i+1][j+1])
			}
		}
	}

	return ans * ans
}

// MinPathSum 最小路径和
func MinPathSum(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	f, n := matrix[0], len(matrix[0])
	for i := 1; i < n; i++ {
		f[i] += f[i-1]
	}

	for i := 1; i < len(matrix); i++ {
		f[0] += matrix[i][0]
		for j := 1; j < n; j++ {
			f[j] = min(f[j-1], f[j]) + matrix[i][j]
		}
	}

	return f[n-1]
}

// Family 打家截舍
func Family(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

// 分发糖果
func candy(ratings []int) int {
	n := 0
	if len(ratings) == 0 {
		return n
	}

	res := make([]int, len(ratings))
	for i := range res {
		res[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			res[i] = max(res[i+1]+1, res[i])
		}
	}

	for _, v := range res {
		n += v
	}

	return n
}
