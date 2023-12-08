# File: excelize/xmlPivotCache.go

在Go生态excelize项目中，excelize/xmlPivotCache.go文件的作用是处理Excel中的PivotCache信息。PivotCache（数据透视表缓存）是Excel中用于支持数据透视表功能的数据存储结构。

该文件中包含了一系列结构体和相关方法，用于解析和生成PivotCache相关的XML内容。下面是对每个结构体的详细介绍：

1. xlsxPivotCacheDefinition：表示PivotCache定义的结构体，包含多个字段，如缓存源、工作表源、多个缓存字段等。其中最重要的是CacheSource字段，它指向了数据源。

2. xlsxCacheSource：表示PivotCache的缓存源，包含了数据源类型、连接字符串等信息。

3. xlsxWorksheetSource：表示PivotCache的工作表源，它描述了从哪个工作表中获取数据。

4. xlsxConsolidation：表示PivotCache的合并选项，用于指定如何合并多个来源数据。

5. xlsxCacheFields：表示PivotCache中的缓存字段集合，包含多个字段。

6. xlsxCacheField：表示PivotCache中的一个字段，包含字段的属性、类型等信息。

7. xlsxSharedItems：表示PivotCache中的共享项集合，用于描述一个字段的不同取值。

8. xlsxMissing：表示PivotCache中的缺失值。

9. xlsxNumber：表示PivotCache中的数值型数据。

10. xlsxTuples：表示PivotCache中的元组数据。

11. xlsxBoolean：表示PivotCache中的布尔型数据。

12. xlsxError：表示PivotCache中的错误类型数据。

13. xlsxString：表示PivotCache中的字符串型数据。

14. xlsxDateTime：表示PivotCache中的日期时间型数据。

15. xlsxFieldGroup：表示PivotCache中的字段分组。

16. xlsxCacheHierarchies：表示PivotCache中的缓存层级。

17. xlsxKpis：表示PivotCache中的指标。

18. xlsxTupleCache：表示数据透视表缓存中的元组缓存。

19. xlsxCalculatedItems：表示PivotCache中的计算项。

20. xlsxCalculatedMembers：表示PivotCache中的计算成员。

21. xlsxDimensions：表示PivotCache中的维度。

22. xlsxMeasureGroups：表示PivotCache中的度量组。

23. xlsxMaps：表示PivotCache中的映射表。

24. xlsxX14PivotCacheDefinition：表示PivotCache中的X14格式定义。

25. decodeX14PivotCacheDefinition：该方法用于解析X14格式的PivotCache定义。

通过这些结构体和相关方法，excelize/xmlPivotCache.go文件实现了对Excel中PivotCache的解析和生成，方便了对数据透视表功能的操作。

