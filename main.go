package main

import (
	"github.com/abiiram10/helloheroku/datos"
	"github.com/abiiram10/helloheroku/memoria"
	"github.com/abiiram10/helloheroku/web"
)

func main() {
	datos.Ini()
	memoria.Ini()
	web.Ini()
}
