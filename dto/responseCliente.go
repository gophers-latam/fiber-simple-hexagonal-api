package dto

type ResponseCliente struct {
	ID              string `json:"cliente_id"`
	Nombre          string `json:"nombre"`
	Ciudad          string `json:"ciudad"`
	CodigoPostal    string `json:"codigo_postal"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	Estatus         string `json:"estatus"`
}
