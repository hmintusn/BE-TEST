package fee

import "letsgo/internal/domain"

type DimensionCalculator struct {
    Coefficients domain.Coefficients
}

func (d DimensionCalculator) Calculate(item domain.Item) float64 {
    if item.WidthM <= 0 || item.HeightM <= 0 || item.DepthM <= 0 || d.Coefficients.DimensionPerM3 <= 0 {
        return 0
    }
    volume := item.WidthM * item.HeightM * item.DepthM
    return volume * d.Coefficients.DimensionPerM3
}


