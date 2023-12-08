# File: excelize/drawing.go

在excelize/drawing.go文件中，定义了一系列与图表和绘图相关的功能函数，用于在Excel文件中插入和处理图表以及绘图。

1. prepareDrawing：准备绘图区域的信息。
2. prepareChartSheetDrawing：准备图表工作表的绘图信息。
3. addChart：添加一个图表到工作表。
4. drawBaseChart：绘制基本的图表（非3D，非饼图）。
5. drawDoughnutChart：绘制圆环图。
6. drawLineChart：绘制折线图。
7. drawLine3DChart：绘制3D折线图。
8. drawPieChart：绘制饼图。
9. drawPie3DChart：绘制3D饼图。
10. drawPieOfPieChart：绘制饼图中的饼图。
11. drawBarOfPieChart：绘制饼图中的条形图。
12. drawRadarChart：绘制雷达图。
13. drawScatterChart：绘制散点图。
14. drawSurface3DChart：绘制3D曲面图。
15. drawSurfaceChart：绘制曲面图。
16. drawBubbleChart：绘制气泡图。
17. drawChartShape：绘制图表形状。
18. drawChartSeries：绘制图表系列。
19. drawChartSeriesSpPr：绘制图表系列形状属性。
20. drawChartSeriesDPt：绘制图表系列数据点。
21. drawChartSeriesCat：绘制图表系列分类（X轴）。
22. drawChartSeriesVal：绘制图表系列值（Y轴）。
23. drawChartSeriesMarker：绘制图表系列的标记点。
24. drawChartSeriesXVal：绘制图表系列的X值。
25. drawChartSeriesYVal：绘制图表系列的Y值。
26. drawCharSeriesBubbleSize：绘制图表系列的气泡大小。
27. drawCharSeriesBubble3D：绘制图表系列的3D气泡。
28. drawChartNumFmt：绘制图表的数字格式。
29. drawChartDLbls：绘制图表的数据标签。
30. drawChartSeriesDLbls：绘制图表系列的数据标签。
31. drawPlotAreaCatAx：绘制绘图区域的分类轴。
32. drawPlotAreaValAx：绘制绘图区域的值轴。
33. drawPlotAreaSerAx：绘制绘图区域的系列轴。
34. drawPlotAreaTitles：绘制绘图区域的标题。
35. drawPlotAreaSpPr：绘制绘图区域的形状属性。
36. drawPlotAreaTxPr：绘制绘图区域的文本属性。
37. drawChartLn：绘制图表的线条属性。
38. drawingParser：解析绘图信息。
39. addDrawingChart：添加一个图表到绘图区域。
40. addSheetDrawingChart：在工作表中添加一个图表到绘图区域。
41. deleteDrawing：删除绘图信息。
42. extractEmbedRID：提取嵌入关联标识符。
43. deleteDrawingRels：删除绘图关联。
44. genAxID：生成轴的标识符。

这些函数通过对Excel文件的底层数据结构进行操作，实现了在Excel文件中插入和处理图表以及绘图的功能。可以通过调用这些函数，根据传入的参数和数据，生成对应的图表和绘图。

