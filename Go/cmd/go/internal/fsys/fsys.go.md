# File: fsys.go

fsys.go是一个标准库包，其作用是提供一个通用的文件和目录接口，可以对文件系统进行一些常用的操作，如读取、写入、复制、创建目录和删除文件等。

该文件定义了一组接口，这些接口支持在任何文件系统中进行文件和目录操作，而不必考虑底层文件系统的实现细节。它还提供了一个文件系统的实现，该实现可以在本地文件系统、远程文件系统和内存中创建虚拟文件系统。

这个文件包含了一些基本的类型和方法来表示文件系统中的文件和目录，例如：

- fs.File：用于表示一个文件，在此类型上可以执行读，写和关闭等操作。

- fs.DirEntry：用于表示一个目录项，它包含了一个文件或者目录的元信息，例如文件名和是否是目录等。

- fs.ReadDirFS：用来读取一个目录下的所有文件和子目录。

- fs.ReadFileFS：用来读取一个文件的内容。

通过上述方法和类型，可以方便地对文件和目录进行操作，同时也提供了一系列的安全校验和错误处理，以确保操作的可靠性，是一个提高开发效率的有用工具。




---

### Var:

### doTrace





### traceFile





### traceMu





### gofsystrace





### gofsystracelog





### gofsystracestack





### OverlayFile





### overlay





### cwd





### errNotDir








---

### Structs:

### OverlayJSON





### node





### fakeFile





### missingFile





### fakeDir





## Functions:

### Trace





### init





### isDir





### isDeleted





### canonicalize





### Init





### initFromJSON





### IsDir





### parentIsOverlayFile





### nonFileInOverlayError





### readDir





### ReadDir





### OverlayPath





### Open





### OpenFile





### openFile





### IsDirWithGoFiles





### walk





### Walk





### Lstat





### Stat





### overlayStat





### Name





### Size





### Mode





### ModTime





### IsDir





### Sys





### String





### Name





### Size





### Mode





### ModTime





### IsDir





### Sys





### String





### Name





### Size





### Mode





### ModTime





### IsDir





### Sys





### String





### Glob





### cleanGlobPath





### volumeNameLen





### cleanGlobPathWindows





### glob





### hasMeta





