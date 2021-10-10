package main

import (
	"github.com/seckill/sk-core/setup"
)

func main() {

	setup.InitZk()
	setup.InitRedis()
	setup.RunService()

}
