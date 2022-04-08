package usuarios

import (
	"abiiram10/helloheroku/datos"
	"fmt"
	"log"
	"sync"
	"time"
)

//GetUsers obtiene los usuarios
func GetUsers(collUsuarios *sync.Map, Modificacion *time.Time) (*time.Time, error) {
	var dtMaximo *time.Time
	sqlQuery := "SELECT int_id_usuario, vch_nombre,vch_paterno,vch_materno,vch_correo,vch_telefono,vch_username,vch_password, bol_enuso, dt_modificacion  FROM catalogo.tbl_usuario "
	if Modificacion == nil {
		sqlQuery += " WHERE bol_enuso"
	} else {
		sqlQuery += " WHERE dt_modificacion>'" + Modificacion.UTC().Format("2006-01-02 15:04:05.999") + "'"
	}

	rows, err_query := datos.ConexionP.Query(sqlQuery)
	if err_query != nil {
		log.Println("Ocurrio un error al intentar obtener los usuarios")
		return dtMaximo, err_query
	}

	defer rows.Close()

	for rows.Next() {
		var xUsuario datos.Usuario

		err_scan := rows.Scan(&xUsuario.Id_usuario, &xUsuario.Vch_nombre, &xUsuario.Vch_paterno, &xUsuario.Vch_materno, &xUsuario.Vch_correo,
			&xUsuario.Vch_telefono, &xUsuario.Vch_username, &xUsuario.Vch_password, &xUsuario.Bol_enuso, &xUsuario.Dt_modificacion)

		if err_scan != nil {
			fmt.Println("Erro al intentar scan GetUsers")
			return dtMaximo, err_scan
		}

		if dtMaximo == nil || dtMaximo.Before(xUsuario.Dt_modificacion) {
			dtMaximo = &xUsuario.Dt_modificacion
		}

		collUsuarios.Store(xUsuario.Id_usuario, xUsuario)
	}

	return dtMaximo, nil
}

func SetUser(nUsuario datos.Usuario) int {
	var IdNuevoUsuario int

	sqlStatement := `INSERT INTO catalogo.tbl_usuario(vch_nombre, vch_paterno, vch_materno, vch_correo, vch_telefono, vch_username, vch_password, bol_enuso, int_usuario_registro, dt_registro, int_usuario_modifico, dt_modificacion)
		   VALUES ('%s','%s','%s','%s','%s','%s','%s',true,%d,Current_timestamp(3) AT TIME ZONE 'UTC',%d,Current_timestamp(3) AT TIME ZONE 'UTC') RETURNING int_id_usuario;`

	sqlQuery := fmt.Sprintf(sqlStatement, nUsuario.Vch_nombre, nUsuario.Vch_paterno, nUsuario.Vch_materno, nUsuario.Vch_correo, nUsuario.Vch_telefono, nUsuario.Vch_username, nUsuario.Vch_password, nUsuario.Int_id_modifico, nUsuario.Int_id_modifico)

	err_scan := datos.ConexionP.QueryRow(sqlQuery).Scan(&IdNuevoUsuario)
	if err_scan != nil {
		return -1
	}

	return IdNuevoUsuario
}

func UpdateUser(nUsuario datos.Usuario) bool {
	sqlStatement := `UPDATE catalogo.tbl_usuario SET 
		vch_nombre='%s', vch_paterno='%s', vch_materno='%s', vch_correo='%s',vch_telefono='%s', vch_username='%s',
		vch_password='%s', bol_enuso=%t,int_usuario_modifico=%d,dt_modificacion=Current_timestamp(3) AT TIME ZONE 'UTC'
		WHERE int_id_usuario= %d;`

	sqlQuery := fmt.Sprintf(sqlStatement, nUsuario.Vch_nombre, nUsuario.Vch_paterno, nUsuario.Vch_materno, nUsuario.Vch_correo, nUsuario.Vch_telefono, nUsuario.Vch_username,
		nUsuario.Vch_password, nUsuario.Bol_enuso, nUsuario.Int_id_modifico, nUsuario.Id_usuario)

	log.Println(sqlQuery)

	return true
}

func DeleteUser(nUsuario datos.Usuario) bool {
	sqlStatement := `DELETE FROM catalogo.tbl_usuario WHERE int_id_usuario=%d;`
	sqlQuery := fmt.Sprintf(sqlStatement, nUsuario.Id_usuario)
	log.Println(sqlQuery)
	return true
}
