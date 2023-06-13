# File: ifacemethod3.go

ifacemethod3.go是Go语言标准库中cmd目录下的一个文件，其主要作用是实现了一个演示示例。该示例展示了在Go语言中如何定义接口、结构体，并将结构体的方法赋给接口，以便对接口进行操作。

ifacemethod3.go文件中定义了一个接口Printer，包含一个方法Print()。然后定义了两个结构体类型，分别是person和animal。它们都定义了一个方法Print()，并且实现了Print()方法。这样，person和animal类型就都实现了Printer接口。

在程序的主函数中，首先创建了一个person类型的变量p，并将其赋给一个Printer类型的变量。这是因为person类型已经实现了Printer接口，所以可以对其进行类型转换。然后通过对p调用Print()方法，实现了对person类型的操作。

接着，创建了一个animal类型的变量a，并将其赋给一个Printer类型的变量。同样的道理，animal类型也已经实现了Printer接口，所以也可以对其进行类型转换。然后通过对a调用Print()方法，实现了对animal类型的操作。最终输出结果为：

```
  name: Ken
  age: 11
I'm a cat and I can climb trees!
```

总之，ifacemethod3.go文件是Go语言官方为了演示接口、结构体和方法之间的关系而编写的示例程序，可以帮助Go语言开发者更好地理解Go语言的接口和结构体的概念，并掌握如何将结构体的方法赋给接口，以便实现对接口的操作。

