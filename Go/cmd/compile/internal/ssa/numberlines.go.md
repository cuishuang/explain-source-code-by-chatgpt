# File: numberlines.go

numberlines.go是一个Go语言程序，用于从标准输入读取文本行并为每行打上行号，然后将其写入标准输出。它可以将任何给定的文本文件转换为带有行号的文件，使其更容易阅读和理解。

该程序的主要目的是帮助程序员在调试代码时更容易地跟踪他们正在查看的代码行。它还有助于其他人阅读和理解代码，因为行号提供了一种更容易确定某些代码段的位置和顺序的方法。

程序的主要步骤如下：

1.创建一个标准输入的读取器和一个标准输出的写入器。

2.循环读取标准输入的每一行，并打印一个行号和该行的文本内容。

3.在每行之前加一个行号。

4.将包含行号的行写入标准输出。

总之，numberlines.go的作用是将文本文件转换为具有行号的文件，以更容易地跟踪和理解文件中的行号。




---

### Structs:

### fileAndPair

fileAndPair结构体是用来存储带有行号的文件内容和行号范围的结构体。

具体来说，fileAndPair结构体有两个字段：

- filename string类型，表示文件名
- pairs []filePair类型，表示一个长度不确定的filePair类型的切片，其中filePair表示一对行号的范围（起始行号和结束行号），用于表示一个文件中需要处理的行号范围。

在numberlines.go文件中，fileAndPair结构体是用来存储需要处理的文件列表和需要处理的行号范围列表的。同时，在读取文件内容时，也将内容和行号范围组合成fileAndPair结构体对象并存储在一个切片中。

fileAndPair结构体的作用是对文件信息和行号范围进行封装，方便在程序中传递和处理。它是numberlines.go文件中实现行号计数功能的重要数据结构之一。



### fileAndPairs

在numberlines.go中，fileAndPairs结构体用于存储文件名和其对应的行号和内容的列表。具体来说，fileAndPairs结构体由下面两个字段组成：

1. fileName string：表示文件名，也就是要统计行数的文件的名称。
2. pairs []pair：表示文件中的每一行内容和其对应的行号组成的列表，其中pair结构体由下面两个字段组成：
   - lineNum int：表示行号。
   - lineText string：表示行内容。

通过使用fileAndPairs这个结构体，可以使得程序能够方便地保存每个文件中的所有行号和内容，并进行相应的统计处理。同时，fileAndPairs结构体也可以输出每行的号码和内容，方便用户进行进一步分析和处理。



## Functions:

### isPoorStatementOp

isPoorStatementOp函数定义在numberlines.go文件中，作用是判断给定的Go语言操作符是否是可行的语句操作符。该函数接受一个操作符字符串类型参数，包括以下几种：

- ":"声明操作符
- "="赋值操作符
- "("逗号操作符

其中，赋值操作符和声明操作符是可行的语句操作符，而逗号操作符不是。因此，该函数会返回一个布尔值，表示被检测的操作符是否是可行的语句操作符。具体实现如下：

```go
// isPoorStatementOp returns whether a token is a poor statement operator.
func isPoorStatementOp(op string) bool {
	return op == ":" || op == "="
}
```

该函数在ParseFileAndPrint函数中被调用，用于判断行号是否需要累加。具体使用方法如下：

```go
for _, group := range f.Comments {
    num, endWithSemicolon := countGroup(group, func(op token.Operator) bool {
        if op == token.COLON {
            return true
        }
        return isPoorStatementOp(op.String())
    })
    if num != 0 {
        fmt.Printf("%d:%d\t%s\t(comment group)\n", start, end, filename)
        start += num
        end = start - 1
        if endWithSemicolon {
            end++
        }
    }
}
```

上面的代码会调用countGroup函数统计注释文本中的操作符数量，并判断是否需要累加行号。其中，isPoorStatementOp函数用于判断给定的操作符是否是可行的语句操作符。如果是，就将行号累加，否则将被忽略。



### nextGoodStatementIndex

在numberlines.go文件中，nextGoodStatementIndex函数的作用是查找下一个要显示的语句的索引。该函数的具体实现逻辑如下：

1. 从startIndex开始遍历语句列表，依次检查每个语句的行号是否在当前屏幕范围内（即是否在start和end之间）。

2. 如果该语句的行号在当前屏幕范围内，就将该语句的索引作为下一个要显示的语句的索引返回。

3. 如果该语句的行号不在当前屏幕范围内，则继续向下遍历，直到找到一个行号在当前屏幕范围内的语句，然后返回该语句的索引作为下一个要显示的语句的索引。

4. 如果遍历结束都没有找到一个行号在当前屏幕范围内的语句，则返回-1，表示下一个要显示的语句不存在。

实际上就是找出在当前显示区域内的最靠上的代码行，以该行的行号为起点，依次向下寻找下一个在范围内的代码行，直到找到下一个要显示的语句或者遍历结束。这里的“好的语句”指的是需要在屏幕上显示的语句。



### notStmtBoundary

notStmtBoundary函数的作用是判断给定的字符是否可以作为一个语句的边界。

具体来说，该函数接受一个rune类型的参数r，表示一个字符，然后返回一个bool类型的值，指示该字符是否可以作为一个语句的边界。如果该字符可以作为语句的边界，返回false；否则返回true。

这个函数主要用在文件行号显示的过程中，用于确定某个字符是否可以作为一个语句的边界。如果可以的话，则该字符不应该被算作该语句的一部分，而是应该被算作下一条语句的一部分。

举个例子，假设有以下一行代码：

fmt.Println("Hello, world!"); // This is a comment

在这行代码中，分号“;”可以作为一个语句的边界。因此，在将该行代码分割成多个语句的时候，我们需要忽略分号之后的内容，只取分号之前的部分作为语句。而在判断一个字符是否为分号的时候，就可以使用notStmtBoundary函数进行判断。

该函数的具体实现如下：

func notStmtBoundary(r rune) bool {
  return !unicode.IsOneOf([]*unicode.RangeTable{syntax.WhiteSpace, syntax.Comment}, r) &&
  !unicode.IsOneOf([]*unicode.RangeTable{syntax.Operator, syntax.Term, syntax.CloseBrace}, r) &&
  r != '#' &&
  r != ',' &&
  r != ';'
}

该函数首先判断该字符是否为单行注释或空白字符，如果是，就视为语句边界。然后判断该字符是否为操作符、限定符或闭合括号，如果是也视为语句边界。其次，该函数过滤掉了“#”、“,”和“;”这三个字符，因为这些字符通常用于注释、字符串或分割符，而不是作为语句的边界。最后，如果该字符既不是语句边界也不是需要过滤的字符，则返回true。



### FirstPossibleStmtValue

FirstPossibleStmtValue 函数是定义在 Go 语言的编译器源码中的一个函数，其作用是找到一个能够视为语句的 Go 语言表达式的节点的位置，并返回其节点值。该函数的主要用途是在 gofmt 等工具中自动化地添加行号信息时，识别语句所在的位置，为其添加行号信息。

具体地说，该函数首先遍历节点中的所有子节点，然后按特定的条件判断当前节点是否为语句。如果当前节点不是语句，则会继续递归其子节点，直到找到一个能够视为语句的节点为止。在实际的实现中，该函数所使用的判断条件主要包括节点类型、标识符等信息，并且会根据不同的节点类型使用不同的判断方法。

总之，FirstPossibleStmtValue 函数的主要作用是辅助实现 Go 语言编译器和其他工具中的行号信息功能，提高代码的可读性和可维护性。



### flc

在numberlines.go文件中，flc（也就是“file line counter”）是一个函数，它的作用是对指定的文件进行行计数，并将计数结果更新到一个map中。

具体而言，flc函数的输入参数为一个文件名（string类型）和一个map（map[string]int类型），用于存储计数结果。函数的实现通过调用bufio包中的NewScanner函数来读入文件内容，并对每一行进行计数。如果在map中已经存在了该文件名对应的计数结果，那么该计数结果将被更新；否则，将会新增一条记录。

具体的代码实现如下：

```go
func flc(fname string, m map[string]int) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if _, ok := m[fname]; ok {
		m[fname] += count
	} else {
		m[fname] = count
	}
}
```

总的来说，flc函数的作用是非常简单明了的：对指定的文件进行行计数，并将计数结果更新到一个map中。这样，我们就可以通过调用flc函数来实现对多个文件的行计数，并将计数结果存储在同一个map中，从而方便进行比较和分析。



### Len

在numberlines.go文件中，Len函数用于计算一个整数的位数。该函数通常用于在输出数字时为其确定正确的 padding 级别。具体地说，它返回给定整数的十进制表示中的数字数量。

例如，对于整数123456，它的十进制表示中有6个数字，因此Len(123456)将返回6。如果给定数字为0，则Len函数将返回1。

下面是Len函数的代码实现：

```
func Len(n int) int {
    if n == 0 {
        return 1
    }

    // 如果 n 是负数，则先将它转换为正数
    if n < 0 {
        n = -n
    }

    // 通过除以10的方式逐步减小 n 的值，直到它变为0
    count := 0
    for n > 0 {
        count++
        n /= 10
    }

    return count
}
```

在该函数中，我们首先检查给定数字是否是0。如果是，则直接返回1。否则，我们将给定数字 n 转换为其绝对值（即如果它是负数，则取其相反数）。之后，我们用一个循环逐步减小 n 的值，每次将其除以10，并将计数器 count 加1。最后，我们返回计数器的值，这就是给定数字的数字数。



### Less

在numberlines.go文件中，Less函数是一个比较函数（comparator），其作用是比较两个文件行的大小关系。Less函数接收两个参数，分别是要比较的两个行的[]byte切片。如果第一个行（a）小于第二个行（b），Less函数返回true，否则返回false。

在main函数中，可以看到Less函数是通过sort.Slice函数调用的。sort.Slice函数用于按照指定的比较函数对一个slice进行排序。在numberlines.go文件中，sort.Slice函数将每个输入文件的行按照从小到大的顺序排序，并将结果输出到标准输出。

Less函数的具体实现如下：

```
func Less(a, b []byte) bool {
    for i := 0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] < b[i]:
            return true
        case a[i] > b[i]:
            return false
        }
    }
    return len(a) < len(b)
}
```

该函数首先使用一个for循环遍历两个输入的行a和b。在每次循环中，它比较第i个字符在a和b中的大小。如果a[i] < b[i]，则返回true，表示a比b小；如果a[i] > b[i]，则返回false，表示a比b大。如果两个字符相同，则继续比较下一个字符。如果所有字符都相等，但是a的长度小于b的长度，则返回true，表示a比b小。否则，返回false，表示a比b大。 



### Swap

在Go语言的cmd包中，numberlines.go文件的Swap函数是一个用于交换两个整数的函数。

这个函数的作用是将切片中的两个元素交换位置。函数的参数是i和j，表示需要交换位置的两个索引值。在函数体内，将i所表示的元素值保存到临时变量temp中，然后将j所表示的元素值赋值给i所表示的元素，最后将temp中保存的值赋值给j所表示的元素。这样，i和j所表示的元素就完成了位置的交换。

这个函数在numberlines.go文件中可以用来实现排序算法。通过使用这个函数，我们可以在排序算法的过程中交换元素的位置，从而实现对元素的排序。这个函数的代码片段如下：

```go
func Swap(a []int, i, j int) {
    temp := a[i]
    a[i] = a[j]
    a[j] = temp
}
```

在排序算法中，我们可以使用这个函数完成对元素位置的交换操作，例如：

```go
func BubbleSort(a []int) {
    for i := 0; i < len(a)-1; i++ {
        for j := 0; j < len(a)-1-i; j++ {
            if a[j] > a[j+1] {
                Swap(a, j, j+1)
            }
        }
    }
}
```

在这个排序算法中，我们使用了Swap函数完成了对元素位置的交换，从而实现了对元素的排序。



### numberLines

numberLines这个函数是一个用于给文本添加行号的函数。具体而言，将原始文本分割成行，然后为每一行添加一个行号并生成一个新的文本。

该函数的实现方式如下：

- 首先使用strings.Split函数将原始文本分割成行；
- 然后使用fmt.Sprintf将行号和行的文本拼接起来，生成一个新的带有行号的字符串；
- 最后使用strings.Join函数将所有带有行号的字符串连接起来，生成最终的包含行号的文本。

例如，对于以下原始文本：

```
Hello, world!
How are you?
```

经过numberLines函数处理后，得到的结果为：

```
1: Hello, world!
2: How are you?
```

由此可见，numberLines这个函数的作用是给文本添加行号，方便阅读和处理。



