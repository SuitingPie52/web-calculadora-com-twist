package main

import ("encoding/json")

type HTTPMessage struct {

	Code int `json: "code"`
	Error string `json: "error"`
	
}

	/*
a. Sucesso (200): {"result": number, "op": string}
b. Operação inválida (400): {"result": "invalid expression", "op": string}
	*/

var TableMessages = map[int]HTTPMessage {


	404: {404, "not found"},
	500: {500, "something went wrong"},


}

func MessageToJson(m HTTPMessage)[]byte {

	j, _ := json.Marshal(m)
	return j

}

