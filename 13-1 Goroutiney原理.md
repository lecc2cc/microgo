### Go语言实践 - Runtime

+ Goroutine原理
+ 内存分配原理
+ GC原理
+ Channel原理

#### Goroutine 原理

Goroutine 是一个与其他 goroutines 并行运行在同一地址空间的 Go 函数或方法。

一个运行的程序由一个或更多个 goroutine 组成。它与线程、协程、进程等不同。



Goroutines 在同一个用户地址空间里并行独立执行 functions，channels 则用于 goroutines 间的通信和同步访问控制。

