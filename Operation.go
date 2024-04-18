package main

import ("math"	
	"errors"
	"strconv"
	"strings"
)

type Operation struct {

	Result float64 `json: "result"`
	Op string `json: "op"`

}

type Sum struct {

	Operation
	
}

func (s Sum) Calculate(var1 float64, var2 float64, op string)(Calculable, error) {

	s.Result = var1 + var2
	s.Op = op

	if isRationalNumber(s.Result) {
	
		return s, errors.New("Número racional")
	
	}
	
	return s, nil

}

type Sub struct {

	Operation
	
}

func (s Sub) Calculate(var1 float64, var2 float64, op string)(Calculable, error) {

	s.Result = var1 - var2
	s.Op = op
	
	if isRationalNumber(s.Result) {
	
		return s, errors.New("Número racional")
	
	}
	
	return s, nil

}

type Mul struct {

	Operation
	
}

func (s Mul) Calculate(var1 float64, var2 float64, op string)(Calculable, error) {

	s.Result = var1 * var2
	s.Op = op
	
	if isRationalNumber(s.Result) {
	
		return s, errors.New("Número racional")
	
	}
	
	return s, nil

}

type Div struct {

	Operation
	
}

func (s Div) Calculate(var1 float64, var2 float64, op string)(Calculable, error) {

	s.Result = var1 / var2
	s.Op = op
	
	if isRationalNumber(s.Result) {
	
		return s, errors.New("Número racional")
	
	}
	
	return s, nil

}

type Pow struct {

	Operation
	
}

func (s Pow) Calculate(var1 float64, var2 float64, op string)(Calculable, error) {

	s.Result = math.Pow(var1, var2)
	s.Op = op
	
	if isRationalNumber(s.Result) {
	
		return s, errors.New("Número racional")
	
	}
	
	return s, nil

}

type Rot struct {

	Operation
	
}

func (s Rot) Calculate(var1 float64, var2 float64, op string)(Calculable, error) {

	s.Result = math.Pow(var1, 1.0/var2)
	s.Op = op
	
	if isRationalNumber(s.Result) {
	
		return s, errors.New("Número racional")
	
	}
	
	return s, nil

}

func isRationalNumber (f float64) bool {

	s := strconv.FormatFloat(f, 'g', -1, 64)
	
	if !strings.Contains(s, ".") {
	
		return true
	
	}
	
	slicedS := strings.Split(s, ".")
	
	return len(slicedS[1]) < 13
	
	

}


