# 字符串

---

- [最长公共前缀](#最长公共前缀)
- [反转字符串](#反转字符串)
- [反转字符串中的单词3](#反转字符串中的单词3)

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


## 反转字符串

`easy`

- [反转字符串](https://leetcode-cn.com/problems/reverse-string/)

### 解题思路

很容易的一题
- 双指针, 指向头和尾, 向中间移动
- 每次移动相互交换元素
- 到头指针大于尾指针, 完成

### 解题代码

[Code](../Reverse_String/code.go)

## 反转字符串中的单词3

`easy`

- [反转字符串中的单词3](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/)

### 解题思路

- 使用双指针, 顺序遍历string
- 遍历指针指向空格时, 从前指针到遍历指针之间的字符, 逆序插入`[]byte`
- 再加一个空格字符, 前指针跳到遍历指针位置, 遍历指针继续向后遍历, 重复第二步
- 遍历完整个字符串, 完成

### 解题代码

[Code](../Reverse_Words_in_a_String_III/code.go)

### Tips

- 使用了`[]byte`和`strings.Builder`来实现
- 理论上`strings.Builder`的性能应该好于`[]byte`, 本地对单实例的Benchmark测试也是如此
- 猜测在长字符串的测试用例下, `[]byte`的性能会更好
