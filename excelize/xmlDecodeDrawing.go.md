# File: excelize/xmlDecodeDrawing.go

在Go生态excelize这个项目中，excelize/xmlDecodeDrawing.go文件的作用是解析Excel文档中的绘图部分。该文件包含了很多函数用于解码并提取绘图信息。

具体函数的作用如下：

1. decodeCellAnchor：解码单元格锚点，其中包括了绘图对象的位置、大小以及绑定的单元格等信息。
2. decodeCellAnchorPos：解码单元格锚点的位置信息，包括是否为绝对位置以及距离定位单元格的偏移量。
3. decodeSp：解码形状对象，包括形状的类型、Id等信息。
4. decodeNvSpPr：解码非视觉形状属性，包括形状名称、形状Id等信息。
5. decodeCNvSpPr：解码图形非视觉属性，包括形状名称、形状Id等信息。
6. decodeWsDr：解码绘图工作区的信息，包括工作区的大小、绘图对象列表等。
7. decodeCNvPr：解码图形非视觉属性，包括图形对象的Id等信息。
8. decodePicLocks：解码图片锁定属性，包括图片是否可以被编辑、调整大小等。
9. decodeBlip：解码图片数据，包括图片的Id、嵌入对象等。
10. decodeStretch：解码图片伸展，包括图片伸展的方式等。
11. decodeOff：解码偏移量，包括x、y方向上的偏移量。
12. decodeAExt：解码自动扩展属性，包括绘图对象是否根据单元格大小自动扩展等。
13. decodePrstGeom：解码预设几何体，包括几何体类型、几何体的参数等。
14. decodeXfrm：解码变换参数，包括平移、缩放、旋转等变换参数。
15. decodeCNvPicPr：解码图片非视觉属性，包括图片Id、嵌入对象等。
16. decodeNvPicPr：解码非视觉图片属性，包括图片的名称等。
17. decodeBlipFill：解码图片填充属性，包括图片数据、伸展等属性。
18. decodeSpPr：解码形状属性，包括形状的填充、线条颜色、阴影等属性。
19. decodePic：解码图片对象，包括图片的Id、填充属性等。
20. decodeFrom：解码起始位置，包括x、y方向上的偏移量等。
21. decodeTo：解码结束位置，包括x、y方向上的偏移量等。
22. decodeClientData：解码客户端数据，包括自定义属性等信息。

这些结构体对应于XML文件中的各种标签和属性，通过解析XML文件中的绘图信息，可以提取出绘图对象的位置、大小、属性等信息，以便在Go中对Excel文档中的绘图进行操作和处理。

