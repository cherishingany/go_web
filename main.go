package main

import (
	_ "github.com/spf13/viper"
	"go_web/engine"
)

func main() {

	router := engine.NewEngine()

	//r.Static("/dwz", "/Users/nationsky/Desktop/goapp/src/go_web/view/html")
	//
	//r.LoadHTMLFiles("/Users/nationsky/Desktop/goapp/src/go_web/view/html/error.htmlindex1.html")

	router.Run() // listen and serve on 0.0.0.0:8080

}
