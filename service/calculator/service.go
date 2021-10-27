package calculator

// Service is responsible for performing mathematical operations.
// Service does not need to know about specific types of values when performing operations.
type Service struct{}

func Initialize() Service {
	return Service{}
}

// MultiplicationFloat64 return product of two given values: multiplier and multiplicand.
func (s Service) MultiplicationFloat64(multiplier float64, multiplicand float64) float64 {
	return multiplier * multiplicand
}

// AdditionFloat64 return sum of given arguments.
func (s Service) AdditionFloat64(addends ...float64) float64 {
	var sum float64
	for _, addend := range addends {
		sum += addend
	}
	return sum
}
