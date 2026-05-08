# Time

一个轻量级的 Go 时间包装库，提供自定义 JSON 序列化和数据库扫描支持。

[English](README.md)

## 特性

- **JSON 序列化**：自动格式化时间为 `"2006-01-02 15:04:05"` 格式
- **数据库支持**：实现了 `sql.Scanner` 和 `driver.Valuer` 接口
- **零值处理**：数据库操作中零值时间返回 `NULL`
- **本地时区**：JSON 反序列化使用本地时区

## 安装

```bash
go get github.com/gin-generator/time
```

## 使用方法

### 基础用法

```go
package main

import (
    "encoding/json"
    "fmt"
	"time"
    _time "github.com/gin-generator/time"
)

type Event struct {
    Name      string          `json:"name"`
    CreatedAt _time.Time      `json:"created_at"`
}

func main() {
    event := Event{
        Name:      "会议",
        CreatedAt: _time.Time(time.Now()),
    }

    // JSON 序列化
    data, _ := json.Marshal(event)
    fmt.Println(string(data))
    // 输出: {"name":"会议","created_at":"2026-05-08 10:30:45"}

    // JSON 反序列化
    var newEvent Event
    json.Unmarshal(data, &newEvent)
}
```

### 数据库用法

```go
type User struct {
    ID        int64           `db:"id"`
    Name      string          `db:"name"`
    CreatedAt _time.Time      `db:"created_at"`
}

// 配合 database/sql 使用
db.Query("SELECT id, name, created_at FROM users")
```

## API 说明

#### `MarshalJSON() ([]byte, error)`
将时间序列化为 `"2006-01-02 15:04:05"` 格式的 JSON 字符串。

#### `UnmarshalJSON(data []byte) error`
使用本地时区将 JSON 字符串反序列化为时间。

#### `String() string`
返回 `"2006-01-02 15:04:05"` 格式的时间字符串。

#### `Value() (driver.Value, error)`
实现 `driver.Valuer` 接口用于数据库操作。零值时间返回 `NULL`。

#### `Scan(v any) error`
实现 `sql.Scanner` 接口用于从数据库读取数据。

## 许可证

MIT
