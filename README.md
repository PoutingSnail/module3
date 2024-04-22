# Module3作业
|文件名|备注|
|:---|:---|
|main.go|go源文件|
|httpsvr|编译后的二进制文件|
|dockerfile|生成镜像配置文件|

### 构建本地镜像
  ```
  docker build  /root/httpsvr -t poutingsnail/httpsvr:v1.0
### 编写 Dockerfile 将模块二作业编写的 httpserver 容器化
``` FROM ubuntu
ENV VERSION=v1.0
ADD ./httpsvr /httpsvr
ENTRYPOINT /httpsvr
将镜像推送至 docker 官方镜像仓库      docker push poutingsnail/httpsvr:v1.0
通过 docker 命令本地启动 httpserver    docker run -d poutingsnail/httpsvr:v1.0
通过 nsenter 进入容器查看 IP 配置  
docker inspect containerid |grep -i pid
nsenter -t pid -n ip a
