# File: excelize/picture.go

在Go生态excelize项目中，excelize/picture.go文件的作用是处理Excel中的图片相关操作。下面逐个介绍该文件中的函数作用：

1. parseGraphicOptions：解析图形选项，返回解析后的图形选项信息。
2. AddPicture：向指定单元格添加图片。
3. AddPictureFromBytes：从字节添加图片。
4. addSheetLegacyDrawing：向工作表添加遗留绘图。
5. addSheetDrawing：向工作表添加绘图。
6. addSheetPicture：向工作表添加图片。
7. countDrawings：统计绘图对象的数量。
8. addDrawingPicture：添加绘图图片。
9. countMedia：统计媒体资源的数量。
10. addMedia：添加媒体资源。
11. GetPictures：获取工作表中的图片信息。
12. GetPictureCells：获取包含图片的单元格。
13. DeletePicture：删除指定单元格中的图片。
14. getPicture：获取指定单元格中的图片。
15. extractCellAnchor：提取单元格锚点。
16. extractDecodeCellAnchor：提取decode后的单元格锚点。
17. getDrawingRelationships：获取绘图关系。
18. drawingsWriter：绘图写入。
19. drawingResize：调整绘图大小。
20. getPictureCells：获取图片单元格。

以上函数分别实现了在Excel表格中操作图片的功能，包括添加、删除、获取、统计等操作。这些函数通过读取并解析Excel文件的相关内容，完成对图片的增删改查等操作。

