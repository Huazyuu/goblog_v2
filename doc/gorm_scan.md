在GORM中，`Scan`方法用于将查询结果映射到对应的结构体，下面详细介绍其机制以及结合你提供的代码进行分析。

### 1. `Scan`方法的工作机制

#### 1.1 字段映射规则
- **名称匹配**：GORM会尝试将查询结果中的列名与目标结构体的字段名进行匹配。匹配时，默认情况下不区分大小写，并且会忽略列名中的下划线（`_`）。例如，查询结果中的列名 `send_user_id` 会匹配结构体中的 `SendUserID` 字段，因为GORM会将列名转换为驼峰命名法后进行匹配。
- **类型匹配**：GORM还会检查查询结果的列数据类型与结构体字段的数据类型是否兼容。如果不兼容，可能会导致数据转换错误。例如，如果查询结果中的列是字符串类型，而结构体字段是整数类型，GORM会尝试进行类型转换，如果转换失败则会报错。

#### 1.2 扫描过程
- **查询执行**：当调用 `Scan` 方法时，GORM会先执行前面构建的SQL查询语句。在你的代码中，查询语句是 `SELECT send_user_id, count(distinct rev_user_id) as count FROM `message` WHERE send_user_nick_name = '测试用户' GROUP BY `send_user_id` LIMIT 10`。
- **结果集处理**：GORM会遍历查询结果集的每一行。
- **映射填充**：对于每一行结果，GORM会根据字段映射规则将列值填充到目标结构体的相应字段中。如果查询结果有多行，GORM会为每一行创建一个新的结构体实例，并将这些实例添加到切片中。

### 2. 结合你的代码分析

#### 2.1 结构体定义
```go
type MessageUserStat struct {
    SendUserID uint
    RevUserID  uint
    Count      int
}
```
这个结构体定义了三个字段，分别是 `SendUserID`、`RevUserID` 和 `Count`，用于存储查询结果。

#### 2.2 查询构建与 `Scan` 调用
```go
err = query.Select("send_user_id, count(distinct rev_user_id) as count").
    Group("send_user_id").
    Offset(offset).
    Limit(limit).
    Scan(&stats).Error
```
- **查询构建**：通过 `Select` 方法指定要查询的列，`Group` 方法进行分组，`Offset` 和 `Limit` 方法进行分页。
- **`Scan` 调用**：将查询结果扫描到 `stats` 切片中。GORM会将查询结果中的 `send_user_id` 列的值填充到 `MessageUserStat` 结构体的 `SendUserID` 字段中，将 `count` 列的值填充到 `Count` 字段中。由于查询结果中没有 `rev_user_id` 列，所以 `RevUserID` 字段的值将保持为其零值（对于 `uint` 类型，零值是 `0`）。

### 3. 示例代码验证
```go
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type MessageUserStat struct {
    SendUserID uint
    RevUserID  uint
    Count      int
}

// 模拟全局数据库连接
var global = struct {
    DB *gorm.DB
}{}

func init() {
    var err error
    global.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
}

// GetMessageUserStats 根据用户id聚合,查看对应昵称的用户有几条消息
func GetMessageUserStats(nickname string, page, limit int) (stats []MessageUserStat, err error) {
    offset := (page - 1) * limit
    query := global.DB.Model(&MessageUserStat{})
    if nickname != "" {
        query = query.Where("send_user_nick_name = ?", nickname)
    }

    err = query.Select("send_user_id, count(distinct rev_user_id) as count").
        Group("send_user_id").
        Offset(offset).
        Limit(limit).
        Scan(&stats).Error
    return stats, err
}

func main() {
    stats, err := GetMessageUserStats("测试用户", 1, 10)
    if err != nil {
        panic(err)
    }
    for _, stat := range stats {
        println(stat.SendUserID, stat.Count)
    }
}
```

在这个示例中，我们模拟了一个数据库连接，并调用 `GetMessageUserStats` 函数进行查询。查询结果会被扫描到 `stats` 切片中，然后遍历切片打印每个 `MessageUserStat` 结构体的 `SendUserID` 和 `Count` 字段的值。