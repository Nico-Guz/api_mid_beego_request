package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
)

// Envia una peticion al endpoint indicado y extrae la respuesta del campo Data para retomarla
func GetRequestNew(endpoint string, route string, target interface{}) error {
	url := beego.AppConfig.String("ProtocolAdmin") + "://" + beego.AppConfig.String(endpoint) + route
	var reponse map[string]interface{}
	var err error
	err = GetJson(url, &reponse)
	err = ExtractData(reponse, &target)
	return err
}

func GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

// Esta funcion extrae la informacion cuando se recibe encapsulada en una estructura
// y da manejo a las respuestas que contienen arreglos de objetos vacios
func ExtractData(respuesta map[string]interface{}, v interface{}) error {
	var err error
	if respuesta["Success"] == false {
		err = errors.New(fmt.Sprint(respuesta["Data"], respuesta["Message"]))
		panic(err)
	}
	datatype := fmt.Sprintf("%v", respuesta["Data"])
	switch datatype {
	case "map[]", "[map[]]": // response vacio
		break
	default:
		err = FillStruct(respuesta["Data"], &v)
		respuesta = nil
	}
	return err
}

func FillStruct(in interface{}, out interface{}) (err error) {
	var str []byte
	if str, err = json.Marshal(in); err != nil {
		return
	}
	err = json.Unmarshal(str, &out)
	return
}
