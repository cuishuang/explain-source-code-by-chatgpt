# File: excelize/xmlDrawing.go

在Go生态excelize项目中，excelize/xmlDrawing.go文件用于处理Excel中的绘图相关操作。该文件提供了一系列结构体和方法，用于处理Excel中的图片、图表、形状等绘图元素。

下面逐一介绍这些结构体的作用：

1. xlsxCNvPr：用于定义绘图对象的非覆盖性属性。
2. xlsxHlinkClick：用于定义超链接操作。
3. xlsxPicLocks：用于定义图片锁定属性。
4. xlsxBlip：用于定义绘图对象的图像的引用。
5. xlsxStretch：用于定义图片的拉伸属性。
6. xlsxOff：用于定义对象在工作表上的位置偏移。
7. aExt：用于定义绘图对象的扩展属性。
8. xlsxPrstGeom：用于定义形状的几何变换。
9. xlsxXfrm：用于定义图形或形状的转换矩阵。
10. xlsxCNvPicPr：用于定义绘图对象的图像的非覆盖性属性。
11. xlsxNvPicPr：用于定义图形或形状的非覆盖性属性。
12. xlsxCTSVGBlip：用于定义绘图对象的SVG图像。
13. xlsxCTOfficeArtExtension：用于定义Office艺术扩展。
14. xlsxEGOfficeArtExtensionList：用于定义Office艺术扩展列表。
15. xlsxBlipFill：用于定义绘图对象的填充属性。
16. xlsxLineProperties：用于定义形状边框的属性。
17. xlsxSpPr：用于定义形状的属性。
18. xlsxPic：用于定义图片对象。
19. xlsxFrom：用于定义从哪个单元格开始的位置。
20. xlsxTo：用于定义到哪个单元格结束的位置。
21. xdrClientData：用于定义图表对象的客户端数据。
22. xdrCellAnchor：用于定义图表对象的单元格定位。
23. xlsxCellAnchorPos：用于定义图表对象的单元格定位的位置。
24. xlsxPoint2D：用于定义点的坐标。
25. xlsxWsDr：用于定义工作表的绘图集。
26. xlsxGraphicFrame：用于定义图表的框架。
27. xlsxNvGraphicFramePr：用于定义图表的非覆盖性属性。
28. xlsxGraphic：用于定义图表的属性。
29. xlsxGraphicData：用于定义图表的数据。
30. xlsxSle：用于定义图表的数据系列。
31. xlsxChart：用于定义图标对象。
32. xdrSp：用于定义形状对象。
33. xdrNvSpPr：用于定义形状对象的非覆盖性属性。
34. xdrCNvSpPr：用于定义形状对象的非覆盖性工作表属性。
35. xdrStyle：用于定义形状对象的样式属性。
36. aRef：用于定义填充颜色属性。
37. aScrgbClr：用于定义填充颜色属性的SCRGB颜色。
38. aFontRef：用于定义字体引用属性。
39. xdrTxBody：用于定义形状对象的文本框属性。
40. Picture：表示图片对象。
41. GraphicOptions：用于定义图形对象的选项。
42. Shape：用于定义形状对象。
43. ShapeLine：用于定义形状对象的线条样式属性。

这些结构体和方法提供了对Excel中绘图元素的读取、设置和处理功能，使得用户能够在Excel文件中插入图片、绘制图形和形状、添加超链接等操作。同时，还可以通过这些结构体和方法对已存在的绘图元素进行修改和删除。

