# Unirest for Go

[![License][license-image]][license-url]

![][unirest-logo]


[Unirest](http://unirest.io) is a lightweight HTTP libraries designed by [Kong](https://github.com/Mashape/kong). This is the **unofficial Go implementation** of the Unirest protocol.

## Features

* Make `GET`, `POST`, `PUT`, `PATCH`, `DELETE` requests
* Both synchronous and asynchronous requests
* Supports form parameters and file uploads
* Supports gzip
* Supports Basic Authentication
* Customizable timeout
* Customizable default headers for every request (DRY)

## Installing

Requirements: **Go >= 1.11** (since it uses go modules)

To utilize unirest, install the `unirest` gem:

```bash
go get github.com/palash25/unirest-go
```

After installing the packageyou can now begin simplifying requests using:

```go
import "github.com/palash25/unirest-go"
```

## Creating Requests

So you're probably wondering how using Unirest makes creating requests in Ruby easier, let's start with a working example:

```go
import (
    "fmt"

    "github.com/palash25/unirest-go"
)

func synchronousPost() {
  c := unirest.NewClient()

  auth := map[string]string{"user": "username", "password": "password"}
  
  params := make(map[string]interface{})
  params["key"] = "value"

  headers := make(map[string]interface{})
  headers["Accept"] = "application/json"
  
  response, err := c.Post("http://httpbin.org/post", headers, params, auth).Do()
  if err != nil {
    panic(err)
  }
  
  fmt.Println(response.Code)
  fmt.Println(response.RawBody)
  fmt.Println(response.Headers)
}
```

## Asynchronous Requests Pool
Unirest-Go also supports asynchronous requests using two channels to feed the requests and obtain the responses:

```go
func asyncGetPool() {
    go func() {

    urls := []string{
        "http://httpbin.org/get",
        "http://httpbin.org/ip",
        "http://httpbin.org/headers",
    }

    for _, u := range urls{
      req := c.Get(u, body, nil, nil)
      // Feed the requests to the input channel
      c.InChan <- req
    }
    close(c.InChan)
  }()
  
  // Run the client in Async mode
  c.DoAsync()

  for i:=0; i<3; i++ {
    // retrieve the requests from the output channel
    r := <-c.OutChan

    // Asynchronous reponses are returned as a struct of the
    // HTTP response and error (if any)
    fmt.Println(r.Resp.Code, r.Err)
  }
}
```

## File Uploads
```go
params := make(map[string]interface{})
params["file"] = "/path/to/file.go"
response, err := c.Post("http://httpbin.org/response-headers", nil, params, nil).Do()
```

### Basic Authentication

Authenticating the request with basic authentication can be done by providing an `auth` map with `user` and `password` keys like:

```go
auth := map[string]string{"user": "username", "password": "password"}
response, err := c.Get("http://httpbin.org/get", nil, nil, auth).Do()
```

# Request
```go
client := unirest.NewClient()

client.Get(url, headers, params, auth)
client.Post(url, headers, params, auth)
client.Delete(url, headers, params, auth)
client.Put(url, headers, params, auth)
client.Patch(url, headers, params, auth)
```
  
- `url` (`String`) - Endpoint, address, or uri to be acted upon and requested information from.
- `headers` (`map[string]interface{}`) - Request Headers as associative array or object
- `parameters` (`interface{}`) - Request Body associative array or object
- `auth` (`map[string]string`) - Basic Auth credentials in the form of a string to string map

# Response
Upon receiving a response Unirest returns the result in the form of an Object.

- `code` - HTTP Response Status Code (Example `200`)
- `headers` - HTTP Response Headers
- `raw_body` - HTTP Response as a string

# Advanced Configuration

You can set some advanced configuration to tune Unirest-Go:

### Timeout

You can set a custom timeout value (in **seconds**):

```go
client.SetTimeout(5) # 5s timeout
```

### Default Request Headers

You can set default headers that will be sent on every request:

```go
client.SetDefaultHeader('Header1','Value1')
client.SetDefaultHeader('Header2','Value2')
```

You can clear the default headers anytime with:

```go
client.ClearSetDefaultHeader()
```

### User-Agent

The default User-Agent string is `unirest-ruby/1.1`. You can customize
it like this:

```go
client.SetUserAgent("custom_user_agent")
```

----

Made with &#9829; by [Palash Nigam](https://github.com/palash25)

### TODO

- [ ] Add more comments for the docs
- [ ] Write unit tests
- [ ] Add examples

[unirest-logo]: http://cl.ly/image/2P373Y090s2O/Image%202015-10-12%20at%209.48.06%20PM.png


[license-url]: https://github.com/Mashape/unirest-ruby/blob/master/LICENSE
[license-image]: https://img.shields.io/badge/license-MIT-blue.svg?style=flat
