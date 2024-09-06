# File: parser.go

go中的parser.go文件是Go语言的解析器。它的作用是将Go源文件转换成抽象语法树（AST），以便进行各种静态分析和编译器操作。

具体来说，parser.go文件负责解析输入的Go代码并将其转换为抽象语法树。该文件中定义了许多函数和结构体，用于处理不同类型的语法结构，如函数、变量声明、表达式等。例如，parser.go文件中的函数parseExpr()用于解析一个表达式，并将其转换为一个AST节点。

解析器还负责检查语法错误，并将错误信息报告给用户。如果一个文件包含无效的语法结构或者不符合Go语言的规范，解析器会在解析期间抛出错误，让用户知道哪里出了问题。

总之，parser.go文件是Go语言编译器中非常重要的一部分，它负责将输入的源码翻译成可供编译器进行处理的数据结构，并在此过程中进行语法验证和错误处理。




---

### Var:

### reserved








---

### Structs:

### parser





### fixupRecord





### importError





## Functions:

### init





### initScanner





### Error





### error





### errorf





### expect





### expectEOL





### expectKeyword





### parseString





### parseUnquotedString





### next





### parseQualifiedName





### parseUnquotedQualifiedName





### parseQualifiedNameStr





### getPkg





### parseExportedName





### parseName





### deref





### parseField





### parseParam





### parseVar





### parseConversion





### parseConstValue





### parseConst





### reserve





### update





### parseNamedType





### parseInt64





### parseInt





### parseArrayOrSliceType





### parseMapType





### parseChanType





### parseStructType





### parseParamList





### parseResultList





### parseFunctionType





### parseFunc





### parseInterfaceType





### parsePointerType





### parseTypeSpec





### lookupBuiltinType





### parseType





### parseTypeAfterAngle





### parseTypeExtended





### skipInlineBody





### parseTypes





### parseSavedType





### parsePackageInit





### maybeCreatePackage





### parseInitDataDirective





### parseDirective





### parsePackage





