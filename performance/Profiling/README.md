#### Go Performance  
> https://mp.weixin.qq.com/s/mTHBJrseyWmEBxgPaqQJtw
#### 可以查看 CPU 占用情况：
> go tool pprof -http=:9090 http://localhost:6060/debug/pprof/profile
> go tool pprof http://localhost:10000/debug/pprof/profile?seconds=60

####  查看内存分配情况：
> go tool pprof -http=:9091 http://localhost:6060/debug/pprof/allocs

#### 生成prof文件 执行上面的命令会自动保存文件到相应目录
> Saved profile in C:\Users\Admin\pprof\go tool pprof pprof.samples.cpu.002.pb.gz
> go tool pprof pprof.samples.cpu.002.pb.gz

#### 根据prof文件查看
> go tool pprof -http=":8081" .\cpu.prof
> go tool pprof -http=":8081" pprof.samples.cpu.002.pb.gz


#### 参考 https://zhuanlan.zhihu.com/p/396363069