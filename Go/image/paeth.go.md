# File: paeth.go

paeth.go这个文件的作用是实现了Paeth滤波器函数，用于PNG图像编码时使用。Paeth滤波器是一种差分滤波器，它可以通过当前像素与周围3个像素之间的预测值差值中最小的那一个来减少压缩后的图像数据大小，从而提高压缩率。

Paeth滤波器的实现过程如下：

1. 获取当前像素(x,y)以及周围3个像素(A,B,C)的值；

2. 计算当前像素与上一个被滤波像素像素(a,b)的预测值p(a,b)；

3. 计算当前像素与周围3个像素之间的差值(px,py,pA,pB,pC)；

4. 选择上述差值种最小的那一个，作为当前像素的滤波结果；

5. 对于第一个被滤波像素(x,y)，没有上一个被滤波像素，因此直接将像素值作为滤波结果。

此外，paeth.go文件还实现了paethPredictor结构体和paethPredictor的方法paethPredictor.Apply方法，这个方法可以对一个图片数据块（一个像素块）进行Paeth滤波处理，并返回处理后的结果。

## Functions:

### abs

在 Go 语言中，abs 是一个内置的函数，用于返回一个数的绝对值。

在 image 包中的 paeth.go 文件中，abs 函数被用于 Paeth 运算中计算两个像素值之间的距离。Paeth 运算是一种用于 PNG 图片压缩中的滤波算法，它通过选择距离目标像素最近的邻居像素来对目标像素进行预测，并将预测值与实际像素值进行比较，来判断是否需要进行编码，并计算编码后的像素值。

具体来说，Paeth 运算使用以下公式来计算两个像素值之间的距离：

    p := a + b - c
    pa := abs(p - a)
    pb := abs(p - b)
    pc := abs(p - c)

其中，a、b、c 分别为三个像素值，p 为预测值，pa、pb、pc 分别为预测值与实际值的偏差值。

在该公式中，abs 函数被用于计算偏差值，目的是保证偏差值为正数，从而保证计算准确性。因为在 Paeth 运算中，偏差值越大，说明预测值与实际值之间的距离越远，编码后的像素值就会更大，所以保证偏差值为正数非常重要。



### paeth

paeth函数实现了Paeth滤波器，它是PNG图像中的一种滤波器类型。滤波器可以让图像在经过压缩后得到更好的压缩比。

Paeth滤波器的原理是通过计算当前像素点的左侧、上侧和左上侧三个像素点的值来得出一个预测值，然后用预测值来减少原始图像数据，从而实现压缩。Paeth滤波器的计算公式如下：

paethPredictor = paeth(left, above, upperLeft)
current = originalValue - paethPredictor

其中，left表示当前像素点左侧的像素值，above表示当前像素点上方的像素值，upperLeft表示当前像素点左上方的像素值。paeth函数就是用来计算paethPredictor值的。

paeth函数的实现源代码如下：

func paeth(left, above, upperLeft uint8) uint8 {
	p := int(left) + int(above) - int(upperLeft)
	pa := abs(p - int(left))
	pb := abs(p - int(above))
	pc := abs(p - int(upperLeft))
	if pa <= pb && pa <= pc {
		return left
	}
	if pb <= pc {
		return above
	}
	return upperLeft
}

这个函数接受三个参数，分别是当前像素点左侧、上方和左上方的像素值。它根据上面的公式计算出paethPredictor值，并返回这个值。具体来说，它首先计算出p值，然后分别计算出pa、pb和pc三个值，并找到其中最小的那个值对应的像素值返回。

总结一下，paeth函数的作用就是计算Paeth滤波器中的paethPredictor值，帮助PNG图像压缩得更好。



### filterPaeth

在PNG图像中，每个扫描行（row）通常会经过一个压缩过程，其中一种压缩方式就是使用过滤器（filter）。而filterPaeth就是PNG图像中的一种过滤器类型。

该函数的作用是对一行像素数据进行Paeth滤波，通过使用前一行的像素值来计算出当前像素的值，从而达到压缩的效果。它使用的算法是一个类似于带权平均数的计算公式，具体实现如下：

- 当前像素的值 = 当前像素值 - 左侧相邻像素值 + 前一行当前列像素值；
- 提供三个编号为a、b、c的像素，Paeth滤波器计算出三个差值pa、pb、pc；
- 选择pa、pb、pc中绝对值最小的一个作为最终像素值，并将其加到像素中。

除了filterPaeth外，image/paeth.go文件中还包含其他的过滤器类型函数，如filterNone、filterSub、filterUp等。它们在不同的情况下应用不同的算法来压缩PNG图像数据。



