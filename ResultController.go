package main

import(
	"net/http"
	"encoding/json"
	"strings"
	"strconv"
)

var ResultController = map[string]map[string]func(http.ResponseWriter, *http.Request) {

	"GET": GetFunctions,

}

var GetFunctions = map[string]func(http.ResponseWriter, *http.Request) {
	
	"/result": GetResult,
	
}

func GetResult (res http.ResponseWriter, req *http.Request) {

	op := req.URL.Query().Get("op")

	noDashOp := strings.ReplaceAll(op, " ", "")

	var CheckOperators = map[bool]string{}

	for k := range TableCalculable {
	
		CheckOperators[strings.Contains(noDashOp, k)] = k
	
	}
	
	if !isAnyOperationValid(CheckOperators) {
	
		WriteInvalidExpression(res, op)
		return
		
	}
	
	slicedOp := strings.Split(noDashOp, CheckOperators[true])
	
	if len(slicedOp) > 2 {
	
		WriteInvalidExpression(res, op)
		return
	
	}
		
	var1, err1 := strconv.ParseFloat(slicedOp[0], 64)
	var2, err2 := strconv.ParseFloat(slicedOp[1], 64)
	
	if err1 != nil || err2 != nil {
	
		WriteInvalidExpression(res, op)
		return		
	
	}
	
	s, ratErr := TableCalculable[CheckOperators[true]].Calculate(var1, var2, op)
	
	if ratErr != nil {
	
		res.Write(MessageToJson(TableMessages[500]))
		return
	
	}
	
	res.Write(CalculableToJson(s))
	
	return

}

func isAnyOperationValid (m map[bool]string)bool {

	return m[true] != ""

}

func WriteInvalidExpression(res http.ResponseWriter, op string) {

	e := map[string]string{"result": "invalid expression", "op": op}
	j, _ := json.Marshal(e)
	res.Write(j)
	return

}

