package main

import (
	"myproject/global"
)

const env = "dev"

func main() {
	_ = global.GetInstance(env)
}
