package main

import (
	"github.com/gin-gonic/gin"
	"myproject/main/global"
)

const env = "dev"

func main() {
	config := global.GetInstance(env)

	engine := gin.Default()

	engine.Run(config.Port())

}
