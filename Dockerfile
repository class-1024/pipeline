
#https://blog.csdn.net/niyuelin1990/article/details/79035728
#源镜像
FROM golang:latest
#作者
MAINTAINER Razil "1965948351@163.com"
#设置工作目录
WORKDIR $GOPATH/src/github.com/kinot-web
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/github.com/mygohttp
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 80
#最终运行docker的命令
ENTRYPOINT  ["./kinot-web"]
