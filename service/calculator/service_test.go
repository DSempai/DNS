package calculator_test

import (
	"DNS/service/calculator"
	"testing"
)

func TestService_Multiply(t *testing.T) {
	tests := []struct {
		name         string
		multiplier   float64
		multiplicand float64
		want         float64
	}{
		{
			name:         "float*float",
			multiplier:   12.1,
			multiplicand: 1,
			want:         12.1,
		},
		{
			name:         "null*null",
			multiplier:   0,
			multiplicand: 0,
			want:         0,
		},
		{
			name:         "float*null",
			multiplier:   12.1,
			multiplicand: 0,
			want:         0,
		},
		{
			name:         "int*float",
			multiplier:   12.0000001,
			multiplicand: 2,
			want:         24.0000002,
		},
		{
			name:         "int*int",
			multiplier:   12,
			multiplicand: 12,
			want:         144,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := calculator.Service{}
			if got := s.MultiplicationFloat64(tt.multiplier, tt.multiplicand); got != tt.want {
				t.Errorf("MultiplyCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AdditionFloat64(t *testing.T) {
	tests := []struct {
		name    string
		addends []float64
		want    float64
	}{
		{
			name:    "few values",
			addends: []float64{1.11, 2.22, 3.33},
			want:    6.66,
		},
		{
			name:    "no values",
			addends: []float64{},
			want:    0,
		},
		{
			name:    "values are integers",
			addends: []float64{1, 2, 3, 4},
			want:    10,
		},
		{
			name:    "single number",
			addends: []float64{1.111111},
			want:    1.111111,
		},
	}
	s := calculator.Service{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.AdditionFloat64(tt.addends...); got != tt.want {
				t.Errorf("AdditionFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
