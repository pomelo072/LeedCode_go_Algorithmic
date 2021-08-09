# 技巧

---
- [双指针](#双指针)
- [滑动窗口](#滑动窗口)


---


## 双指针

- [有序数组的平方](./array.md#有序数组的平方)
- [移动零](./array.md#移动零)
- [两数求和--输入有序数组](./array.md#两数求和_输入有序数组)
- [链表的中间结点](./array.md#链表的中间结点)
- [删除链表的倒数第N个结点](./array.md#删除链表的倒数第N个结点)

## 滑动窗口

- [无重复字符的最长子串](#无重复字符的最长子串)
- [字符串的排列](#字符串的排列)

### 无重复字符的最长子串

`medium`

- [无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

#### 解题思路

- 使用滑动窗口, 窗口包含当前不重复字符, 窗口向前扩张
- 使用map哈希, 搜索重复字符, 降低搜索时间, 当搜索到重复字符, 将窗口与重复字符的右侧重新扩展
- 单独的常数变量记录最长子串长度

#### 解题代码

[Code](../Longest_Substring_Without_Repeating_Characters/code.go)

#### Tips

- 两个月前字节Camp的第一题, 由于算法水平太差A不出来, 害
- 没有很好的掌握滑动窗口的思想, 又涉及字符串令我实在有些难顶.

### 字符串的排列

`medium`

- [字符串的排列](https://leetcode-cn.com/problems/permutation-in-string/)

#### 解题思路


#### 解题代码

[Code](../Permutation_in_String/code.go)

#### Tips

