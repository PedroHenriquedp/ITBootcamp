package calculo_imposto

import "testing"

func TestCalcularImpostoAbaixo50000(t *testing.T) {
	salario := 9000.0
	impostoEsperado := 0.0

	impostoCalculado := CalcularImposto(salario)

	if impostoCalculado != impostoEsperado {
		t.Errorf("Imposto incorreto para salario abaixo de 50000. Era esperado %.2f, valor obtido: %.2f", impostoEsperado, impostoCalculado)
	}
}

func TestCalcularImpostoEntre50000E150000(t *testing.T){
	salario := 100000.0
	impostoEsperado := salario * 0.17
	
	impostoCalculado := CalcularImposto(salario)

	if impostoCalculado != impostoEsperado {
		t.Errorf("Imposto incorreto para salario entre 50000 e 150000. Era esperado %.2f, valor obtido: %.2f", impostoEsperado, impostoCalculado)
	}
}

func TestCalcularImpostoAcima150000(t *testing.T) {
	salario := 200000.0
	impostoEsperado := salario * 0.27

	impostoCalculado := CalcularImposto(salario)

	if impostoCalculado != impostoEsperado {
		t.Errorf("Imposto incorreto para salario acima de 150000. Esperado: %.2f e valor obtido: %.2f", impostoEsperado, impostoCalculado)
	}
}

