# 应用指南

## 项目介绍

- 官网接口地址：https://pexpayapi.github.io/apidocs/spot/cn/#185368440e
- qq群：154056587

## 应用目录

```go
├── README.md
├── go.mod
├── go.sum
├── pkg
│   ├── common.go
│   ├── interface.go
│   ├── service.go
│   └── struct.go
└── utils

```

## 使用帮助

```go
package main

import (
	"log"
	"github.com/ZzzWClock/pexpay-go-sdk/pkg"
)

func main() {

	// 初始化
	pexPay := pkg.PexPayInterface(&pkg.PexPayStruct{
		ApiKey:    "You Apikey",
		ApiSecret: "You ApiSecret",
	})

	// 获取分页订单列表
	data, err := pkg.QueryC2COrderListPage(&pkg.C2COrderListPageForm{
		Page: "1",
		Rows: "10",
	}, pexPay)

	if err != nil {
		log.Println("pkg.QueryC2COrderListPage : ", err)
		return
	}
	
	log.Println("data : ", string(data))
}
```

## 赞赏

![](./static/THEW6Ei4Hdf4bSSRBZNvmkpMMsinyXH6ka.jpeg)
