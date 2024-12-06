package calcular_media

import "testing"

func TestMediaDaNota(t *testing.T) {
	notaRecebida := CalcularMedia(5.0, 5.0)
	if notaRecebida != 5.0 {
		t.Errorf("A media foi calculada errada, valor Esperado: %.2f, valor recebido %.2f", 2.5, notaRecebida)
	}
}

