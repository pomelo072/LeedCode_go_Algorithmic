# 字符串

---

- [最长公共前缀](#最长公共前缀)

---

## 最长公共前缀

`easy`

- [最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/)

### 解题思路

- 通过对两两字符串之间比较找出最小的前缀子串, 直到子串为空或者遍历完整个string切片

### 解题代码

[Code](../Longest_Common_Prefix/code.go)

### Tips

混淆了`:=`和`=`赋值符号对覆盖单个变量的问题

- `:=` 会分配新的内存, 产生一个全新的变量并赋值
- `=` 只是将值赋予变量
