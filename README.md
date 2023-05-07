# OPS

### 技术选型的思考

1. viper能完成修改配置文件自动生效，就不在使用配置中心，如果后期需要上Nacos或者是Apollo，分布式配置中心介绍 https://www.cnblogs.com/liuqingzheng/p/16322508.html
2. 不在使用传统mysql的分库分表，而是借助tidb的自动分区来实现，数据库使用tidb
3. 目前虽然是单服务，不排除后面扩展，引入网关中心kong，介绍 https://www.cnblogs.com/liuqingzheng/p/16367226.htm
4. 由于gin框架本身不支持热加载，引入air