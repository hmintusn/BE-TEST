package fee

import (
    "testing"
    "letsgo/internal/domain"
)

func TestWeightCalculator(t *testing.T) {
    calc := WeightCalculator{Coefficients: domain.Coefficients{WeightPerKg: 11}}
    item := domain.Item{WeightKg: 2}
    got := calc.Calculate(item)
    want := 22.0
    if got != want {
        t.Fatalf("expected %v, got %v", want, got)
    }
}


