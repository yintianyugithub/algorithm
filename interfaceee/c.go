package interfaceee

import (
	"container/heap"
	"strconv"
	"strings"
)

// FindPeak 寻找峰值
func FindPeak(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

// SearchCycleArr 搜索旋转排序数组
// 先找左右两边哪边有序
// 判断target所在的范围，逐渐缩小一半
func SearchCycleArr(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		// 左边有序
		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

// SearchFirstEnd 在排序数组中查找第一个和最后一个位置
// 第一步找到大于等于target的位置
// 第二部找到大于target的位置
// 判断位置一是否等于，等于返回[位置一,位置二-1]
// 否则返回[-1,-1]
func SearchFirstEnd(nums []int, target int) []int {
	lower := func() int {
		l, r := 0, len(nums)
		for l < r {
			mid := l + (r-l)/2
			if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}

	up := func() int {
		l, r := 0, len(nums)
		for l < r {
			mid := l + (r-l)/2
			if nums[mid] <= target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}

	s := lower()
	if s == len(nums) || nums[s] != target {
		return []int{-1, -1}
	}

	return []int{s, up() - 1}
}

// FindMin 寻找旋转排序数组中的最小值
// 中间值和右边值对比：小于说明最小值在中间值左边，大于说明在中间值右边
func FindMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return nums[l]
}

// AddBinary 二进制字符串相加
// 和字符串整数逐位相加一样
func AddBinary(a, b string) string {
	i, j, carry := len(a)-1, len(b)-1, 0
	var buildStr strings.Builder
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		buildStr.WriteRune(rune('0' + sum%10))
		carry = sum / 10
	}

	res := []rune(buildStr.String())
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}

	return string(res)
}

// ReverseBit 颠倒二进制位
// 取数字的最后一位，res向左移动一位腾出位置(最后一位一定是0)
// 把取出的最后一位，放在res的最右边(0|1、0|0)
// 最后数字向右移动一位
func ReverseBit(n int) int {
	var res int

	for range 32 {
		res = (res << 1) | (n & 1)
		n >>= 1
	}

	return res
}

// hammingWeight 位1的个数
func hammingWeight(n int) int {
	s := strconv.FormatInt(int64(n), 2)
	ans := 0
	for _, v := range s {
		if v == 1+'0' {
			ans++
		}
	}
	return ans
}

// FindAppearFirstNumber 只出现一次的数字
func FindAppearFirstNumber(nums []int) int {
	var ans int

	for _, v := range nums {
		ans ^= v
	}

	return ans
}

// Oil 加油站
func Oil(nums []int, oils []int) int {
	rest, minRest, start := 0, 0, 0

	for i, v := range nums {
		rest += oils[i] - v
		if rest < minRest {
			minRest = rest
			start = i
		}
	}

	if rest >= 0 {
		return start
	}

	return -1
}

// FindAppearFirstNumber2 只出现一次的数字2
// 统计第i位是1的次数(出现三次的数字%3==0)、%3!=0的是出现一次的数字(把对应i位设置成1)
func FindAppearFirstNumber2(nums []int) int {
	var ans int

	for i := range strconv.IntSize {
		cnt := 0
		for _, v := range nums {
			if v>>i&1 == 1 {
				cnt++
			}
		}
		if cnt%3 != 0 {
			ans |= 1 << i
		}
	}

	return ans
}

// ExceptSelf 除自身外数组的乘积
// 先算i左边的数乘积
// 再算i右边数的乘积，左右相乘就是结果
func ExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1

	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= r
		r *= nums[i]
	}

	return res
}

// LongestConsecutive 最长连续序列
func LongestConsecutive(nums []int) int {
	// 先找到最小的数字
	// 从最小的数字累加
	// 长度大于从长度一半时停止
	if len(nums) == 0 {
		return 0
	}

	ans, m := 1, make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	for _, v := range nums {
		if m[v-1] {
			continue
		}

		increase := v + 1
		for m[increase] {
			increase++
		}

		ans = max(ans, increase-v)
		if ans*2 >= len(nums) {
			break
		}
	}

	return ans
}

type pair struct {
	i, j int
	sum  int
}

// MinHeap 小根堆(最小堆)堆顶最小 TopK问题：比对顶的值大才有资格进入堆
// 大根堆：堆顶最大，比堆顶小才有资格进入
type MinHeap []pair

func (m *MinHeap) Len() int {
	return len(*m)
}

func (m *MinHeap) Less(i, j int) bool {
	// 如果求TopK是不是换成 (*m)[i].sum < (*m)[j].sum 就可以了
	return (*m)[i].sum < (*m)[j].sum
}

func (m *MinHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(pair))
}

func (m *MinHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func FindKPairNumber(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return nil
	}

	minHap := &MinHeap{}
	for i, num := range nums1 {
		heap.Push(minHap, pair{
			i:   i,
			j:   0,
			sum: num + nums2[0],
		})
	}

	res := make([][]int, 0, k)
	for k > 0 && minHap.Len() > 0 {
		cur := heap.Pop(minHap).(pair)
		res = append(res, []int{nums1[cur.i], nums2[cur.j]})
		next := cur.j + 1
		if next < len(nums2) {
			heap.Push(minHap, pair{
				i:   cur.i,
				j:   next,
				sum: nums1[cur.i] + nums2[next],
			})
		}
		k--
	}

	return res
}

// SnakesAndLadders 蛇形梯
func SnakesAndLadders(board [][]int) int {
	n := len(board)
	target := n * n
	getPosition := func(s int) (int, int) {
		rowBottom := (s - 1) / n
		col := (s - 1) % n
		row := n - 1 - rowBottom
		if rowBottom%2 == 1 {
			col = n - 1 - col
		}

		return row, col
	}

	visited := make([]bool, target+1)
	visited[1] = true
	queue := []int{1}

	steps := 0

	for len(queue) > 0 {
		for _, cur := range queue {
			queue = queue[1:]
			if cur == target {
				return steps
			}
			for i := 1; i <= 6; i++ {
				next := cur + i
				if next > target {
					break
				}
				row, col := getPosition(next)
				if board[row][col] != -1 {
					next = board[row][col]
				}
				if visited[next] {
					continue
				}
				visited[next] = true
				queue = append(queue, next)
			}
		}

	}

	return -1
}
