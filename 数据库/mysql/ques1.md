# 慢SQL问题排查
-- 永久配置（my.cnf）
[mysqld]
slow_query_log = 1
slow_query_log_file = /var/log/mysql/mysql-slow.log
long_query_time = 1
[ ] 确认慢查询日志开启
[ ] 分析TOP 10慢SQL
[ ] 检查执行计划类型是否为ALL
[ ] 验证索引是否被实际使用
[ ] 检查是否存在临时表或文件排序
[ ] 评估是否需要分库分表

```

## 执行计划分析
```text
EXPLAIN SELECT * FROM users WHERE age > 20;
| 字段 | 重点关注项 | |---------------|-----------------------------------| 
| type | ALL（全表扫描）需优化为range/index | | key | 实际使用的索引 | | rows | 预估扫描行数 | | Extra | Using filesort/Using temporary需警惕 |
```

# 索引正常，sql无优化空间，如何提高查询速度
```text
主要优化方向优先级建议：
1. 先尝试读写分离 + 缓存（成本最低）
应用层 -> Redis缓存 -> 数据库（缓存未命中时）
缓存策略：
热点数据：主动预热（如商品详情）
缓存更新：写DB后立即更新缓存（Cache-Aside模式）
防击穿：BloomFilter拦截无效查询

2. 数据量持续增长再考虑分片
-- 按用户ID哈希分片（示例分为4库）
user_id % 4 = 0 -> db_0
user_id % 4 = 1 -> db_1
分片策略：

水平分片：按行拆分（适合单表数据量大）
垂直分库：按业务拆分（如订单库、用户库分离）

3. 最后评估分布式方案
当单机方案达到极限时：
分布式数据库：TiDB（HTAP架构）、OceanBase
搜索引擎：Elasticsearch（复杂条件查询）
预计算：物化视图/预聚合（如Kylin）

4. 列式存储
ClickHouse：高性能列式数据库
MySQL列存引擎：TokuDB（已弃用）/ MyRocks

5. 硬件升级
关键指标：

内存：确保足够innodb_buffer_pool_size（建议占物理内存70%）
SSD：替换机械硬盘（尤其对随机IO敏感场景）
CPU：多核优化（配置innodb_thread_concurrency）
```