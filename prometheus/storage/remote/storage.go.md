# File: util/teststorage/storage.go

在Prometheus项目中，`util/teststorage/storage.go`文件是用于提供测试环境中的存储功能。

该文件中定义了TestStorage结构体，该结构体是基于Prometheus的存储层接口（Storage）实现的一个简单的存储引擎，用于在测试环境中模拟存储和查询数据。

TestStorage结构体具有以下几个主要成员：
1. `Mutex`：用于对并发访问进行互斥锁定。
2. `appenderMap`：用于存储各个时间序列的ExemplarAppender接口实例。
3. `globalExemplarAppender`：用于全局的ExemplarAppender接口实例。
4. `exemplarQueryable`：用于查询样本数据的ExemplarQueryable接口实例。

下面是TestStorage结构体的几个主要函数：
1. `New()`：创建一个新的TestStorage实例，并返回其指针。
2. `Close()`：关闭TestStorage实例，释放相关资源。
3. `ExemplarAppender(labels []labels.Label) (exemplar.Appender, error)`：为给定的标签创建并返回一个ExemplarAppender接口实例，用于添加样本数据。
4. `ExemplarQueryable(labels labels.Labels) promql.ExemplarQueryable`：为给定的标签返回一个ExemplarQueryable接口实例，用于查询样本数据。
5. `AppendExemplar(app exemplar.Appender, ts int64, lb labels.Labels, value float64) error`：通过给定的appender接口实例，将样本数据添加到存储中。

总结：
- TestStorage是一个模拟存储引擎，用于测试环境中模拟存储和查询数据。
- TestStorage提供了创建、关闭和操作样本数据的功能。
- ExemplarAppender用于添加样本数据，ExemplarQueryable用于查询样本数据。
- `TestStorage`可以通过 `New()` 函数创建，并通过 `Close()` 函数关闭。我们可以使用 `ExemplarAppender` 接口来创建并获取 `ExemplarAppender` 的实例，然后使用 `AppendExemplar()` 函数将样本数据添加到存储中，并可以使用 `ExemplarQueryable` 接口查询样本数据。

