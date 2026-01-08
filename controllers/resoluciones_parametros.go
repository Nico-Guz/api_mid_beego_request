package controllers

import (
	"github.com/Nico-Guz/api_mid_beego_request/helpers"
	"github.com/astaxie/beego"
)

type ResolucionesParametrosController struct {
	beego.Controller
}

func (c *ResolucionesParametrosController) URLMapping() {}

// GetAll ...
// @Title Get All
// @Description get Obtener resoluciones y parametros
// @Success 200 {object} models.ResolucionesParametros
// @Failure 400 bad request
// @Failure 500 Internal Server Error
// @router / [get]
func (c *ResolucionesParametrosController) GetAll() {
	defer helpers.ErrorController(c.Controller, "ResolucionesParametrosController")

	if v, err := helpers.ListarResolucionesParametros(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Listado consultado con exito", "Data": v}
	} else {
		panic(err)
	}

	c.ServeJSON()
}
