package main

import (
	burn "github.com/BurnFramework/burnweb"
)

func main() {
	br := burn.New()

	// handlers
	br.Get("/helloworld", indexpage)

	// start server
	br.Start(":8080")
}

func indexpage(ctx burn.Context) {
	ctx.SendString("Hello World")
}
