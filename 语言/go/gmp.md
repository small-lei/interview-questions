

````text
Go 的 GMP 模型是什么？
G：Goroutine；

M：Machine，操作系统线程；

P：Processor，调度器调度队列（维护可执行G）；

一个P绑定一个M，运行G；

Go运行时负责调度G->M->P；

支持 work stealing、G 的复用等优化。
````