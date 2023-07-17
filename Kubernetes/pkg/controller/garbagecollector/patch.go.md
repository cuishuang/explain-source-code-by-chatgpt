# File: pkg/controller/garbagecollector/patch.go

pkg/controller/garbagecollector/patch.go是Kubernetes项目中垃圾回收控制器的一个核心文件。它实现了管理controller的finalizers、owner references以及从object中删除这些信息的功能。下面我们就把它分为两部分进行讲解。

第一部分：结构体

1. objectForFinalizersPatch

这个结构体描述了需要添加和删除finalizers的对象信息。它包括一个名称和finalizers列表。

2. ObjectMetaForFinalizersPatch

这个结构体代表了一个对象的元数据信息。它包括一个名称和一个finalizers列表。

3. objectForPatch

这个结构体描述了需要添加或删除owner reference的对象信息。它包括一个名称、UID、API版本和资源类型以及owner reference信息。

4. ObjectMetaForPatch

这个结构体代表了一个对象的元数据信息。它包括一个名称、UID、命名空间、labels、annotations以及owner reference信息。

5. jsonMergePatchFunc

这个结构体提供了一个函数，用于将两个JSON merge patch数据合并。它是通过调用标准库中的json MergePatch函数来实现的。

第二部分：函数

1. getMetadata

这个函数用于将指定的对象信息转化为ObjectMetaForPatch结构体。它主要包括名称、UID、命名空间、labels、annotations和owner reference信息。

2. patch

这个函数用于在一个对象上添加或移除finalizers。它接收一个需要修改的对象、需要添加或删除的finalizers、一个需要执行修改的patch类型和一个客户端，然后执行添加或删除finalizers操作。

3. deleteOwnerRefJSONMergePatch

这个函数用于删除owner reference。它接收一个需要修改对象的UID、需要删除的owner reference信息和一个字节数组，然后执行删除操作。

4. unblockOwnerReferencesStrategicMergePatch

这个函数用于在对象的owner references字段中删除一个阻塞存在。它接收一个需要修改的对象、需要删除的owner reference信息、需要执行修改的patch类型和一个客户端，然后执行删除操作。

5. unblockOwnerReferencesJSONMergePatch

这个函数用于在对象的owner references字段中添加一个阻塞存在。它接收一个需要修改的对象、需要添加的owner reference信息、需要执行修改的patch类型和一个客户端，然后执行添加操作。

总的来说，pkg/controller/garbagecollector/patch.go实现了Kubernetes项目中核心的垃圾回收器控制器功能。它提供了添加或删除finalizers的方法，进行owner reference的修改，以及从对象中删除这些信息的方法。这些方法对于Kubernetes体系的控制器来说非常重要，因为它们帮助控制器在修复或删除对象时，能够正确执行响应的操作，并保证对象的一致性。

