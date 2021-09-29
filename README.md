# simple-url-shortener

[![Go](https://github.com/Zchary-Ma/project52-url-shortner/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/Zchary-Ma/project52-url-shortener/actions/workflows/go.yml)

A url shortener writing in Go.

## How To Start

1. initiate redis config `config.yml` file in the root directory.(example:`config.yml.exmaple`)
2. run `go get -u`
3. run `go run main.go` server will start at default port. `http://localhost:8080` 

## endpoint

- `POST`:`/shorten` 
- `GET`:`/redirect`

## Bonus

> 如何抗住大qps，抗大流量 。url 哈希函数怎么设计（怎么存，怎么统计qps）
