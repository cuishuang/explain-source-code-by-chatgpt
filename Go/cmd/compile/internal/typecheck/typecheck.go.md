# File: typecheck.go

typecheck.go是Go编译器中的类型检查器。类型检查是编译器中的一个重要部分，它负责检查程序中的表达式是否具有正确的类型。如果表达式具有错误的类型，那么编译器就会报错。这种检查有助于在编译期间捕捉到许多错误，避免在运行时出现错误。

typecheck.go的功能包括：

1. 对所有表达式进行类型检查，确保它们都具有正确的类型。

2. 处理类型断言，以确保类型断言语句的断言类型与实际类型相匹配。

3. 处理类型转换，以确保转换类型是有效的。

4. 处理函数调用，以确保参数与函数签名匹配，返回类型与函数签名返回类型相同。

5. 处理结构体，确保结构体中的字段具有正确的类型。

6. 处理接口，确保接口类型与实现类型匹配。

7. 在类型检查期间生成类型信息，以便在生成代码时使用。

8. 执行类型推断，以确定变量的类型。

总之，typecheck.go是Go编译器中的非常重要的组件，它确保程序具有正确的类型，避免了由于类型错误引起的代码问题。




---

### Var:

### InitTodoFunc





### inimport





### TypecheckAllowed





### NeedRuntimeType





### traceIndent





### _typekind





### typecheck_tcstack





## Functions:

### AssignExpr





### Expr





### Stmt





### Exprs





### Stmts





### Call





### Callee





### tracePrint





### Resolve





### typecheckslice





### typekind





### cycleFor





### cycleTrace





### Func





### typecheck





### indexlit





### typecheck1





### typecheckargs





### RewriteNonNameCall





### RewriteMultiValueCall





### checksliceindex





### checksliceconst





### implicitstar





### needOneArg





### needTwoArgs





### Lookdot1





### typecheckMethodExpr





### derefall





### Lookdot





### nokeys





### hasddd





### typecheckaste





### errorDetails





### sigrepr





### fmtSignature





### fielddup





### typecheckarraylit





### visible





### nonexported





### checklvalue





### checkassign





### checkassignto





### stringtoruneslit





### checkmake





### checkunsafesliceorstring





### Conv





### ConvNop





