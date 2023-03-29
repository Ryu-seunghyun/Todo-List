package main

import "github.com/Ryu-seunghyun/Todo-List/routes"

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
