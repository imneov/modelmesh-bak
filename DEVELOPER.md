# 开发手册

## 搭建开发环境

环境要求 ：
- go 1.18
- protoc 3.20.3
- MindSpore 2.0.0-cp37-cp37m-linux_x86_64

## 流程

生成 go 文件
```bash
make proto 
```


## Install 


## 参考库
https://github.com/yireyun/go-queue


https://github.com/joncrlsn/dque  
- 看起来不错
- 支持 内存以及 文件存储

https://github.com/iamduo/workq
- Workq is a job scheduling server strictly focused on simplifying job processing and streamlining coordination. It can run jobs in blocking foreground or non-blocking background mode.
- Workq runs as a standalone TCP server and implements a simple, text based protocol. Clients interact with Workq over a TCP socket in a request/response model with text commands. Please refer to the full protocol doc for details.
- 结构比较清晰




## Troubleshooting

### 1. ImportError: dlopen(//python3.9/site-packages/grpc/_cython/cygrpc.cpython-39-darwin.so, 0x0002): symbol not found in flat namespace '_CFRelease'

```bash
pip uninstall grpcio
conda install grpcio==1.47.0
```


