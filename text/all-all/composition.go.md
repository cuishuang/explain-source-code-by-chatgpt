# File: text/unicode/norm/composition.go

text/unicode/norm/composition.go文件是Go语言中text/unicode/norm包中用来处理Unicode字符组合的文件。它实现了Unicode文本的组合和规范化。

在这个文件中，有几个重要的结构体和函数：

1. ssState：该结构体用于表示组合状态，记录当前的组合状态。

2. streamSafe：该结构体用于表示流安全状态，记录文本是否处于流安全状态。

3. reorderBuffer：该结构体用于记录重新排序的缓冲区，以确保正确的组合顺序。

4. insertErr：该结构体用于表示插入错误，用于处理插入时可能发生的错误情况。

接下来是一些核心函数的作用：

1. first：获取第一个Unicode码点。

2. next：获取下一个Unicode码点。

3. backwards：获取前一个Unicode码点。

4. isMax：判断指定的Unicode码点是否为最大的可组合码点。

5. init：初始化组合的内部状态。

6. initString：初始化字符串。

7. setFlusher：设置刷新器。

8. reset：重置组合状态。

9. doFlush：执行刷新操作。

10. appendFlush：附加刷新操作。

11. flush：刷新组合序列。

12. flushCopy：刷新并复制组合序列。

13. insertOrdered：按顺序插入组合序列。

14. insertFlush：插入并刷新组合序列。

15. insertUnsafe：不安全地插入组合序列。

16. insertDecomposed：插入分解组合序列。

17. insertSingle：插入单个码点。

18. insertCGJ：插入组合序列CGJ（Combining Grapheme Joiner）。

19. appendRune：附加一个码点。

20. assignRune：分配一个码点。

21. runeAt：获取指定位置的码点。

22. bytesAt：获取指定位置的字节。

23. isHangul：判断指定码点是否为韩文字符。

24. isHangulString：判断指定字符串是否都是韩文字符。

25. isJamoVT：判断指定码点是否为Jamo VT字符。

26. isHangulWithoutJamoT：判断指定码点是否为没有Jamo T的韩文字符。

27. decomposeHangul：分解韩文字符。

28. combineHangul：组合韩文字符。

29. compose：对Unicode文本进行组合和规范化。

这些结构体和函数的目的是为了实现Unicode文本的正确组合和规范化，以确保不同的字符可以正确地组合成为一个整体。

