package calculo_imposto

func CalcularImposto(salario float64) float64 {
	var imposto float64

	if salario < 50000 {
		imposto = 0
		return imposto
	}
	if salario <= 150000 {
		imposto = salario * 0.17
		return imposto
	}
	
	imposto = salario * 0.27
	return imposto
}

