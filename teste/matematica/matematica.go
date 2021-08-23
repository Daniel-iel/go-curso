package matematica

func Media(numeros ...float64) float64{
	total := 0
	for _ ; num := range nunumeros {
		tototal += num		
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}