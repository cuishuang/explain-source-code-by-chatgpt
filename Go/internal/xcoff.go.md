# File: xcoff.go

xcoff.go是Go语言编译工具链中的一个文件，主要用于处理IBM AIX系统上的XCOFF格式的可执行文件。XCOFF是针对IBM AIX操作系统设计的一种可执行文件格式，类似于Linux上的ELF格式。

在Go编译器中，xcoff.go文件主要包含了处理XCOFF格式文件的函数和结构体，以及用于在Go语言中表示XCOFF文件头、节表和符号表等元素的定义。它的作用是将Go语言编译后生成的二进制文件转换成适用于IBM AIX系统的XCOFF格式文件，以便在该系统上执行。

xcoff.go中的一些主要函数包括：

- func openXcoff(name string, mode FileMode) (file *xcoffFile, err error): 用于打开一个XCOFF格式的文件，并返回一个包含文件信息的xcoffFile结构体。
- func (f *xcoffFile) ReadAt(b []byte, off int64) (n int, err error): 用于读取XCOFF文件中指定偏移量处的字节序列。
- func (f *xcoffFile) ImportDirs() ([]ImportDir, error): 用于获取XCOFF文件中的导入目录信息。
- func (f *xcoffFile) Symbols() ([]Sym, error): 用于获取XCOFF文件中的符号表信息。

除了处理XCOFF格式文件外，xcoff.go还提供了一些与IBM AIX系统相关的辅助函数，如dlopen、dlsym和dlclose等，用于实现动态链接库相关的操作。

总的来说，xcoff.go是Go语言编译器工具链中的一个重要文件，在Go语言支持IBM AIX系统的编译、链接和执行方面发挥了重要作用。




---

### Structs:

### FileHeader32

在Go语言的内部包中，xcoff.go文件中的FileHeader32结构体用于表示PE文件的头部信息，其中包括如下重要的字段：

- Magic：用于标识该文件是否为xcoff格式的标识符，其长度为4个字节。
- NumSections：用于表示该文件中节的数量，其长度为2个字节。
- TimeDateStamp：用于表示该文件被创建的时间戳，其长度为4个字节。
- PointerToSymbolTable：用于表示符号表的起始地址，其长度为4个字节。
- NumberOfSymbols：用于表示符号表的大小，其长度为4个字节。
- OptionalHeaderSize：用于表示可选头的大小，其长度为2个字节。
- Characteristics：用于表示该文件的特性，其中包括是否是可执行文件、是否是DLL文件等。

FileHeader32结构体的作用是解析和读取PE文件的头部信息，可以通过该结构体中的字段获取PE文件的基本属性和特性，并进一步进行解析和处理。



### FileHeader64

FileHeader64结构体是用于表示XCOFF（Executable and Linking Format for IBM PowerPC）格式的64位可执行文件头部信息。XCOFF是一种用于PowerPC架构的可执行文件格式，被用于AIX和一些其他的操作系统中。FileHeader64结构体包含了64位可执行文件的所有必要信息，例如程序入口点、程序头表偏移地址、节头表偏移地址、符号表偏移地址等，这些信息用于描述可执行文件的基本属性和结构。其定义如下所示：

```go
type FileHeader64 struct {
    Magic         uint16
    TouOff        uint16
    LenMod        uint16
    Tstamp        uint32
    Symptr        uint64
    Nsyms         int32
    OptHeaderFlag uint16
    Flags         uint16
    Mach          uint16
    Reserved      uint16
    Ssize         uint32
    Ls1           uint32
    LoadAddr      uint64
    Toc           uint64
    Modtype       uint64
}
```

具体每个字段所表示的含义如下：

- Magic：用于标识XCOFF格式的魔数，其值为0x01df。
- TouOff：用于表示文件头大小。实际值需要乘以2的TouOff次方才是真正的大小（例如，当TouOff=1时，文件头大小为2字节，当TouOff=2时，文件头大小为4字节），同时也表示程序文本段的节头表偏移地址。
- LenMod：用于表示程序的文本段长度模数。程序的长度将会被16K对齐，LenMod就是用于表示对齐的大小。
- Tstamp：用于表示可执行文件的时间戳。
- Symptr：用于表示符号表偏移地址。
- Nsyms：用于表示符号表的数量。
- OptHeaderFlag：该字段已不再使用，应始终设置为0。
- Flags：用于表示可执行文件的标志位。例如，是否为动态库、是否启用栈溢出保护等等。
- Mach：用于表示目标机器架构的类型。
- Reserved：保留字段，应该始终设置为0。
- Ssize：用于表示字符串表的大小。
- Ls1：该字段已不再使用，应始终设置为0。
- LoadAddr：用于表示程序文本段的加载地址。
- Toc：用于表示程序的TOC（Table Of Contents）地址，一般只有PPC架构才会用到。
- Modtype：该字段已不再使用，应始终设置为0。

因此，FileHeader64结构体的作用就是定义了XCOFF格式的64位可执行文件的头部信息，使用该结构体可以方便地读取和解析可执行文件的相关信息。



### SectionHeader32

SectionHeader32结构体用于表示XCOFF文件中的32位节表头。XCOFF（Executable and Object Code Format）是一种用于执行和静态链接可执行文件和目标文件的二进制格式，常用于AIX操作系统。

SectionHeader32结构体包含以下字段：

- NameOffset uint32：该节表头对应节的名称在字符串表中的偏移量。
- PhysicalAddress uint32：如果该节被加载到内存中，则该字段为其在内存中的物理地址。否则为0。
- VirtualAddress uint32：如果该节被加载到内存中，则该字段为其在内存中的虚拟地址。否则为0。
- Size uint32：该节的大小，以字节为单位。
- SectionType uint32：该节的类型（如代码、数据、重定位等）。
- Flags uint32：该节的标志，如可读、可写、可执行等。
- Link uint16：该节与其他节的关联信息，例如重定位节。
- Info uint16：该节其他信息。
- Alignment uint32：该节在文件和内存中的对齐方式。
- EntrySize uint32：如果该节包含一组固定大小的项目，则该字段表示每个项目的大小。否则为0。

SectionHeader32结构体的作用是描述XCOFF文件中的节。使用节表头，可以在加载或链接节时确定节的大小、类型、位置和其他属性，以便正确地操作和管理代码和数据。



### SectionHeader64

在IBM AIX操作系统中的可执行文件和目标文件中，存储数据和代码的区域被称作Section，每个Section需要记录自身的一些元数据信息，如大小、偏移量、对齐方式等，SectionHeader64结构体就用来记录这些元数据信息。

在xcoff.go文件中，SectionHeader64结构体定义了AIX操作系统下Section的元数据信息，包括：

- Name：Section的名称，在xcoff文件中是ASCII字符串，长度为8位；
- Type：Section的类型，标识Section包含的数据或代码的类型，如代码，数据，BSS等；
- Flags：Section的各种属性标志，如是否可读、是否可写、是否可执行等；
- VirtAddr：Section的虚拟地址，指的是加载到内存时的地址，是相对于文件头中ImageBase的偏移量；
- FileOffset：Section在文件中的偏移量，指的是Section的数据在文件中的位置；
- Size：Section的大小，指的是Section中包含数据或代码的总长度，以字节为单位；
- Link：Section的链接信息，用来关联其他Section；
- Info：Section的附加信息，通常被用来存储Section的序号或关联信息；
- Align：Section的对齐方式，指的是Section中数据或代码的内存对齐方式，通常为2的幂；

通过这些元数据信息，程序就可以根据SectionHeader64结构体的定义，在运行时动态解析文件中的Section，加载到内存中执行。



### SymEnt32

SymEnt32是xcoff文件格式中的一个结构体，用于表示32位符号表的符号表条目。具体来说，它包含了以下字段：

- N: 符号的名称在字符串表中的偏移量
- V: 符号的值（地址或常量），如果符号是一个未定义的符号，这个值应该为0
- Type: 符号类型。这个字段是一个复合字段，包含了多个子字段，例如符号的存储类别、符号的标记、符号的类型等等
- Sclass: 符号的存储类别。这个字段指示了符号是全局符号、局部符号、静态符号等等
- Numaux: 符号条目后面附加的辅助信息的数量，每个辅助信息都是一个AuxEnt32结构体

SymEnt32结构体的作用是对xcoff文件中的符号表进行表示和解析。在xcoff文件中，符号表是一个非常重要的数据结构，它包含了所有定义和未定义的符号，以及对这些符号的引用和重定位信息。通过解析符号表，程序可以获得非常重要的信息，例如程序中使用的函数和全局变量的名称、地址等等。因此，SymEnt32结构体可以帮助开发人员分析和调试xcoff格式的可执行文件和库文件。



### SymEnt64

SymEnt64是一个结构体类型，用于表示XCOFF文件格式中的一个符号表项（Symbol Entry）。每个符号表项对应一个符号（Symbol），可以是函数、变量等程序中的标识符。

具体来说，SymEnt64结构体包含以下字段：

- Noffset uint32：符号名在文件中的偏移量（Offset），即符号表中的字符表索引（String Table Index）。
- Nvalue uint64：符号的地址或值，即指向符号所在区段（Section）的偏移量。对于未定义的外部符号，该字段为0。
- Nscnum int16：符号所在的区段编号（Section Number），从1开始。对于未定义的外部符号，该字段为0。
- Ntype uint16：符号类型（Symbol Type），包括函数、变量、常量等多种类型。具体类型定义在xcoff.h文件中。
- Nsclass uint8：符号存储类别（Storage Class），表示符号的作用域、链接性等信息。具体存储类别定义在xcoff.h文件中。
- Nnumaux uint8：附加数据项（Auxiliary Entries）的数量。这些数据项可用于存储符号的类型信息、大小信息等。

SymEnt64结构体的作用是将XCOFF文件中的符号表项与Go语言中的结构体对应起来，方便程序对符号表进行读取和解析。通过结构体的字段可以获取符号的名称、地址、类型以及存储类别等信息，从而对程序进行静态分析、重定位等操作。



### AuxFile64

AuxFile64是一个用于描述文件名和其它相关信息的XCOFF（Executable and Object Code Format）辅助记录。XCOFF是一种用于可执行文件和目标代码的二进制格式，常用于UNIX系统中。

AuxFile64结构体包含以下字段：

```go
type AuxFile64 struct {
    size      int32
    type_     int32
    name      [14]byte
    xfilename COFFString
    _         [4]byte
    xauxtype  int16
    xpad      [2]byte
}
```

- size：记录的大小（以字节为单位）。
- type_：记录的类型码。
- name：记录包含文件名字符串时的名称/值对的名称。
- xfilename：COFFString类型，表示文件名。
- xauxtype：表示附加记录类型的值。
- xpad：用于对齐。

AuxFile64结构体的作用是提供了一种描述文件名和其它相关信息的方式，例如附加文件的版本号、作者、编译器和其它指针等信息。该结构体可以在XCOFF文件中作为符号表的一部分出现，也可以与COFF（Common Object File Format）一起使用。在操作XCOFF文件或COFF文件时，可以使用该结构体来方便地获取文件名和其它相关信息。



### AuxFcn32

AuxFcn32是一个结构体类型，用于表示32位XCOFF文件中的函数辅助记录信息。函数辅助记录是一种附加信息，可以用于帮助解析函数符号。

AuxFcn32结构体包含以下字段：

- FcnSize uint32: 函数大小，表示函数在代码段中的大小。
- Lnno uint32: 包含行号记录的行号段（Line Number Section）的第一个符号表入口的索引。可以用于将行号信息与函数关联起来。
- ParmHid uint32: 参数的大小（以4字节为单位）和每个离散参数的类型。参数可以是整数、浮点数或结构体。这个字段记录参数信息，可以帮助解析函数参数。
- Freg uint32: 调用者保存的浮点寄存器（是一个位图）。被调用的函数在进入时必须将改变的浮点寄存器保存到堆栈上。
- Lreg uint32: 调用者保存的整数寄存器（是一个位图）。被调用的函数在进入时必须将改变的整数寄存器保存到堆栈上。
- Pcreg uint32: 当前函数用作栈指针的寄存器编号。在大多数情况下，这是R1寄存器（即SP，栈指针寄存器）。
- Pad [2]byte: 占位符，保证结构体大小为16字节。

当解析32位XCOFF文件中的函数符号时，可以使用AuxFcn32中的信息来获取函数大小、参数信息、寄存器信息以及行号信息等，以便更好地理解和调试程序的执行过程。



### AuxFcn64

AuxFcn64结构体是用于存储64位XCOFF可执行文件中的函数相关辅助信息的数据结构。

在XCOFF可执行文件中，每个函数都有一些附加信息，如函数的地址、大小、入口点等等。这些信息被存储在AuxFcn64结构体中，以便在链接期间和调试期间使用。

具体来说，AuxFcn64结构体的域包括：

- Tag：表示该结构体的类型，一般为AuxFcn64Tag。
- ParmLocal：表示函数参数和局部变量在堆栈中的大小。
- EntryParm：表示函数入口点的偏移量。
- RegMask：表示函数调用时需要保存哪些寄存器的位图。
- RegSave：表示在函数调用时需要保存的寄存器的位图。
- CrSave：表示在函数调用时需要保存的CR寄存器的位图。
- Toc：表示函数的TOC指针。
- SymIndex：表示该函数在符号表中的索引。

通过这些信息，链接器和调试器可以更准确地了解函数的内部结构和行为，帮助程序员进行调试和优化。



### AuxSect64

AuxSect64结构体是用于描述xcoff文件中的辅助节的数据结构。 

在xcoff文件格式中，辅助节是与主节相关联的附加信息。这些信息可以包括调试、符号表、重定位信息等。AuxSect64结构体存储了辅助节的元数据，如辅助节的名称、长度、类型、标志位等。 

具体来说，AuxSect64结构体包含以下字段： 

-  Name：辅助节名称，长度为8个字节。
-  Type：辅助节类型，表示该节所包含的信息类型，如SYMENT、SYMINFO、RELOC等。
-  Flags：标志位，表示辅助节的属性，如是否为只读节，是否为初始化节等。
-  Reserved：保留字段，长度为2个字节，未使用。
-  Size：辅助节长度，表示辅助节占用的字节数。
-  Link：指向辅助节所关联的主节的偏移量。
-  Info：附加信息，根据Type字段的不同而有所不同。

通过AuxSect64结构体，程序可以获取xcoff文件中的辅助节信息，从而进行相关的分析和处理。



### AuxCSect32

AuxCSect32结构体是用于表示COFF文件格式中的COMMON节（通常用于存储未初始化的全局变量）中的段的辅助信息。它包含与COMMON节相关的辅助信息，如节长度、节对齐、节的文件偏移量等。

具体来说，AuxCSect32结构体有如下作用：

1. 描述COMMON节的大小和位置：AuxCSect32中的Size表示COMMON节的大小，并且COMMON节的开始位置是从AuxCSect32的Offset字段指定的偏移量处开始的。这些信息使得可执行文件加载时可以快速定位COMMON节的位置和大小。

2. 描述COMMON节的对齐方式：AuxCSect32中的Align字段指定COMMON节的对齐方式，这对于使COMMON节中的数据按照正确的方式被加载到内存中是非常重要的。

3. 描述COMMON节的字符列表：AuxCSect32中的Characteristics字段用于存储一些二进制标志，这些标志描述了COMMON节的一些特性，例如是否可以重定向、是否可以被删除、是否含有COMDEF语句等等。

综上，AuxCSect32结构体是用于描述COMMON节的辅助信息，它将COMMON节的位置、大小、对齐方式、特性等都进行了标记和描述。这些信息使得可执行文件在加载COMMON节时能够快速地检索和定位COMMON节的内容。



### AuxCSect64

AuxCSect64是xcoff格式（Executable and Linkable Format for Extended COFF）中的一个辅助记录，记录了一个64位符号表部分的代码段的信息。

具体来说，一个xcoff文件可能包含多个代码段，每个代码段都有一个对应的符号表。AuxCSect64结构体记录了一个代码段的一些额外信息，包括：

- csectName：代码段名称
- length：代码段长度
- nreloc：该代码段关联的重定位表项数目
- nlinno：该代码段关联的行号表项数目
- flags：标志位，描述该代码段的属性，如是只读、是否是代码段等等
- reserved：保留字段

AuxCSect64结构体的作用是，为解析xcoff文件中的代码段提供辅助信息，帮助构建代码段相关的数据结构。同时，它也为调试器、反汇编器等工具提供了代码段的额外信息，方便进行代码分析和调试。



### LoaderHeader32

LoaderHeader32结构体用于描述XCOFF文件的加载头部部分。XCOFF是一种可执行文件格式，在某些Unix系统中使用，它包含了代码、数据以及其他相关信息。LoaderHeader32记录了加载程序所需的一些基本信息，以便正确地加载和执行XCOFF文件。

LoaderHeader32结构体中包含以下字段：

- Magic：标识XCOFF文件类型的标志，取值为0x01DF；
- Vstamp：版本戳；
- Tsize：代码段（text segment）的尺寸，以字节为单位；
- Dsize：数据段（data segment）的尺寸，以字节为单位；
- Bsize：bss段（block started by symbol）的尺寸，以字节为单位；
- Entry：指向程序入口点的地址；
- Text_start：代码段实际加载的起始地址；
- Data_start：数据段实际加载的起始地址；
- Tosb：堆栈底部偏移；
- Modtype：模块类型。

上述字段记录了XCOFF文件的基本信息，例如代码、数据和堆栈的尺寸、加载地址等。这些信息对于加载程序来说非常重要，可以帮助其正确地加载和执行XCOFF文件。



### LoaderHeader64

LoaderHeader64结构体的作用是用来描述XCOFF可执行文件的加载器头部。该结构体包含了可执行文件的多项信息，如入口点、段表和符号表等。具体来说，该结构体包含以下字段：

- Lversion：加载器头部的版本号。
- Lnr_relocs：需要重定位的节的数量。
- Lentloc：条目表中的条目数。
- Loffsets：偏移量表中的偏移量数量。
- Lreloc_pos：重定位节相对于文件起始位置的偏移量。
- Ledata_start：执行代码段的位置。
- Ledata_length：执行代码段的长度。
- Lbss_start：BSS段的位置。
- Lbss_length：BSS段的长度。
- Lloader_length：加载器占用的长度。

通过LoaderHeader64结构体，我们可以获取关于XCOFF可执行文件的各种信息，以便进行加载和执行。



### LoaderSymbol32

LoaderSymbol32是一个结构体，它用于描述XCOFF文件中的32位符号表。XCOFF是一种可执行文件格式，用于描述位于POWER和PowerPC处理器上运行的程序。每个XCOFF文件都包含一个符号表，其中包含所有全局和静态符号的定义和引用。这些符号用于编译器，链接器和调试器等工具。

LoaderSymbol32结构体是符号表中每个具有32位大小的符号记录的表示。它包含以下字段：

- Name：符号名称的字符串表示形式
- Value：符号值，指向符号所在的地址
- SectionNumber：符号所在的节（Section）的索引
- Type：符号类型
- StorageClass：符号存储类别
- NumberOfAuxSymbols：符号记录对应的附加符号数量

通过访问符号表中的LoaderSymbol32结构体，链接器可以确定符号地址、类型和存储类别等相关信息，以便正确地解决相互引用的符号。同时，调试器也可以使用符号表中存储的信息定位特定符号，帮助开发者调试程序。

因此，LoaderSymbol32结构体在XCOFF文件中的作用非常重要，它提供了符号表中符号的基本信息。



### LoaderSymbol64

LoaderSymbol64 这个结构体的作用是表示一个 64 位 XCOFF 文件中的符号表项。具体来说，它包含以下字段：

- Name: 字符串类型，表示符号的名称。
- Value: uint64 类型，表示符号的地址或常量值。
- SectionNumber: int16 类型，表示符号所属的节（section）编号。如果该符号是外部符号，则这个字段为 0。如果该符号是绝对符号，则这个字段为 -1。
- Type: uint16 类型，表示符号的类型。
- StorageClass: uint8 类型，表示符号的存储类别。
- NumberOfAuxEntries: uint8 类型，表示该符号在符号表中的后续附加信息条目数。

通过这些字段，LoaderSymbol64 可以表示一个 64 位 XCOFF 文件中的符号表项，包括符号的名称、地址、类型、存储类别和后续附加信息等。在加载 XCOFF 文件时，可以使用 LoaderSymbol64 来解析符号表，并根据需要进行符号重定位、符号绑定等操作。



### Reloc32

Reloc32是一个结构体，用于表示XCOFF二进制文件中的32位重定位类型。重定位是一种修改可执行文件中指向符号的指针地址的方法，以使得程序能够正确地执行。重定位类型指定了重定位需要做什么样的修改。

Reloc32结构体包含以下字段：

- Type uint8：重定位类型，指定了重定位需要做什么样的修改。
- SymbolNum uint24：符号表中符号的编号，该符号需要与重定位位置相关联。
- VirtualAddress uint32：虚拟地址，表示需要进行重定位的指针地址。

具体来说，Reloc32结构体描述了在XCOFF可执行文件中需要进行的重定位操作。其中，Type字段指定了重定位类型，例如R_XT_REL和R_TOC，表示需要使用何种方式进行重定位。SymbolNum字段指定了需要与重定位位置相关联的符号的编号，VirtualAddress字段指定了需要进行重定位的指针地址。通过这些信息，可执行文件加载器可以根据重定位类型将符号信息合并到正确的位置。

在程序编译时，编译器会通过对源代码的分析生成一组Reloc32结构体，它们描述了程序中需要进行的所有重定位操作。在生成可执行文件时，这些结构体将会被打包到ELF/XCOFF二进制文件中，以指示程序动态装载器在程序运行时如何修改指针地址，从而实现正确的符号引用、链接和运行。



### Reloc64

Reloc64是一个表示64位xcoff可重定位节中的重定位条目的结构体。它包含以下字段：

- Offset：重定位入口在节中的偏移量。
- SymbolNum：与重定位入口关联的符号表条目的编号。如果这个值是0，那么这个重定位入口是一个相对于节本身的未命名符号的引用。
- Type：重定位类型。它定义了如何修改地址，以使得指令或数据可以正确地链接到其他节或符号。
- Addend：如果重定位类型需要一个修正值，那么这里会存储该值，否则为0。

Reloc64的作用是在链接时解决符号引用，以便将符号定义的地址插入到对应的重定位入口中，从而正确地链接源代码。重定位类型指定了将该地址融入源代码的方式，如修改指针值或插入一个间接跳转。Addend字段则是相对于符号地址的修正值，用于使链接后的地址正确。



