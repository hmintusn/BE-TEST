package service

import "letsgo/internal/domain"

// ShippingService aggregates fee calculators and returns the maximum fee for an item.
type ShippingService struct {
    Calculators []domain.FeeCalculator
}

func (s ShippingService) ShippingFee(item domain.Item) float64 {
    maxFee := 0.0
    for i := 0; i < len(s.Calculators); i++ {
        calc := s.Calculators[i]
        fee := calc.Calculate(item)
        if fee > maxFee {
            maxFee = fee
        }
    }
    return maxFee
}