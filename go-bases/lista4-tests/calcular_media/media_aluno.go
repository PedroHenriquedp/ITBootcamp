package calcular_media

func CalcularMedia(notas ...float64) float64{
	var soma float64
	len := len(notas)
	for _, nota := range notas {
		soma += nota
	}
	mediaFinal := soma/float64(len)

	return mediaFinal
}
