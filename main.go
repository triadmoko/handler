package main

import "github.com/triadmoko/handler/logs"

func main() {
	logs.NewLogger("./logs/logger.log", true)

}
