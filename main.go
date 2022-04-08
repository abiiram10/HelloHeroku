package main

import (
	"abiiram10/helloheroku/datos"
	"abiiram10/helloheroku/memoria"
	"abiiram10/helloheroku/web"
)

func main() {
	datos.Ini()
	memoria.Ini()
	web.Ini()
}
