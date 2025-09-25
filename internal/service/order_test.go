package service

import (
    "testing"
    "letsgo/internal/domain"
    "letsgo/internal/fee"
)

func TestOrderGrossPrice(t *testing.T) {
    coeff := domain.Coefficients{WeightPerKg: 11, DimensionPerM3: 11}
    ship := ShippingService{Calculators: []domain.FeeCalculator{
        fee.WeightCalculator{Coefficients: coeff},
        fee.DimensionCalculator{Coefficients: coeff},
    }}
    svc := OrderService{Shipping: ship}

    order := domain.Order{Items: []domain.Item{
        {AmazonPrice: 100, WeightKg: 0.3, WidthM: 0.07, HeightM: 0.015, DepthM: 0.15},
        {AmazonPrice: 500, WeightKg: 0.01, WidthM: 0.03, HeightM: 0.03, DepthM: 0.03},
    }}

    got := svc.GrossPrice(order)
    want := 603.41 
    if diff := got - want; diff > 1e-9 || diff < -1e-9 {
        t.Fatalf("expected %v, got %v", want, got)
    }
}


