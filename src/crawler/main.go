package main

import (
	"crawler/engine/zhenai"
	zhenai2 "crawler/seed/zhenai"
)

func main() {
	zhenai.Run(zhenai2.Seed())
}
