package dto

type FindParameterByCodeDto struct {
	Nombre                  string `json:"nombre"`
	Codigo                  string `json:"codigo"`
	ValorCodigo             string `json:"valor_codigo"`
	ValorDescripcion        string `json:"valor_descripcion"`
	ParametroDetallePadreId string `json:"parametro_detalle_padre_id"`
}
