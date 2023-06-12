# mall-demo
采用Golang、gin、gorm、mysql、redis、rabbitmq实现的单机秒杀商城后端demo

| 文件                     | 说明           |
|------------------------|--------------|
| api/v1                 | 接口文件         |
| cmd                    | main函数程序启动入口 |
| config                 | 配置文件         |
| internal               | 程序主体         |
| internal/cache         | redis缓存      |
| internal/dao           | 数据持久层        |
| internal/dto           | 数据传输层        |
| internal/initialize    | 初始化mysql文件   |
| internal/middleware    | 中间件          |
| internal/model         | 数据库实体        |
| internal/pkg           | 工具类          |
| internal/pkg/e         | 响应码          |
| internal/pkg/res       | 响应类          |
| internal/pkg/util      | 工具类          |
| internal/rabbitMQ      | 消息队列         |
| internal/service       | 业务逻辑层        |
| internal/vo            | 视图层          |
| rabbitMQServer         | 消息队列服务       |
| rabbitMQServer/cmd     | 服务启动入口       | 
| rabbitMQServer/service | 消费者业务逻辑层     |
| router                 | 路由           |
| go.mod                 | 依赖           |


>使用多线程模拟并发秒杀，简单实现了四种方式：
>1. 无锁秒杀，会出现超卖
>2. 使用mutex互斥锁秒杀，正常，但效率很低
>3. 使用for update排他锁秒杀，正常
>4. 将商品信息提前存入redis，使用lua脚本实现检查库存、减库存、保存用户信息，秒杀成功将订单信息通过MQ发布，消费者监听消息并写入数据库。


>事务隔离级别采用RC,不加锁采用Serializable级别会发生死锁，原因：
> * 在串行化隔离级别下，若有两个线程A和B
> * A在查库存时获取读锁
> * B在查库存时也获取读锁
> * A在查完库存后想要更新库存，获取写锁
> * B在查完库存后也想要更新库存，B等待A释放写锁
> * 但串行化事务中，读锁被获取后，写锁需要等待
> * 造成A等待B释放读锁，B等待A释放写锁
> * 循环等待，造成死锁


>TODO：
>1. 主从数据库
>2. 搭建redis分布式集群
>3. 微服务实现