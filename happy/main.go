package main

import (
	"happy-save-api/common"
	"happy-save-api/server/api"
)

func main() {
	common.RedisInit()
	api.Start()
}
