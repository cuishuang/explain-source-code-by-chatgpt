# File: filetransport_test.go

filetransport_test.go是Go语言标准库中net包中filetransport的测试文件。该文件包含了对filetransport的各种测试函数，以确保其正确性和可靠性。

filetransport是一个实现了基于文件的数据传输的net.Conn接口，它允许将输入和输出的数据保存在一个文件中。filetransport_test.go中测试函数主要包括对filetransport的构建、读写数据、关闭连接等方面的测试。

例如，TestRead函数测试了从filetransport读取数据的正确性，使用了一个测试文件来进行测试，并比较了读取的结果与文件中的数据是否一致。TestWrite函数则测试了向filetransport写入数据的正确性，并比较了写入的结果与文件中的数据是否一致。TestClose函数测试了正确关闭filetransport连接的能力。

通过这些测试函数，我们可以确保filetransport的实现满足设定的要求，并可靠地处理数据传输任务，为用户提供高质量的服务。

## Functions:

### checker

checker这个函数是用来对比两个文件内容是否一致的。在该文件中，通过创建一个临时文件，并将源文件复制到临时文件中去，再通过比较临时文件和目标文件的内容，来判断源文件和目标文件是否一致。

checker函数的具体实现如下：

```
func checker(src, dest string) error {
    sf, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sf.Close()
    df, err := os.Open(dest)
    if err != nil {
        return err
    }
    defer df.Close()
    s := bufio.NewReader(sf)
    d := bufio.NewReader(df)
    for {
        sb, serr := s.ReadByte()
        db, derr := d.ReadByte()
        if serr == io.EOF && derr == io.EOF {
            return nil
        }
        if serr != derr {
            return fmt.Errorf("files not equal: %v %v", serr, derr)
        }
        if sb != db {
            return fmt.Errorf("files not equal: %v %v", sb, db)
        }
    }
}
```

具体做法就是读取源文件和目标文件中的每一个字节，并逐个比较，如果遇到不同的字节，则说明两个文件不同。如果一直到文件末尾，两个文件都没有不同的地方，则说明两个文件一致。

checker函数的作用是对于测试用例中的预期输出文件和实际输出文件进行比对，从而判断实现是否正确。



### TestFileTransport

TestFileTransport函数是一个单元测试函数，它测试了filetransport.go文件中定义的FileTransport类型的方法。

具体来说，TestFileTransport函数创建了一个本地监听器(localListener)和另一个实际服务器(server)，并使用FileTransport类型的本机客户端(localClient)向服务器发送数据。TestFileTransport测试了FileTransport类型的数据传输是否正确，以及FileTransport类型在处理传输问题时是否能够正确地处理错误情况，以保证数据传输的正确性。

该函数在测试中使用了Go标准库中的net/http/httptest、net/http、io、bufio、os、path/filepath等模块来模拟客户端和服务器的通信，对于FileTransport类型的方法进行测试，以保证在使用FileTransport进行数据传输时的正确性和稳定性。



