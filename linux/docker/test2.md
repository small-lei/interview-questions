# 列出至少 6 个常用的 Docker 命令
以下是 6 个常用的 Docker 命令，涵盖镜像管理、容器操作、网络管理等：

docker build
用于从 Dockerfile 构建镜像。
```text
bash
复制代码
docker build -t my-image:latest .
选项：
-t：指定镜像标签（tag）。
.：当前目录下的 Dockerfile。
```

docker run
创建并启动一个新的容器。
```text
bash
复制代码
docker run -d -p 8080:80 my-image
选项：
-d：后台运行容器（detached）。
-p：端口映射，将主机端口映射到容器端口。
```

docker ps
查看正在运行的容器列表。
```text
bash
复制代码
docker ps
选项：
-a：显示所有容器，包括已停止的容器。
```

docker exec
在运行的容器内执行命令。
```text
bash
复制代码
docker exec -it my-container /bin/bash
选项：
-it：交互模式打开容器终端。
```

docker images
列出本地所有 Docker 镜像。
```text
bash
复制代码
docker images
```

docker logs
查看容器的日志。
```text
bash
复制代码
docker logs my-container
选项：
-f：实时跟踪日志输出。
```

这些命令涵盖了 Docker 容器和镜像的基本操作，是日常使用 Docker 时最常见的命令。