# File: inspector.go

inspector.go文件是Go语言自带的调试器——命令行交互式调试器（Command-line Interactive Debugger）的实现代码。它负责实现调试器交互界面的命令行参数解析、侦听程序进程的启动、断点设置、单步执行、变量查看等基本操作，帮助开发者在调试程序时定位问题，提高代码质量和调试效率。

具体来说，inspector.go文件中包含以下函数和结构体：

1. func main()：程序入口函数，解析命令行参数。

2. type inspectorCommand：一个命令结构体，包含名称和执行函数。

3. func (ic *inspectorCommand) Run(c *command) (state, error)：用于执行命令的函数，其中c是传递给命令的参数。

4. type command：命令结构体，包含各种命令及其参数。

5. func (c *command) execute() error：在command结构体创建之后，用于执行该command实例所描述的命令。

6. type handler：处理器结构体，包含用于通信、控制和监视调试进程的所有信息。主要方法包括attach、continue、next、step和stop等。

7. func newHandler(pid int, args []string, cwd string, exitChan chan error, interactive bool) (*handler, error)：创建一个新的处理器结构体，pid是要调试的进程ID，args是调试进程的参数，cwd是调试进程的工作目录，exitChan是一个用于监听目标进程退出的通道，interactive表示调试器是否使用交互模式。

8. func (h *handler) start() error：在新线程中启动处理器。

9. func (h *handler) detach()：从正在监控的进程上分离处理器。

10. func (h *handler) wait() error：等待处理器停止。

总而言之，inspector.go文件实现了一个命令行交互式调试器，使得开发者可以在调试程序时通过命令行输入各种指令对程序进行单步执行、变量查看、断点设置等操作，以快速定位问题。

