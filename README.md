# Vodka

`Vodka` is a lightweight abstraction over `net/http` to build HTTP services in Go. `Vodka` provides an easy way to build APIs, with good middleware support.

## Install

```go
go get github.com/Jamlie/vodka
```

## Examples

```go
package main

import (
    "log"
    "net/http"
    "strconv"

    "github.com/Jamlie/vodka"
    "github.com/Jamlie/vodka/middleware"
)

func main() {
    v := vodka.New()

    v.Use(middleware.Logger)
    v.Use(middleware.CORS)

    v.GET("/", func(c vodka.Context) { // /
        c.HTML(http.StatusOK, "<h1>Hello, Vodka</h1>")
    })

    v1 := v.Route("/api")

    v1.POST("/users/{id}", func(c vodka.Context) { // /api/users/{id}
        userId := c.Query("id")

        id, err := strconv.Atoi(userId)
        if err != nil {
            c.JSON(http.StatusBadRequest, map[string]string{
                "error": err.Error(),
            })
        }

        c.JSON(http.StatusOK, map[string]int{
            "id": id,
        })
    })

    log.Fatal(http.ListenAndServe(":8080", v)) // or log.Fatal(v.Start(":8080"))
}
```
