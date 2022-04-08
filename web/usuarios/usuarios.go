package usuarios

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abiiram10/HelloHeroku/datos"
	"github.com/abiiram10/HelloHeroku/memoria/usuarios"
)

// GetUsuariosPOST Obtiene los usuarios de la tabla
func GetUsuariosPOST(Response http.ResponseWriter, Request *http.Request) {
	defer Request.Body.Close()

	if Request.Method != http.MethodPost { //Si es diferente de POST
		Response.WriteHeader(http.StatusBadRequest)
		return
	}

	parametros := struct {
		Modificacion *time.Time
	}{}

	err_decoder := json.NewDecoder(Request.Body).Decode(&parametros)
	if err_decoder != nil {
		Response.WriteHeader(http.StatusBadRequest)
		return
	}

	arrayUsers := usuarios.GetUserPOST(parametros.Modificacion)

	RespuestaBytes, err_marshal := json.Marshal(map[string]interface{}{
		"Datos":     arrayUsers,
		"Resultado": 1,
	})

	if err_marshal != nil {
		Response.WriteHeader(http.StatusBadRequest)
		return
	}

	Response.Header().Set("Access-Control-Allow-Origin", "*")
	Response.Header().Set("Content-Type", "application/json")
	Response.WriteHeader(http.StatusOK)
	Response.Write(RespuestaBytes)
}

func SetUsuarioPOST(Response http.ResponseWriter, Request *http.Request) {
	defer Request.Body.Close()

	if Request.Method != http.MethodPost { //Si es diferente de POST
		Response.WriteHeader(http.StatusBadRequest)
		return
	}

	parametros := struct {
		Nombre            string
		Paterno           string
		Materno           string
		Correo            string
		Telefono          string
		Username          string
		Password          string
		IdUsuarioModifico int
	}{}

	err_decoder := json.NewDecoder(Request.Body).Decode(&parametros)
	if err_decoder != nil {
		Response.WriteHeader(http.StatusBadRequest)
		return
	}

	var nUsuario datos.Usuario
	nUsuario.Vch_nombre = parametros.Nombre
	nUsuario.Vch_paterno = parametros.Paterno
	nUsuario.Vch_materno = parametros.Materno
	nUsuario.Vch_correo = parametros.Correo
	nUsuario.Vch_telefono = parametros.Telefono
	nUsuario.Vch_username = parametros.Username
	nUsuario.Vch_password = parametros.Password
	nUsuario.Bol_enuso = true
	nUsuario.Dt_modificacion = time.Now().UTC()
	nUsuario.Int_id_modifico = parametros.IdUsuarioModifico

	idResponse := usuarios.SetUserPOST(nUsuario)

	RespuestaBytes, err_marshal := json.Marshal(map[string]interface{}{
		"Resultado": idResponse,
	})

	if err_marshal != nil {
		Response.WriteHeader(http.StatusBadRequest)
		return
	}
	Response.Header().Set("Access-Control-Allow-Origin", "*")
	Response.Header().Set("Content-Type", "application/json")
	Response.WriteHeader(http.StatusOK)
	Response.Write(RespuestaBytes)

}
