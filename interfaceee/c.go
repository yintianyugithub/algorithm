package interfaceee

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
