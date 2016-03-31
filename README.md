# ctxgrp

A route group + middle ware library that integrates with any http router.

Usage:
`go get github.com/peak6/ctxgrp/httprouter`  

(work in progress example) 

```go
import (
  "net/http"
  "github.com/peak6/ctxgrp"
  "github.com/peak6/ctxgrp/httprouter"
)

func main(){
  root := httprouter.New()
  api := root.NewGroup("/api").Use(middleware)
  other := root.NewGroup("/other")
  
  api.GET("/foo", fooHandler)
  other.GET("/bar", barHandler)
}

func middleware(h ctxgrp.Handler) ctxgrp.HandlerFunc{
  return func(ctx context.Context, w http.ResponseWriter, r *http.Request){
    log.Println("middleware!")
    h.ServeHTTP(ctx,w,r)
  }
}
func fooHandler(ctx context.Context, w http.ResponseWriter, r *http.Request){
  w.Write([]byte("foo handler"))
}

func otherHandler(ctx context.Context, w http.ResponseWriter, r *http.Request){
  w.Write([]byte("other handler"))
}
```
