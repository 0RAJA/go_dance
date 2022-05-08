package main

func main() {
	// 	1.了解下其他语言的编码规范，是否和 Go 语言编码规范有相通之处，注重理解哪些共同点？
	//	文件：小写+下划线
	//	变量：采用驼峰命名法
	//	常量：大写+下划线
	//	注释：可以通过 /* …… */ // ……增加注释， //之后应该加一个空格 ，// 注释内容需要在文件/方法/变量上方
	//	当接受者类型是一个结构体并且很庞大，或者是一个大数组，建议使用指针传递来提高性能，其他场景使用值传递即可
	//

	/*
		2.编码规范或者性能优化建议大部分是通用的，有没有方式能够自动化对代码进行检测？
			使用 gofmt 工具格式化代码。
			使用 goimports 工具检查导入。
			使用 golint 工具检查代码规范。推荐golangci-lint
			使用 go vet 工具静态分析代码实现。
			可以将系列检查操作集成到CI里，在每次提交代码时进行检查。
	*/
	/*
		3.从 https://github.com/golang/go/tree/master/src 中选择感兴趣的包，看看官方代码是如何编写的？
			看了下sort包的编写，首先的感受就是注释非常多，并且有十分全面的单元测试，性能测试和示例测试。
			而且命名也非常规范，接口设计也十分合理
	*/
	/*
		4.使用 Go 进行并发编程时有哪些性能陷阱或者优化手段？
			1. 不控制goroutine数量引发CPU性能问题：使用缓冲channel，sync，或者协程池解决
			2. 协程泄露问题，比如超时问题，或者是在消费者停止消费后忘记通知生产者等导致的泄露问题。
			3. 对于临时变量可以使用sync.Pool进行复用，减小GC压力
	*/
	/*
		5. 在真实的线上环境中，每个场景或者服务遇到的性能问题也是各种各样，搜索下知名公司的官方公众号或者博客，里面有哪些性能优化的案例？
			https://blog.csdn.net/g6U8W7p06dCO99fQ3/article/details/122646325
			Uber的Go语言GC优化
	*/
	/*
		6. Go 语言本身在持续更新迭代，每个版本在性能上有哪些重要的优化点？
			我认为go的垃圾回收机制的演进就十分明显
			1. Go 1.0 完全串行的标记和清除 需要暂停整个程序 几百ms
			2. Go 1.1 多核主机上并行执行垃圾收集的标记和清除
			3. Go 1.3 运行时增加了栈内存的精准扫描支持
			4。Go 1.5 实现基于三色标记清除的并发垃圾回收器
				垃圾收集延迟降低为10ms以下
			5. Go 1.6 实现去中心化的垃圾收集器
			6。Go 1.7 使用并行栈搜索
				延迟降低为2ms以内
			7. Go 1.8 使用混合写屏障
				延迟降低为0.5ms以内
			......
	*/
}
