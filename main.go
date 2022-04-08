package main

import (
	"github.com/abiiram10/HelloHeroku/datos"
	"github.com/abiiram10/HelloHeroku/memoria"
	"github.com/abiiram10/HelloHeroku/web"
)

func main() {
	datos.Ini()
	memoria.Ini()
	web.Ini()
}
