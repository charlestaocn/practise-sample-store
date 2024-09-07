package main

import "practise-sample-store/router"

func main() {
	engine := router.Router()
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}
