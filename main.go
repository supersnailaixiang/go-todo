package main

import (
	_ "ToDo/routers"
	_ "ToDo/init"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

