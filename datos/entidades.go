package datos

import "time"

//Usuario : struct of the Usuario table
type Usuario struct {
	Id_usuario      int       `json:"Id_usuario"`
	Vch_nombre      string    `json:"Nombre"`
	Vch_paterno     string    `json:"Paterno"`
	Vch_materno     string    `json:"Materno"`
	Vch_correo      string    `json:"Correo"`
	Vch_telefono    string    `json:"Telefono"`
	Vch_username    string    `json:"Username"`
	Vch_password    string    `json:"Password"`
	Bol_enuso       bool      `json:"Bol_enuso"`
	Dt_modificacion time.Time `json:"Dt_modificacion"`
	Int_id_modifico int       `json:"-"`
}
