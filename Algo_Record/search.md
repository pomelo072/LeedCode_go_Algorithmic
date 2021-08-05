# 查找算法

---
- [二分查找](#二分查找)
---

## 二分查找

- [二分查找](https://leetcode-cn.com/problems/binary-search/) `easy`
- [第一个错误的版本](https://leetcode-cn.com/problems/first-bad-version/) `easy`
- [搜索插入位置](https://leetcode-cn.com/problems/search-insert-position/) `easy`


#### 解题思路

基础思想为二分查找, 时间复杂度 O(log n)

通过对目标和中心元素比较, 判断目标元素位置
- 相等, 则中心元素为目标元素
- 小于, 则目标元素在此中心元素左侧
- 大于, 则目标元素在此中心元素右侧

循环对中心元素判断减小区间范围, 直至等于某个中心元素或区间内只剩一个元素.

#### 解题代码

- [二分查找](../Binary_Search/code.go)
- [第一个错误的版本](../First_Bad_Version/code.go)
- [搜索插入位置](../Search_Insert_Position/code.go)

#### Tips

- 注意判断边界位置
- 在区间变换时注意不同语言的开合条件