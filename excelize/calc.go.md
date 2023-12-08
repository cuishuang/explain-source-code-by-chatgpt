# File: excelize/calc.go

在Go生态excelize项目中，`excelize/calc.go`文件是用于处理Excel中的计算逻辑的文件。

该文件中的变量`tokenPriority`定义了各个运算符的优先级，`month2num`定义了月份的名称与对应数字的映射关系，`dateFormats`、`timeFormats`和`dateOnlyFormats`定义了日期和时间的格式化方式，`addressFmtMaps`定义了地址的格式化方式，`formulaFormats`定义了公式的格式化方式，`formulaCriterias`定义了公式的条件格式，`conversionUnits`和`unitConversions`定义了单位之间的转换关系，`romanTable`定义了罗马数字的转换表。

`calcContext`用于存储计算的上下文信息，`cellRef`和`cellRange`表示单元格引用和单元格范围，`formulaCriteria`用于计算公式的条件，`ArgType`表示处理的参数类型，`formulaArg`表示公式的参数，`formulaFuncs`表示公式的函数，`conversionUnit`表示单位转换，`romanNumerals`表示罗马数字，`roundMode`表示四舍五入模式，`calcInverseIterator`表示计算迭代器，`trendGrowthMatrixInfo`表示趋势增长矩阵的信息，`calcDatabase`表示计算数据库。

`Value`用于表示Excel中的值，`ToNumber`用于将值转换为数字，`ToBool`用于将值转换为布尔值，`ToList`用于将值转换为列表，`CalcCellValue`用于计算单元格的值，`calcCellValue`用于计算单元格的值并返回结果，`getPriority`用于获取运算符的优先级，`newNumberFormulaArg`、`newStringFormulaArg`、`newMatrixFormulaArg`、`newListFormulaArg`、`newBoolFormulaArg`、`newErrorFormulaArg`、`newEmptyFormulaArg`用于创建不同类型的公式参数，`evalInfixExp`用于计算中缀表达式，`evalInfixExpFunc`用于计算中缀表达式的函数，`prepareEvalInfixExp`用于准备计算中缀表达式，`calcPow`、`calcEq`、`calcNEq`、`calcL`、`calcLe`、`calcG`、`calcGe`、`calcSplice`、`calcAdd`、`calcSubtract`、`calcMultiply`、`calcDiv`用于计算不同的运算，`calculate`用于计算公式，`parseOperatorPrefixToken`用于解析运算符前缀标记，`isFunctionStartToken`、`isFunctionStopToken`、`isBeginParenthesesToken`、`isEndParenthesesToken`、`isOperatorPrefixToken`、`isOperand`用于判断不同的标记类型，`tokenToFormulaArg`和`formulaArgToToken`用于在标记和公式参数之间进行转换，`parseToken`用于解析标记，`parseRef`用于解析引用，`prepareCellRange`用于准备单元格范围，`parseReference`用于解析引用，`prepareValueRange`、`prepareValueRef`分别用于准备值范围和值引用，`cellResolver`用于解析单元格，`rangeResolver`用于解析范围，`callFuncByName`用于通过函数名调用函数，`formulaCriteriaParser`用于解析公式条件，`formulaCriteriaEval`用于计算公式条件。

以上是`excelize/calc.go`文件中部分重要结构体和函数的介绍，涉及到的功能包括处理公式运算、处理日期和时间、处理单位转换、处理罗马数字、计算数据库等。

