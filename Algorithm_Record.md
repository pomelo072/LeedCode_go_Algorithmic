# 算法记录

> 开始漫长的算法之旅, 这玩意怎么这么难的???
> 
> 好好练习吧

## 最长公共前缀 `easy`

- [最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/)

### 解题思路

- 通过对两两字符串之间比较找出最小的前缀子串, 直到子串为空或者遍历完整个string切片

### 解题代码

[Code](./Longest_Common_Prefix/code.go)

### Tips 

混淆了`:=`和`=`赋值符号对覆盖单个变量的问题

- `:=` 会分配新的内存, 产生一个全新的变量并赋值
- `=` 只是将值赋予变量

## 两数相加 `medium`

- [两数相加](https://leetcode-cn.com/problems/add-two-numbers/)

### 解题思路

- 对链表遍历相加节点
    1. 注意末尾进位, 把进位相加, 最后一位为进位时, 在插入一次节点
    2. 注意空链表
    3. 注意不等长链表, 空链表或者不等长链表都应该和一个值为0的节点相加, 知道全部节点遍历完成

### 解题代码

[Code](./Add_Two_Numbers/code.go)

### Tips

- 链表的遍历
- 确定参数返回值的返回参数

## 二分查找

- [二分查找](https://leetcode-cn.com/problems/binary-search/) `easy`
- [第一个错误的版本](https://leetcode-cn.com/problems/first-bad-version/) `easy`
- [搜索插入位置](https://leetcode-cn.com/problems/search-insert-position/) `easy`

### 解题思路

基础思想为二分查找, 时间复杂度 O(log n)
   
通过对目标和中心元素比较, 判断目标元素位置
- 相等, 则中心元素为目标元素
- 小于, 则目标元素在此中心元素左侧
- 大于, 则目标元素在此中心元素右侧 

循环对中心元素判断减小区间范围, 直至等于某个中心元素或区间内只剩一个元素.

### 解题代码

- [二分查找](./Binary_Search/code.go)
- [第一个错误的版本](./First_Bad_Version/code.go)
- [搜索插入位置](./Search_Insert_Position/code.go)

### Tips

- 注意判断边界位置
- 在区间变换时注意不同语言的开合条件
