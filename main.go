package main

import (
	burn "github.com/BurnFramework/burnweb"
)

func main() {
	br := burn.New()

	// handlers
	br.Get("/helloworld", indexpage)

	br.Static("/main", "./index.html")

	// start server
	br.Start(":8080")
}

func indexpage(ctx burn.Context) {
	ctx.SendString("Hello World")
}
