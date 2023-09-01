# model 数据库表生成model层

## roadmap

### todo

1. 设计转换 datatypes 定制化类型转换

## 设计

### 基础数据类型映射  

Db数据类型映射到Golang数据类型 ```conf/config.yaml / db_type_map```

### 数据对象空

参考文档: 

1. [SqlNull处理](https://iamdual.com/en/posts/handle-sql-null-golang/)
2. [golang 官方wiki](https://github.com/golang/go/wiki/SQLInterface)
3. [sql文档](https://pkg.go.dev/database/sql#ColumnType.Nullable)

config中 ```go_nullable_map``` 定义空类型映射关系

### sql 基础字段约定
1. tbl{content} tbl标识业务字段表
2. ORM常用 GROM 实现关系映射, GORM规范 gorm.Model 定义基础类型

```sql
CREATE TABLE `tblModelTable` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',

  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据表模型';
```

### comment 注释识别类型
通过 ```gorm.io/datatypes``` 扩展数据类型。 json 定义类型

1. ```datatypes.Date``` 日期类型，使用Sql NullTime 原始类型。
2. ```datatypes.RawMessage``` comment 包含json, 转换该类型
3. ```datatypes.Time``` comment 包含timestamp 转换
4. ```datatypes.URL``` 非空转换类型 


### 通过Dao 封装Context 上下文

1. ```type {XXX}Dao struct{ ctx Context }``` 定义context 上下文
2. ```New{XXX}Dao() *{XXX}Dao``` 创建实体对象上下文

### 函数
1. Update 实体对象修改。
2. Insert BatchInsert 添加批量添加函数
3. Delete 通过id 删除实体对象
4. QueryByIds 通过 Ids列表查询


## template 上下文定义

字段定义
```json
{
    "TableName":"tblModelTable",
    "BeanName":"ModelTable",
    "Columns":[
        {
            "ColumnName":"id",
            "FieldType":"int64",
            "Comment":"主键id"
        },
        {
            "ColumnName":"name",
            "FieldType":"string",
            "Comment":""
        }
    ],
    "Comment":"数据表模型"
}
```