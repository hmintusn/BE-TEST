package fee

import "letsgo/internal/domain"

type WeightCalculator struct {
    Coefficients domain.Coefficients
}

func (w WeightCalculator) Calculate(item domain.Item) float64 {
    if item.WeightKg <= 0 || w.Coefficients.WeightPerKg <= 0 {
        return 0
    }
    return item.WeightKg * w.Coefficients.WeightPerKg
}


