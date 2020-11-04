# Go Gin Example [![rcard](https://goreportcard.com/badge/github.com/tianxinbaiyun/mws)](https://goreportcard.com/report/github.com/tianxinbaiyun/mws) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/tianxinbaiyun/mws) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/EDDYCJY/go-gin-example/master/LICENSE)

An example of gin contains many useful features

[简体中文](https://github.com/tianxinbaiyun/mws/blob/master/README_ZH.md)

## Installation
```
$ go get github.com/tianxinbaiyun/mws
```

## How to run

### Required

- Mysql
- Redis

### Ready

Create a **blog database** and import [SQL](https://github.com/tianxinbaiyun/mws/blob/master/docs/sql/blog.sql)

### Conf

You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```

### Run
```
$ cd $GOPATH/src/go-gin-example

$ go run main.go 
```

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)


Listening port is 8000
Actual pid is 4393
```
Swagger doc


## Features

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis