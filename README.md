# Go Simple LB

[![Build Status](https://travis-ci.com/cloudingcity/go-simple-lb.svg?branch=master)](https://travis-ci.com/cloudingcity/go-simple-lb)
[![codecov](https://codecov.io/gh/cloudingcity/go-simple-lb/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudingcity/go-simple-lb)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudingcity/go-simple-lb)](https://goreportcard.com/report/github.com/cloudingcity/go-simple-lb)

A simple load balancer written in Go.

- Use Round Robin algorithm
- Health check to recovery for unhealthy servers in every 1 min 

## Usage

```shell script
$ go-simple-lb -port 8080 -servers=http://localhost:8081,http://localhost:8082,http://localhost:8083
```
## Demo

```shell script
$ docker-compose up --build

# send request to load balancer
$ curl :8080
{"hostname":"25e08d5ffec5"}
$ curl :8080
{"hostname":"2df976f69607"}
$ curl :8080
{"hostname":"5d96ca9bd34d"}
$ curl :8080
{"hostname":"25e08d5ffec5"}
```
