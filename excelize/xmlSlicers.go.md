# File: excelize/xmlSlicers.go

在Go生态excelize项目中，`excelize/xmlSlicers.go`文件的作用是处理Excel中的切片器（Slicer）相关的XML数据。切片器是一种数据筛选工具，可用于动态过滤数据。

该文件中定义了以下结构体和函数：

1. `xlsxSlicers`: 切片器的集合，包含多个`xlsxSlicer`对象。
2. `xlsxSlicer`: 单个切片器的定义，包含切片器的属性和切片器的缓存定义。
3. `xlsxSlicerCacheDefinition`: 切片器缓存的定义信息，包含切片器缓存的属性和切片器缓存的数据。
4. `xlsxSlicerCachePivotTables`: 切片器缓存所关联的透视表的集合，包含多个`xlsxSlicerCachePivotTable`对象。
5. `xlsxSlicerCachePivotTable`: 单个透视表与切片器缓存的关联关系，包含透视表的ID。
6. `xlsxSlicerCacheData`: 切片器缓存的数据信息，包含切片器缓存的项目和项目的属性。
7. `xlsxTabularSlicerCache`: 切片器缓存的定义信息，用于表格切片器，包含切片器缓存的属性和切片器缓存的数据。
8. `xlsxTabularSlicerCacheItems`: 表格切片器的项目信息，包含多个`xlsxTabularSlicerCacheItem`对象。
9. `xlsxTabularSlicerCacheItem`: 单个表格切片器的项目，包含项目的属性。
10. `xlsxTableSlicerCache`: 切片器缓存的定义信息，用于表格切片器。
11. `xlsxX14SlicerList`: Excel 2010版本中的切片器列表，包含多个`xlsxX14Slicer`对象。
12. `xlsxX14Slicer`: Excel 2010版本中的切片器定义，包含切片器的属性和切片器的缓存定义。
13. `xlsxX14SlicerCaches`: Excel 2010版本中的切片器缓存列表，包含多个`xlsxX14SlicerCache`对象。
14. `xlsxX14SlicerCache`: Excel 2010版本中的切片器缓存定义，包含切片器缓存的属性和切片器缓存的数据。
15. `xlsxX15SlicerCaches`: Excel 2013版本中的切片器缓存列表，包含多个切片器缓存对象。
16. `decodeTableSlicerCache`: 解析表格切片器缓存定义的函数。
17. `decodeSlicerList`: 解析切片器列表的函数。
18. `decodeSlicer`: 解析切片器定义的函数。
19. `decodeSlicerCaches`: 解析切片器缓存列表的函数。
20. `xlsxTimelines`: Excel中的时间轴定义。
21. `xlsxTimeline`: 时间轴的定义，包含时间轴的属性和时间轴的缓存定义。

这些结构体和函数定义了Excel中切片器的各种属性和数据，并提供了解析和处理切片器相关XML数据的功能。可以使用这些功能来读取、修改和生成Excel文件中的切片器信息。

