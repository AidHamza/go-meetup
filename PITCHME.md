# Go Web Server!
## @AidHamza

---

## $ whoami

<img src="07-10-2017/assets/gopher.png" height="230px" />

* @AidHamza
* Software engineer @Schibsted Media Group

---

## References

https://github.com/AidHamza/go-meetup

* Previous Presentations
* Demo Source Code

---

## Go Web Server

+++

### net/http Package

Freaking Simple!

``` go
http.HandleFunc("/", helloGophers)
http.ListenAndServe(":8080", nil)
```

@[1](Handle / request by helloGophers func)
@[2](Run the HTTP server on port 8080)

+++

### net/http Handler interface
Once implemented the `http.Handler` can be passed to `http.ListenAndServe`, which will call the `ServeHTTP` method on every incoming request.

```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```

+++

#### Example

``` go
package main

import (
  "fmt"
  "log"
  "net/http"
)

type helloHandler struct{}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
}

func main() {
  err := http.ListenAndServe(":9999", helloHandler{})
  log.Fatal(err)
}
```

+++

### HTTP request multiplexer

It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

Patterns may optionally begin with a host name, restricting matches to URLs on that host only. Host-specific patterns take precedence over general patterns, so that a handler might register for the two patterns "/codesearch" and "codesearch.google.com/" without also taking over requests for "http://www.google.com/".

Check : https://caddyserver.com/

+++

#### Example :

``` go
package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  h := http.NewServeMux()

  h.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, you hit the blog page!")
  })

  h.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, you hit contact us!")
  })

  h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(404)
    fmt.Fprintln(w, "You're lost, go home")
  })

  err := http.ListenAndServe(":8080", h)
  log.Fatal(err)
}
```

+++

#### Other mux Packages

There are numerous replacements for http.ServeMux like gorilla/mux which give you things like automatically pulling variables out of paths, easily asserting what http methods are allowed on an endpoint, and more.

+++

#### Composability

http package is composable, it is very easy to create re-usable pieces of code and glue them together into a new working application. The http.Handler interface is the way all pieces communicate with each other.

+++

### Middlware

Serving endpoints is nice, but often there’s functionality you need to run for every request before the actual endpoint’s handler is run. For example, access logging. A middleware component is one which implements http.Handler, but will actually pass the request off to another http.Handler after doing some set of actions.

---

## Go Templates

Most of the modern programming languages implements a templating package / library; Golang has a very powerful template package.

Templates are executed by applying them to a data structure. Annotations in the template refer to elements of the data structure (typically a field of a struct or a key in a map) to control execution and derive values to be displayed.

Reference : https://golang.org/pkg/text/template/

+++

### Example

Let's demo an email template example

---

## Go Routines to create Worker Pools

In this chapter, we will discover how to do concurrent tasks the right way.
Handling and running Go Routines is not an easy task, its complex how to manage the routine state, the resources you have, and syncronize between those routines.

+++

### Thread Pool

Thread pool is a design pattern aiming to achieve tasks concurrency execution in our programs.

a Thread Pool contains multiple threads waiting to execute tasks

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/0c/Thread_pool.svg/580px-Thread_pool.svg.png">

+++

#### Thread pool in Go

Worker pools are a model in which a fixed number of m workers (implemented in Go with goroutines) work there way through n tasks in a work queue (implemented in Go with a channel). Work stays in a queue until a worker finishes up its current task and pulls a new one off.


+++

#### Go workers pool is simple!

``` go
func worker(id int, jobs <-chan int, results chan<- int) {
  for j := range jobs {
    fmt.Println("worker", id, "processing job", j)
    time.Sleep(time.Second)
    results <- j * 2
  }
}

func main() {
  jobs := make(chan int, 100)
  results := make(chan int, 100)

  for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
  }

  for j := 1; j <= 9; j++ {
    jobs <- j
  }
  close(jobs)

  for a := 1; a <= 9; a++ {
    <-results
  }
}
```

---



## Any question before demo ?



---

## Let's get hands dirty

---

## Raffle time with Jetbrains Love
Thanks Jetbrains for the continious support, and helping many communities growing.

<img src="https://confluence.jetbrains.com/download/attachments/37945345/global.logo?version=5&modificationDate=1449748383000&api=v2" />

---

## Avito.ma

I would like to thank Avito.ma for the community support, and trying to provide all the possible help to the tech communities in Morocco.

<img src="https://www.avito.ma/img/logo-avito-og.png" />

Work with us ?
https://www.avito.ma/emploi

---

## Thank you!