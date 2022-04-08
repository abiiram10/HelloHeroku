package usuarios

import (
	"abiiram10/helloheroku/datos"
	"abiiram10/helloheroku/datos/usuarios"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	collUsuarios sync.Map
	FechaMaxima  *time.Time
)

//Inicia dd
func Ini() {
	collUsuarios = sync.Map{}

	dtMax, err := usuarios.GetUsers(&collUsuarios, FechaMaxima)
	if err != nil {
		log.Println("Ocurrio un error: TBL_USUARIOS")
	}

	if dtMax != nil {
		FechaMaxima = dtMax
		fmt.Println("Nueva fecha -" + FechaMaxima.UTC().Format("2006-01-02 15:04:05.999"))
	}

	go actualizar()
}

func actualizar() {
	Timer := time.NewTimer(time.Second * 5)
	for range Timer.C {
		Timer.Stop()

		dtMax, err := usuarios.GetUsers(&collUsuarios, FechaMaxima)
		if err != nil {
			log.Println("Ocurrio un error: TBL_USUARIOS")
		}

		if dtMax != nil {
			FechaMaxima = dtMax
			fmt.Println("Nueva fecha -" + FechaMaxima.UTC().Format("2006-01-02 15:04:05.999"))
		}
		Timer.Reset(time.Second * 5)
	}
}

//GetUserPOST Obtiene los usuarios en memoria
func GetUserPOST(Modificacion *time.Time) []datos.Usuario {
	arrayUser := []datos.Usuario{}

	collUsuarios.Range(func(_, value any) bool {
		xUsuario, ok := value.(datos.Usuario)
		if ok {
			if Modificacion != nil {
				if xUsuario.Dt_modificacion.After(*Modificacion) {
					arrayUser = append(arrayUser, xUsuario)
				}
			} else {
				arrayUser = append(arrayUser, xUsuario)
			}
		}
		return true // t = seguir, f = detener
	})

	return arrayUser
}

func SetUserPOST(nUsuario datos.Usuario) int {
	var idNuevoUsuario int = usuarios.SetUser(nUsuario)
	if idNuevoUsuario > 0 {
		nUsuario.Id_usuario = idNuevoUsuario
		collUsuarios.Store(nUsuario.Id_usuario, nUsuario)
	}
	return idNuevoUsuario
}
