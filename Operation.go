package main

import ("math")

type Operation struct {

	Result float64 `json: "result"`
	Op string `json: "op"`

}

type Sum struct {

	Operation
	
}

func (s Sum) Calculate(var1 float64, var2 float64, op string)Calculable {

	s.Result = var1 + var2
	s.Op = op
	return s

}

type Sub struct {

	Operation
	
}

func (s Sub) Calculate(var1 float64, var2 float64, op string)Calculable {

	s.Result = var1 - var2
	s.Op = op
	return s

}

type Mul struct {

	Operation
	
}

func (s Mul) Calculate(var1 float64, var2 float64, op string)Calculable {

	s.Result = var1 * var2
	s.Op = op
	return s

}

type Div struct {

	Operation
	
}

func (s Div) Calculate(var1 float64, var2 float64, op string)Calculable {

	s.Result = var1 / var2
	s.Op = op
	return s

}

type Pow struct {

	Operation
	
}

func (s Pow) Calculate(var1 float64, var2 float64, op string)Calculable {

	s.Result = math.Pow(var1, var2)
	s.Op = op
	return s

}

type Rot struct {

	Operation
	
}

func (s Rot) Calculate(var1 float64, var2 float64, op string)Calculable {

	s.Result = math.Pow(var1, 1.0/var2)
	s.Op = op
	return s

}

