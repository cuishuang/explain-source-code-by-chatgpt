# File: tools/gopls/internal/lsp/cache/view.go

文件`tools/gopls/internal/lsp/cache/view.go`的作用是实现了一个视图（view）的缓存，用于维护整个工作区的状态，并提供对源代码以及相关信息的访问和管理。

下面是每个变量和结构体的作用：

1. `checkPathCase`：用于检查路径大小写的变量。
2. `timeNow`：用于获取当前时间的函数。
3. `modFlagRegexp`：用于匹配go.mod文件中的module行的正则表达式。

结构体：

1. `View`：表示一个视图，它是对工作区状态的封装，包含了源代码文件和其他相关信息的缓存。该结构体提供了许多方法来处理工作区中的文件、目录和模块等。
2. `workspaceInformation`：表示工作区的信息，包括根目录、模块文件和模块路径等。
3. `ViewType`：表示视图的类型，可以是文件视图、模块视图或目录视图。
4. `go111module`：表示go111module环境变量的值，可以是"on"、"off"或"auto"。
5. `goEnv`：表示环境变量的集合，包括GOPATH、GO111MODULE等。
6. `workspaceMode`：表示工作区的模式，可以是模块模式、GOPATH模式或临时模式。
7. `ignoreFilter`：表示过滤器，用于排除不感兴趣的文件。

函数：

1. `effectiveGO111MODULE`：获取有效的go111module环境变量的值，考虑环境变量的设置以及工作区的模式等。
2. `ViewType`：获取视图的类型。
3. `moduleMode`：检查当前是否处于模块模式。
4. `GOWORK`：获取工作目录。
5. `GO111MODULE`：获取go111module环境变量的值。
6. `load`：加载视图。
7. `String`：返回结构体的字符串表示。
8. `vars`：获取给定目录的环境变量。
9. `ID`：返回视图的唯一标识符。
10. `tempModFile`：获取临时模块文件。
11. `Name`：返回视图的名称。
12. `Folder`：返回视图对应的目录。
13. `minorOptionsChange`：检查是否发生次要选项的更改。
14. `SetFolderOptions`：设置文件夹的选项。
15. `setViewOptions`：设置视图的选项。
16. `viewEnv`：获取视图的环境变量。
17. `RunProcessEnvFunc`：运行一个进程，并设置环境变量。
18. `fileHasExtension`：检查文件是否具有给定的扩展名。
19. `locateTemplateFiles`：在视图中定位模板文件。
20. `contains`：检查列表是否包含给定的元素。
21. `filterFunc`：过滤函数，用于过滤不感兴趣的文件。
22. `relevantChange`：检查给定的修改是否是相关的。
23. `markKnown`：标记文件为已知。
24. `knownFile`：检查文件是否是已知的。
25. `shutdown`：关闭视图。
26. `IgnoredFile`：检查文件是否被忽略。
27. `newIgnoreFilter`：创建一个新的过滤器来排除不感兴趣的文件。
28. `ignored`：返回文件是否被忽略。
29. `checkIgnored`：检查文件是否被忽略。
30. `Snapshot`：表示视图的快照，包含了视图中文件的信息和状态。
31. `getSnapshot`：获取当前视图的快照。
32. `initialize`：初始化视图。
33. `loadWorkspace`：加载工作区的信息。
34. `invalidateContent`：使文件的内容无效。
35. `getWorkspaceInformation`：获取工作区的信息。
36. `findWorkspaceModFile`：在工作区中查找模块文件。
37. `findRootPattern`：查找工作区的根目录模式。
38. `defaultCheckPathCase`：检查路径的大小写。
39. `IsGoPrivatePath`：检查给定的路径是否是Go私有路径。
40. `ModuleUpgrades`：表示模块的升级信息，包括当前模块和升级目标模块的版本等。
41. `RegisterModuleUpgrades`：注册模块的升级信息。
42. `ClearModuleUpgrades`：清除模块的升级信息。
43. `Vulnerabilities`：表示模块的安全漏洞信息。
44. `SetVulnerabilities`：设置模块的安全漏洞信息。
45. `GoVersion`：表示Go的版本信息，包括主版本号、次版本号和补丁号等。
46. `GoVersionString`：获取Go版本的字符串表示。
47. `globsMatchPath`：检查给定的路径是否与通配符模式匹配。
48. `vendorEnabled`：检查是否启用了vendor模式。
49. `allFilesExcluded`：检查所有文件是否被排除。
50. `pathExcludedByFilterFunc`：检查路径是否被过滤函数排除。
51. `pathExcludedByFilter`：检查路径是否被过滤器排除。
52. `buildFilterer`：构建一个过滤器函数，用于过滤不感兴趣的文件。

