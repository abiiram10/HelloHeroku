package memoria

import (
	"log"

	"abiiram10/HelloHeroku/memoria/usuarios"
)

//Ini inicia lo de memoria
func Ini() {
	log.Println("Servicios memoria")
	usuarios.Ini()
}
