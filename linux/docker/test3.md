# 如何编写 Dockerfile？列出常用的命令

编写一个 Dockerfile 时，你需要按照一系列步骤定义如何构建 Docker 镜像。Dockerfile 包含了构建镜像所需的指令和配置。

Dockerfile 常用命令
FROM
```text
说明：指定基础镜像，所有指令都基于此镜像执行。
示例：
dockerfile
复制代码
FROM ubuntu:20.04
```

RUN
```text
说明：在构建镜像时执行的命令，通常用于安装软件包、更新系统等。
示例：
dockerfile
复制代码
RUN apt-get update && apt-get install -y nginx
```

COPY
```text
说明：将主机文件系统中的文件或目录复制到镜像中。
示例：
dockerfile
复制代码
COPY ./src /app/src
```

ADD
```text
说明：类似于 COPY，但可以解压归档文件，支持 URL 下载。
示例：
dockerfile
复制代码
ADD https://example.com/file.tar.gz /app/file.tar.gz
```

CMD
```text
说明：指定容器启动时运行的默认命令，可以被 docker run 时的命令覆盖。
示例：
dockerfile
复制代码
CMD ["nginx", "-g", "daemon off;"]
```

ENTRYPOINT
```text
说明：设置容器启动时的命令，不能被覆盖，常用于指定主进程。
示例：
dockerfile
复制代码
ENTRYPOINT ["nginx"]
```

WORKDIR
```text
说明：设置工作目录，后续的指令将在此目录中执行。
示例：
dockerfile
复制代码
WORKDIR /app
```

ENV
```text
说明：设置环境变量。
示例：
dockerfile
复制代码
ENV NODE_ENV=production
```

EXPOSE
```text
说明：声明容器要监听的端口（仅声明作用，无实际端口绑定）。
示例：
dockerfile
复制代码
EXPOSE 80
```

VOLUME
```text
说明：声明挂载点，可以将容器中的某个目录映射为持久化卷。
示例：
dockerfile
复制代码
VOLUME /app/data
```

ARG
```text
说明：定义在构建时传递的变量，与 ENV 不同，它只在构建阶段有效。
示例：
dockerfile
复制代码
ARG build_version
```

USER
```text
说明：指定容器内的用户权限，通常用于避免以 root 权限运行。
示例：
dockerfile
复制代码
USER appuser
```

Dockerfile 示例
```text
dockerfile
复制代码
# 使用基础镜像
FROM node:16-alpine

# 设置工作目录
WORKDIR /app

# 复制 package.json 文件
COPY package.json .

# 安装依赖
RUN npm install

# 复制应用源代码
COPY . .

# 声明环境变量
ENV NODE_ENV=production

# 暴露应用监听的端口
EXPOSE 3000

# 设置容器启动时执行的命令
CMD ["npm", "start"]
这个示例 Dockerfile 会创建一个基于 Node.js 的应用环境，安装依赖，复制代码，并设置默认启动命令。

```

### 总结
常用命令如 FROM, RUN, COPY, CMD 等，定义了如何构建镜像和启动容器。通过这些命令的组合，你可以创建可移植的容器化应用程序。