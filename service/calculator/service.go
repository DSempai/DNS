package calculator

type Service struct{}

func Initialize() Service {
	return Service{}
}

func (s Service) MultiplicationFloat64(multiplier float64, multiplicand float64) float64 {
	return multiplier * multiplicand
}

func (s Service) AdditionFloat64(addends ...float64) float64 {
	var sum float64
	for _, addend := range addends {
		sum += addend
	}
	return sum
}
