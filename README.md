## Middlewares in Beego


Basic example middleware for Beego framework.

### Installation

Standard `go get`:

```
$ go get github.com/ojardila/beego_middleware
```

#### Dependencies

```
  go get github.com/astaxie/beego 
```

#### Configuring Beego Middleware

Add the following lines into the ```routers/routers.go``` file which will initialize the filter to run on all requests (BeforeStatic, BeforeRouter, BeforeStatic, AfterExec and FinishRouter)


```go
 import "github.com/ojardila/beego_middleware"
 
 func init() {
    beego_keenio.InitMiddleware()

  
 }
```
