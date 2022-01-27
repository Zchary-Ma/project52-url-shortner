# url-shortening-server

[![Go](https://github.com/Zchary-Ma/url-shortening-server/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/Zchary-Ma/url-shortening-server/actions/workflows/go.yml)

A url shortening server written in Go.
## endpoint

- `GET`:`/api/{shortened}` redirect shortened url
- `POST`:`/api/s` shorten url 
- `GET`:`/api/h` health check
- `GET`:`/` html index page 

## Bonus

> 如何抗住大qps，抗大流量 。url 哈希函数怎么设计（怎么存，怎么统计qps）
