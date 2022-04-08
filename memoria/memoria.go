package memoria

import (
	"abiiram10/helloheroku/memoria/usuarios"
	"log"
)

//Ini inicia lo de memoria
func Ini() {
	log.Println("Servicios memoria")
	usuarios.Ini()
}
