package main

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

