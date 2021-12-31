package main

import (
	burn "github.com/BurnFramework/burnweb"
)

type todo struct {
	Done string
	Time int64
}

var todolist = []todo{
	{Done: "Gym", Time: 2.00},
}

func main() {
	br := burn.New()

	// handlers
	br.Get("/helloworld", indexpage)
	br.Get("/todo", todopage)

	br.Static("/main", "./index.html")

	// start server
	br.Start(":8080")
}

func indexpage(ctx burn.Context) {
	ctx.SendString("Hello World")
}

func todopage(ctx burn.Context) {
	ctx.SendJSON(todolist)
}
