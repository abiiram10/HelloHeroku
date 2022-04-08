package main

import (
	"abiiram10/HelloHeroku/datos"
	"abiiram10/HelloHeroku/memoria"
	"abiiram10/HelloHeroku/web"
)

func main() {
	datos.Ini()
	memoria.Ini()
	web.Ini()
}
