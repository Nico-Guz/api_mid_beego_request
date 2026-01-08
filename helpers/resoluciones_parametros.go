package helpers

import (
	"fmt"

	"github.com/Nico-Guz/api_mid_beego_request/models"
	"github.com/astaxie/beego/logs"
)

func ListarResolucionesParametros() (resolucionParametros models.ResolucionesParametros, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"function": "ListarResolucionesParametros", "error": err, "status": "500"}
			panic(outputError)
		}
	}()

	var parametros []models.Parametro
	// var resoluciones []models.Resolucion
	var resoluciones []models.DatosLocales

	url := "parametro?query=TipoParametroId__CodigoAbreviacion:TR&limit=0"
	if err := GetRequestNew("UrlCrudParametros", url, &parametros); err != nil {
		logs.Error(err)
		panic(err.Error())
	}
	fmt.Println("PARAMETROS", parametros)

	// urlResoluciones := "resolucion?limit=0"
	urlResoluciones := "usuario?limit=0"
	if err := GetRequestNew("UrlCrudResoluciones", urlResoluciones, &resoluciones); err != nil {
		logs.Error(err)
		panic(err.Error())
	}
	fmt.Println("RESOLUCIONES", resoluciones)

	resolucionParametros.Resoluciones = resoluciones
	resolucionParametros.Parametros = parametros

	fmt.Println(resolucionParametros)
	return resolucionParametros, outputError
}
