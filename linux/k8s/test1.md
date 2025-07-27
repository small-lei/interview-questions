# k8s一个pod创建的过程
一个 Pod 从创建到运行的大致流程如下：

1.用户提交 Pod 定义 → 2. API Server 接收请求并验证 → 3. Scheduler 选择节点 → 4. Kubelet 在节点上拉取镜像并启动容器 → 5. Kubelet 更新状态 → 6. Pod 运行

这些步骤的核心组件包括：API Server（接收和验证请求）、Scheduler（调度 Pod 到合适的节点）、Kubelet（在节点上启动和管理容器）、容器运行时（启动容器）、etcd（持久化状态）。通过这些组件的协作，Kubernetes 实现了自动化的容器编排和管理。


# kubernetes 的架构是怎么样的?这个问题很大，拆成 apiserver、controller、kubelet、scheduler 讲了一下
Kubernetes 的架构可以拆分为几个关键组件，每个组件在整个集群中负责不同的任务。以下是对 apiserver、controller、kubelet 和 scheduler 的详细讲解：

#### API Server (kube-apiserver)
功能:
API Server 是 Kubernetes 控制平面的核心组件，负责处理来自用户、外部组件、以及内部组件的 API 请求。
   作用:

   所有集群的操作都通过 API Server 进行，例如创建、更新、删除资源（如 Pods、Services 等）。
   它将所有请求转换为 REST API 调用，然后与 etcd 进行交互以存储集群的状态。
   API Server 是 Kubernetes 中的前端服务，所有管理和控制操作通过它来执行。
#### Controller Manager
   功能: Controller Manager 是 Kubernetes 中负责执行集群中控制循环的组件。控制循环的作用是监视集群状态，确保实际状态与期望状态保持一致。
   
作用:
   Node Controller: 监控节点的健康状况，如果节点失联，将会从调度中移除该节点的 Pod。

   Replication Controller: 确保指定数量的 Pod 副本保持运行。

   Endpoint Controller: 关联服务与 Pod。

   Job Controller: 确保 Job 任务顺利执行完成。

   其他还有很多不同的控制器，每个控制器负责集群中不同资源的状态维护。
#### Kubelet
   功能:
   Kubelet 是 Kubernetes 节点上的代理，负责确保在该节点上的容器正常运行。

   作用:
   
Kubelet 会从 API Server 获取调度到该节点的 Pod 说明，并确保这些容器正常运行。
   它定期报告节点的状态（如资源使用情况）回到控制平面。
   Kubelet 不直接管理容器，而是通过容器运行时（如 Docker、containerd）启动和管理容器。
#### Scheduler
   功能:
   Scheduler 负责为未指定运行节点的 Pod 选择合适的节点，并将调度决定反馈给 API Server。
   
作用:

   Scheduler 会分析每个 Pod 的资源需求和约束条件（如 CPU、内存、节点标签等），根据集群当前的资源可用情况选择最适合的节点来运行 Pod。
   它会根据多种调度策略（如负载均衡、资源限制、优先级等）来做出决策，保证集群资源的最佳利用。

# 怎么实时查看k8s内存占用的

在 Kubernetes 中，实时查看内存占用有多种方式，以下是几种常见的做法：
#### 使用 kubectl top 命令
   kubectl top 是 Kubernetes 自带的一个命令，可以显示当前节点和 Pod 的资源使用情况（包括 CPU 和内存）。要使用它，集群中需要安装 metrics-server 组件，它负责收集和提供资源使用数据。
安装 metrics-server
如果集群中还没有 metrics-server，可以通过以下方式安装：
```text
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```
查看节点的内存使用
使用以下命令可以查看所有节点的内存和 CPU 使用情况：
```text
kubectl top nodes

NAME               CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%
node-1             100m         5%     500Mi           25%
node-2             250m         10%    1Gi             50%
```
查看 Pod 的内存使用
要查看某个命名空间下的 Pod 的资源使用情况，可以使用以下命令：
```text
kubectl top pods -n default
NAME                       CPU(cores)   MEMORY(bytes)
nginx-1                    10m          50Mi
nginx-2                    15m          60Mi
```

#### 使用 kubectl describe 命令
   虽然 kubectl describe 主要用来查看 Pod 或节点的详细信息，但也可以通过查看内存请求（requests）和限制（limits）来了解 Pod 的资源使用情况：
   
```text
这会返回该 Pod 的详细信息，其中包括内存的请求和限制。虽然这不是实时内存使用，但可以用来参考当前配置的资源限制。
```

#### 使用 Prometheus + Grafana
#### 使用 Lens
#### 使用 cadvisor

#### 总结
简单实时查看：使用 kubectl top 命令查看当前的节点和 Pod 的内存占用。

更复杂的监控：使用 Prometheus + Grafana 提供历史和实时监控数据。

图形化管理：使用 Lens 或 cAdvisor 实时查看内存、CPU 等资源使用情况。