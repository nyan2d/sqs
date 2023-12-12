# SQS

## WTF

A simple golang session storage that has no external dependencies and uses *std sql.DB* interface.

## How to install

```
go get -u github.com/nyan2d/sqs
```

## How to add sqs to your project

```go
import "github.com/nyan2d/sqs"
```

## Usage

```go
package main

import (
    "database/sql"

    "github.com/nyan2d/sqs"
    _"modernc.org/sqlite"
)

func main() {
    db, _ := sql.Open("sqlite", "database.db")
    ss, _ := sqs.NewSessionStorage(db, "sessions")

    userID := 35
    validTo := time.Now().Add(12*time.Hour)

    token, _ := ss.AddSession(userID, validTo)
}
```

The code above will give you a token that you can, for example, store in a cookies.

*to be continued*