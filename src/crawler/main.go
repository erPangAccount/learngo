package main

import (
	"crawler/fetche"
	"crawler/parser/zhenai"
)

func main() {
	//zhenai.Run(zhenai2.Seed())
	bytes, e := fetche.Fetcher("https://album.zhenai.com/u/1406875062")
	if e != nil {
		panic(e)
	}
	zhenai.UserInfoParser(bytes)
}
