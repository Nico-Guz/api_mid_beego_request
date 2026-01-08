package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/Nico-Guz/api_mid_beego_request/controllers:ResolucionesParametrosController"] = append(beego.GlobalControllerRouter["github.com/Nico-Guz/api_mid_beego_request/controllers:ResolucionesParametrosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
