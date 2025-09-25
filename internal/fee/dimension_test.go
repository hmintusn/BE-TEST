package fee

import (
    "testing"
    "letsgo/internal/domain"
)

func TestDimensionCalculator(t *testing.T) {
    calc := DimensionCalculator{Coefficients: domain.Coefficients{DimensionPerM3: 11}}
    item := domain.Item{WidthM: 1, HeightM: 2, DepthM: 3}
    got := calc.Calculate(item)
    want := 66.0 // 1*2*3*11
    if got != want {
        t.Fatalf("expected %v, got %v", want, got)
    }
}


