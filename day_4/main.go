package main

func main() {
	/*
		从业务层和语言运行时层进行优化分别有什么特点？
			业务层的优化需要针对具体问题具体分析；
			而语言运行时层的优化针对的是更通用的问题，需要在多方面进行 tradeoff。两者都需要自动化性能分析工具的支持。

		从软件工程的角度出发，为了保证语言SDK的可维护性和可拓展性，在进行运行时优化时需要注意什么？
			保证优化后代码的行为与优化前的行为相同，提供完善的文档说明，对于功能的改动，最好通过选项进行隔离，保证新增改动在不打开的情况下不影响原本的功能。

		自动内存管理技术从大类上分为哪两种，每一种技术的特点以及优缺点有哪些？
			追踪垃圾回收:
				回收内存的条件：回收不可达的对象
				回收时，首先会扫描 GC roots，例如栈上的对象、全局变量等；
				从 GC roots 出发，沿着指针指向，追踪所有的可达对象；
				追踪完成后，回收所有不可达对象的内存。
			引用计数:
				对象有一个与之关联的引用数目。当且仅当引用数大于 0 时对象存活；
				当对象的引用数目为 0 时，内存可以被回收；
				优点：
					内存管理的操作被平摊到程序运行中：指针传递的过程中进行引用计数的增减。
					不需要了解 runtime 的细节：因为不需要标记 GC roots，因此不需要知道哪里是全局变量、线程栈等
				缺点：
					开销大，因为对象可能会被多线程访问，对引用计数的修改需要原子操作保证原子性和可见性
					无法回收环形数据结构
					每个对象都引入额外存储空间存储引用计数
					虽然引用计数的操作被平摊到程序运行过程中，但是回收大的数据结构依然可能引发暂停

		什么是分代假说？分代 GC 的初衷是为了解决什么样的问题？
			大量的对象会很快死去。分代 GC 的思想是：针对不同生命周期的对象采取不同策略的内存管理机制。

		Go 是如何管理和组织内存的？
			Go 使用 TCMalloc 风格的内存管理方式。
				a.  TC 是 thread caching 的简写。每个线程都绑定 cache 方便线程快速分配内存；
				b.  内存被划分为特定大小的块。根据对象是否包含指针，将内存块分为 scan 和 noscan 两种；
				c.  根据内存分配请求的大小，选择合适的内存块返回，完成一次内存分配操作；
				d.  回收的内存不会立刻还给操作系统，而是在 Go 内部缓存起来，方便下次分配。

		为什么采用 bump-pointer 的方式分配内存会很快？
			每个线程都持有用于对象分配的 buffer，因此指针碰撞方式的内存分配无需加锁或使用 CAS 操作；对象分配的操作非常简单。

		为什么我们需要在编译器优化中进行静态代码分析？
			通过静态分析，我们可以获取更多关于程序的非平凡特性 (non-trivial properties)；
			这些关于程序的知识可以指导编译器优化。
			例如通过逃逸分析得知对象并未逃逸出当前函数，因此对象可以在栈上分配，避免频繁在堆上分配对象，降低 GC 的压力。

		函数内联是什么，这项优化的优缺点是什么？
			函数内联：将被调用函数的函数体的副本替换到调用位置上，同时重写代码以反映参数的绑定。
			a. 优点
				① 消除函数调用；
				② 由于没有了函数调用，过程间分析转化为过程内分析；
			b. 缺点
				① 函数体变大；
				② 编译生成的 Go 镜像变大。

		什么是逃逸分析？逃逸分析是如何提升代码性能的？
			逃逸分析：分析代码中指针的动态作用域，即指针在何处可以被访问。
			通过逃逸分析得知对象并未逃逸出当前函数，因此对象可以在栈上分配，避免频繁在堆上分配对象，降低 GC 的压力。
	*/
}
