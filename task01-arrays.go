package homework

func average(input [15]float32) (result float32) {
	for _, i := range input {
		result += i
	}
	result = result / 15

	return
}
