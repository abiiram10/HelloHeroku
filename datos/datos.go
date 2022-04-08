package datos

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var ConexionP *sql.DB

// Ini inicio del modulo
func Ini() {
	var ErrCon error
	ConexionP, ErrCon = obtenerConexion("ec2-52-3-60-53.compute-1.amazonaws.com", 5432, "mbpkzcukfieopj", "a61f6ed2bd0a18704352b5f564f0abbf0adee26fed94ba116095c8c4ca577d27", "dc1h538edouk8")
	if ErrCon != nil {
		log.Println("Ocurrio un error")
	} else {
		log.Println("Conexion establecida")
	}
}

func obtenerConexion(_DireccionIP string, _Puerto int, _User string, _Pass string, _BDname string) (Conexion *sql.DB, err_open error) {
	CadenaConexion := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", _DireccionIP, _Puerto, _User, _Pass, _BDname)
	Conexion, err_open = sql.Open("postgres", CadenaConexion)
	if err_open != nil {
		log.Println(err_open.Error())
	}

	ErrPing := Conexion.Ping()
	if ErrPing != nil {
		log.Println(ErrPing.Error())
	}

	return Conexion, err_open
}

//Closeconexion cierra la conexion abierta al inicio
func Closeconexion() {
	log.Println("Se cerró la conexión")
	ConexionP.Close()
}
