# File: isprint.go

isprint.go是Go语言中strconv包的一个文件，其作用是判断一个Unicode字符是否可以用打印方式显示。

在Go语言的strconv包中，isprint.go文件中定义了一个函数isPrint(r rune) bool，该函数接收一个rune类型参数r，代表一个Unicode字符，返回一个bool类型值表示该字符是否可以用打印方式显示。具体实现算法如下：

1. 如果该字符是ASCII字符，只需要判断该字符的ASCII码是否在0x20~0x7E的范围内，表示该字符是可打印字符。

2. 如果该字符是非ASCII字符，需要判断该字符是否属于以下几种范畴中的一种：

   a. C0控制字符（0x00~0x1F和0x7F）和删除字符（0x7F）

   b. C1控制字符（0x80~0x9F）

   c. 字符解释器中指定的非打印字符，例如U+2028和U+2029

如果符合以上条件之一，则该字符不可打印，否则该字符可打印。

在Go语言中，strconv包的isPrint函数通常被用于验证用户的输入是否可打印，例如在处理文件名、密码等场景中，确保用户输入的字符不含有非法字符，以免造成不必要的程序错误。




---

### Var:

### isPrint16

isPrint16是一个长度为65536的布尔型数组，它的定义如下：

var isPrint16 = [utf8.MaxRune + 1]bool{
    // C0 control characters
    // 0x00 NUL, 0x01 SOH, 0x02 STX, 0x03 ETX, 0x04 EOT, 0x05 ENQ, 0x06 ACK, 0x07 BEL
    // 0x08 BS, 0x09 HT, 0x0A LF, 0x0B VT, 0x0C FF, 0x0D CR, 0x0E SO, 0x0F SI
    '\x00': false, '\x01': false, '\x02': false, '\x03': false,
    '\x04': false, '\x05': false, '\x06': false, '\x07': false,
    '\x08': false, '\x09': true, '\x0A': true, '\x0B': false,
    '\x0C': false, '\x0D': true, '\x0E': false, '\x0F': false,
    // C1 control characters
    // 0x80 PAD, 0x81 HOP, 0x82 BPH, 0x83 NBH, 0x84 IND, 0x85 NEL, 0x86 SSA, 0x87 ESA
    // 0x88 HTS, 0x89 HTJ, 0x8A VTS, 0x8B PLD, 0x8C PLU, 0x8D RI, 0x8E SS2, 0x8F SS3
    // 0x90 DCS, 0x91 PU1, 0x92 PU2, 0x93 STS, 0x94 CCH, 0x95 MW, 0x96 SPA, 0x97 EPA
    // 0x98 SOS, 0x99 SGCI, 0x9A SCI, 0x9B CSI, 0x9C ST, 0x9D OSC, 0x9E PM, 0x9F APC
    '\x80': false, '\x81': false, '\x82': false, '\x83': false,
    '\x84': false, '\x85': false, '\x86': false, '\x87': false,
    '\x88': false, '\x89': false, '\x8A': false, '\x8B': false,
    '\x8C': false, '\x8D': false, '\x8E': false, '\x8F': false,
    '\x90': false, '\x91': false, '\x92': false, '\x93': false,
    '\x94': false, '\x95': false, '\x96': false, '\x97': false,
    '\x98': false, '\x99': false, '\x9A': false, '\x9B': false,
    '\x9C': false, '\x9D': false, '\x9E': false, '\x9F': false,
    // other ASCII printable characters
    ' ': true, '!': true, '"': true, '#': true, '$': true, '%': true, '&': true, '\'': true,
    '(': true, ')': true, '*': true, '+': true, ',': true, '-': true, '.': true, '/': true,
    '0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true,
    '8': true, '9': true, ':': true, ';': true, '<': true, '=': true, '>': true, '?': true,
    '@': true, 'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true,
    'H': true, 'I': true, 'J': true, 'K': true, 'L': true, 'M': true, 'N': true, 'O': true,
    'P': true, 'Q': true, 'R': true, 'S': true, 'T': true, 'U': true, 'V': true, 'W': true,
    'X': true, 'Y': true, 'Z': true, '[': true, '\\': true, ']': true, '^': true, '_': true,
    '`': true, 'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true,
    'h': true, 'i': true, 'j': true, 'k': true, 'l': true, 'm': true, 'n': true, 'o': true,
    'p': true, 'q': true, 'r': true, 's': true, 't': true, 'u': true, 'v': true, 'w': true,
    'x': true, 'y': true, 'z': true, '{': true, '|': true, '}': true, '~': true,
    // keep NUL just to be safe, even though it's not a printable so it is always false
    '\uFFFD': false, // Unicode replacement character
}

isPrint16数组的作用是判断一个unicode码点是否对应可打印字符。Unicode码点的范围是0x0000到0x10FFFF，因此这个数组的长度为65536，即utf8.MaxRune+1。数组中存储着每个Unicode码点是否可打印的信息，true表示可打印，false表示不可打印。这个数组在strconv包中的几个函数中都被使用到，比如在Atoi函数中，用于判断字符串是否能被转换为整数。



### isNotPrint16

isNotPrint16是一个包级别变量，其类型为[256]bool，表示ASCII字符集中每个字符是否为不可打印字符（即不是可见字符）。其中，isNotPrint16[x]为true表示字符x是不可打印字符，反之为false表示字符x是可打印字符。

在isprint.go中，isNotPrint16变量主要用于帮助isPrintFunc函数检查给定的字符是否为可打印字符。isPrintFunc是一个函数类型，其定义为 func(rune) bool，用于判断给定的rune是否为可打印字符。调用isPrintFunc函数时，首先会检查给定的字符rune是否在ASCII字符集范围内，如果是，则会根据isNotPrint16变量来判断rune是否为可打印字符。如果rune在ASCII字符集外，则默认认为其为可打印字符。

总之，isNotPrint16变量是用于检查ASCII字符集中哪些字符是不可打印字符，帮助判断输入的字符是否为可打印字符。



### isPrint32

isPrint32是一个长度为256的布尔类型的数组，用于判断一个给定的Unicode字符是否为可打印字符。在isPrint32数组中的每个索引位置上，如果该索引对应的Unicode字符是可打印字符，则isPrint32数组的对应位置的值为true，否则为false。

该变量的作用是帮助实现Go语言字符串转换函数中的一些特定功能，例如在解析浮点数时，需要过滤掉一些不可打印的字符，而这些字符可能会影响浮点数的转换结果。

isPrint32变量依赖于Go语言中版本的Unicode表格，因此在提供相同功能的同时，也保证了跨平台的兼容性。



### isNotPrint32

isNotPrint32变量是一个32位的位图，每一位表示一个Unicode字符是否不可打印（non-printable）。具体来说，如果isNotPrint32的第n位为1，表示Unicode编码为n的字符不可打印；如果isNotPrint32的第n位为0，表示Unicode编码为n的字符可打印。

这个变量的作用是在strconv包中用于判断一个Unicode字符是否可打印，如果该字符可打印，则返回true；否则返回false。在字符串转换时，如果给定的字符串中包含非可打印字符，则其可能会被转换成其他字符（如"\t"会被转换成制表符），造成结果偏差，因此需要过滤掉这些非可打印字符。



### isGraphic

isGraphic是一个包级别变量（即在整个strconv包内都可见），类型为一个长度为256的布尔类型数组。这个数组对应了ASCII码表中256个字符的可打印性（printable）。

在isprint.go文件中，isGraphic变量被定义和初始化为一个长度为256的布尔类型数组，其中每一个元素表示对应的ASCII码值是否是可打印字符（graphic character）。当一个字符被传入strconv包的IsGraphic函数时，该函数会首先检查这个字符的ASCII码值，再查看isGraphic数组对应的元素是否为true，如果是，那么该字符就是可打印的，否则就是不可打印的。

isGraphic变量的作用是提供一个可重用的工具，允许通过一个简单的检查过程来判断一个给定的字符是否是可打印字符。在strconv包的其他函数中，例如ParseInt和QuoteRune，都会用到这个isGraphic变量来判断字符的可打印性，以便执行一些相应的操作，例如将字符转换为数字或者将字符转义为相应的转义序列。

总之，isGraphic变量是strconv包内部实现的一个重要部分，用于提供一种简单的检查机制，以判断一个字符是否是可打印的，帮助其他函数实现相应的操作。



