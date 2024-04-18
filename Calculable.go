package main

import ("encoding/json")

type Calculable interface {

	Calculate(float64, float64, string)Calculable

}

var TableCalculable = map[string]Calculable{

	"+":   Sum{},
	"sum": Sum{},
	"-":   Sub{},
	"sub": Sub{},
	"*":   Mul{},
	"mul": Mul{},
	"/":   Div{},
	"div": Div{},
	"^":   Pow{},
	"pow": Pow{},
	"&":   Rot{},
	"rot": Rot{},
	
}

func CalculableToJson(c Calculable)[]byte {

	j, _ := json.Marshal(c)
	return j

}

