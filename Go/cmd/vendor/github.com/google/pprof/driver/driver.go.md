# File: driver.go

driver.go是Go语言的编译器源代码，其中定义了一个驱动程序的接口。这个接口定义了一个可以编译和链接程序的标准化方式，让编译器能够将源代码转换为可执行文件。

驱动程序的接口充分考虑了编译器的可扩展性和可定制性，用户可以根据自己的需求自定义不同的编译器驱动程序。在编译程序时，程序会根据目标平台的不同选择不同的驱动程序来进行编译。

其中的Driver接口定义如下：

// Driver is the main interface implemented by each of the specific
// compilers.
type Driver interface {
    // Name returns the name of the driver.
    Name() string

    // version should return a version string.  Something like
    // "5.1.0".
    Version() string

    // Init initializes the driver. cmdName is the name of the binary that
    // will be emitting IR.
    Init(c *Config, cmdName string)

    // IR returns true if the compiler can emit an IR file.
    IR() bool

    // StopBefore returns a sensible stopping point for the driver to exit.
    // The codes are generally defined by the language front end, e.g.
    // CFrontendStopBefore is used generally to stop before optimization.
    // StopBeforeHeaderParse is used to parse only the header files for C
    // and C++.
    StopBefore(stage string) bool

    // Run invokes the compiler with the specified command-line arguments.
    Run(c *Config, args []string) error

    // RequiresPkgConfig reports whether the given package is required
    // by the driver for the build tags and options it provides.
    RequiresPkgConfig(pkg string) bool

    // Machine returns the underlying *sys.Machine on which the driver
    // executes. If the machine is unknown, it should return nil.
    Machine() *sys.Machine

    // GCCMachine returns the gccgo target string to use
    // for this driver's machine.
    GCCMachine() string

    // Driver link order.
    LinkOrder([]*Action) ([]*Action, error)

    // Archive builds an archive from a set of object files.
    Archive(c *Config, w io.Writer, objs []*Action, pkgdir, limitArch string) error
}

驱动程序接口中的每个方法都有其定义的作用，例如Name()方法返回驱动程序的名称，Version()方法返回驱动程序的版本号等。总之，Driver接口的目的是为Go语言编译器提供一个可扩展和灵活的框架，以便根据不同的特定需求来创建不同的驱动程序。

