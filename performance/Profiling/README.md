#### Go Performance  
> https://mp.weixin.qq.com/s/mTHBJrseyWmEBxgPaqQJtw
#### 可以查看 CPU 占用情况：
>  go tool pprof -http=:9090 http://localhost:6060/debug/pprof/profile

####  查看内存分配情况：
> go tool pprof -http=:9091 http://localhost:6060/debug/pprof/allocs