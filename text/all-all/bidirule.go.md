# File: text/secure/bidirule/bidirule.go

在Go的text项目中，text/secure/bidirule/bidirule.go文件的作用是实现文本双向规则（bidirectional rule）的处理。双向规则是指在一段文本中同时存在从左到右（LTR）和从右到左（RTL）的文字的情况。该文件中的函数和结构体用于解析和处理双向规则。

ErrInvalid是一个错误变量，表示提供的文本不是有效的双向规则。

transitions是一个字符串切片，用于存储当前字符状态和下一个字符的状态转换规则。

asciiTable是一个ASCII码表，用于检测字符是否属于RTL（从右到左）或LTR（从左到右）方向。

ruleState是一个枚举类型，表示字符的状态，可以是LTR、RTL、Neutral等。

ruleTransition结构体表示字符状态的转换规则。

Transformer结构体用于存储双向规则的转换器。它包含了转换规则的数据以及转换过程中的状态信息。

Direction是一个整型常量，表示文本的方向，可以是LTR或RTL。

DirectionString是一个函数，根据给定的方向常量返回相应的字符串表示。

Valid函数用于检查给定的文本是否是有效的双向规则。

ValidString函数用于检查给定的字符串表示的方向是否是有效的。

New函数用于创建一个新的转换器。

isRTL函数用于判断给定的字符是否属于RTL方向。

Reset函数用于重置转换器的状态。

Transform函数用于根据转换规则将文本从一种方向转换为另一种方向。

Span函数用于确定给定文本中从指定位置开始的字符的方向。

init函数用于初始化转换规则和转换表。

advance函数用于根据转换规则和字符状态将转换器向前移动一个字符。

advanceString函数用于根据转换规则和字符状态将转换器向前移动一个字符串。

