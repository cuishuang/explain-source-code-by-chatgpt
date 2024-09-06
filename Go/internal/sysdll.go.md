# File: sysdll.go

sysdll.go是Go语言标准库的一部分，主要用于在Windows平台上加载和访问系统动态链接库（DLL）。

在Windows系统下，许多系统功能和API都以DLL的形式存在，应用程序需要使用它们来实现一些操作，如网络编程、多线程、文件IO等。Go语言提供了在Windows平台上加载和使用DLL的功能，这些功能都封装在sysdll.go文件中。

sysdll.go文件中定义了一些Windows平台上常用的DLL函数，比如LoadLibrary、GetProcAddress、FreeLibrary等，这些函数可以用来加载DLL、获取函数地址和释放DLL资源。同时，sysdll.go还定义了一些常量、结构体和函数类型，以方便用户使用。

通过sysdll.go文件提供的功能，用户可以在Go语言中调用Windows平台上的系统API，并实现各种高级操作。比如，用户可以使用Windows DLL来操作网络、实现多线程、读取注册表等。




---

### Var:

### IsSystemDLL

IsSystemDLL这个变量是用于判断Windows系统动态链接库（.dll文件）是否为系统级别的dll，即是否为Windows操作系统的核心组件。在Windows操作系统中，一些核心的系统dll组件会被保护，禁止对其进行修改或篡改。

IsSystemDLL变量存放了一些Windows系统dll组件的名称，如果某个dll文件的名称匹配该列表中的任何一个名称，那么该dll就会被认为是系统级别的dll。这个变量的作用是确保在加载dll文件时，只有安全可靠的dll才能被加载到程序中，从而保证程序的正常运行和安全性。

该变量的作用在Go语言的内部实现中体现较为明显，一般不会在程序开发中使用。对于开发者而言，更为重要的是了解和遵守Windows系统对系统dll的保护机制，以确保程序的稳定和安全性。



## Functions:

### Add

Add这个func的作用是将DLL文件添加到系统DLL列表中。

在Windows系统中，每个进程都有一个系统DLL列表，其中包括可被共享的DLL文件，这些文件包含了一些操作系统API的实现。在程序运行时，系统会按照一定的顺序查找这些DLL文件，以获得需要的API函数的实现。

而Add函数的作用就是将用户制定的DLL文件添加到系统DLL列表中。具体来说，Add函数会将DLL文件的路径添加到系统环境变量PATH中，以使得系统能够定位到该DLL文件并使用其中的API函数。

Add函数的代码如下：

```go
func Add(dllpath string) error {
    added, err := hasAddedDLL(dllpath)
    if err != nil || added {
        return err
    }
    _, err = LoadLibrary(dllpath)
    if err != nil {
        return err
    }
    addedDLLs = append(addedDLLs, dllpath)
    return nil
}
```

其中，hasAddedDLL函数用于检查该DLL文件是否已经被添加到系统DLL列表中。如果该DLL文件已经被添加过了，则不会重复添加，并返回一个错误；如果还没有被添加过，则调用LoadLibrary函数将该DLL文件加载到内存中，并将其路径添加到addedDLLs列表中，以表示已经添加过。

总之，Add函数的作用是为程序提供一个方便的方式来将DLL文件添加到系统DLL列表中，以获得更多的API函数的支持。



