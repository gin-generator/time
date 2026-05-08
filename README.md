# Time

A lightweight Go package that wraps `time.Time` with custom JSON serialization and database scanning support.

[中文文档](README_ZH.md)

## Features

- **JSON Serialization**: Automatically formats time as `"2006-01-02 15:04:05"` (DateTime format)
- **Database Support**: Implements `sql.Scanner` and `driver.Valuer` interfaces
- **Zero Value Handling**: Returns `NULL` for zero time values in database operations
- **Local Timezone**: Uses local timezone for JSON unmarshaling

## Installation

```bash
go get github.com/gin-generator/time
```

## Usage

### Basic Usage

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
    CreatedAt _time.Time `json:"created_at"`
}

func main() {
    event := Event{
        Name:      "Meeting",
        CreatedAt: _time.Time(time.Now()),
    }

    // JSON Marshal
    data, _ := json.Marshal(event)
    fmt.Println(string(data))
    // Output: {"name":"Meeting","created_at":"2026-05-08 10:30:45"}

    // JSON Unmarshal
    var newEvent Event
    json.Unmarshal(data, &newEvent)
}
```

### Database Usage

```go
type User struct {
    ID        int64           `db:"id"`
    Name      string          `db:"name"`
    CreatedAt customtime.Time `db:"created_at"`
}

// Works with database/sql
db.Query("SELECT id, name, created_at FROM users")
```

## API

#### `MarshalJSON() ([]byte, error)`
Serializes time to JSON string in `"2006-01-02 15:04:05"` format.

#### `UnmarshalJSON(data []byte) error`
Deserializes JSON string to time using local timezone.

#### `String() string`
Returns time formatted as `"2006-01-02 15:04:05"`.

#### `Value() (driver.Value, error)`
Implements `driver.Valuer` for database operations. Returns `NULL` for zero values.

#### `Scan(v any) error`
Implements `sql.Scanner` for reading from database.

## License

MIT
