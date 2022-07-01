# golang_maximum_product_subarray

Given an integer array `nums`, find a contiguous non-empty subarray within the array that has the largest product, and return *the product*.

The test cases are generated so that the answer will fit in a **32-bit** integer.

A **subarray** is a contiguous subsequence of the array.

## Examples

**Example 1:**

```
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

```

**Example 2:**

```
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

```

**Constraints:**

- `1 <= nums.length <= 2 * 104`
- `10 <= nums[i] <= 10`
- The product of any prefix or suffix of `nums` is **guaranteed** to fit in a **32-bit** integer.

## 解析

給一個整數陣列 nums 

要求寫一個演算法找出一個子陣列使得所有元素的乘積數最大，回傳該數值

對每個元素可以發現

1. 如果所有元素都是大於 0 的狀況下

乘積最大值 res 就是 包含當下元素 nums[i] * 包含前 i-1 元素的乘積最大值

2 如果元素都小於 0 的狀況下

乘積最大值 res  因為有可能因為負數讓包含前 i-1 元素的乘積最大值 * nums[i] 變成乘積最小值

如果 包含前 i-1 元素的乘積最大值  > 0

乘積最大值 res 就是 包含當下元素 包含前 i-1 元素的乘積最大值

包含當下元素 nums[i] 乘積最大值 = nums[i]

包含當下元素 nums[i] 乘積最小值 = nums[i]*當下元素 nums[i] 乘積最大值 

如果 包含前 i-1 元素的乘積最大值  < 0

乘積最大值 res 就是 包含當下元素 nums[i] * 包含前 i-1 元素的乘積最小值

包含當下元素 nums[i] 乘積最大值 = 包含當下元素 nums[i] * 包含前 i-1 元素的乘積最小值

包含當下元素 nums[i] 乘積最小值 = nums[i] 

所以在有可能具有小於 0 的元素下, 必須透過 

計算包含該元素 nums[i] 的乘積最大值還有最小值 來跟 nums[i]本身做比較找出目前為止最大值

算式如下 

cur_max_include(i) = max(nums[i], cur_max_include(i-1)*nums[i], cur_min_include(i-1)*nums[i] )

cur_min_include(i) = max(nums[i], cur_max_include(i-1)*nums[i], cur_min_include(i-1)*nums[i] )

res = max(res, cur_max_include(i))

3 如果元素有 0 的狀況下

也是一樣的狀況

![](https://i.imgur.com/s2hwiJd.png)
## 程式碼
```go
package sol

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProduct(nums []int) int {
	res := nums[0]
	cur_max, cur_min := 1, 1

	for _, num := range nums {
		cur_max, cur_min = max(num, max(cur_max*num, cur_min*num)),
			min(num, min(cur_max*num, cur_min*num))
		res = max(res, cur_max)
	}
	return res
}

```
## 困難點

1. 要思考如何利用前 i-1 的計算結果來減少重複運算
2. 要看出乘積的遞迴關係並不直觀

## Solve Point

- [x]  先初始花 cur_max, cur_min = 1, 還有 res = nums[0]
- [x]  透過比較累成的結果比較有包含前 i-1 與沒包含大小 哪一個比較大