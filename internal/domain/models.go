package domain

type Item struct {
    AmazonPrice float64
    WeightKg    float64
    WidthM      float64
    HeightM     float64
    DepthM      float64
    ProductType string
}

// Order represents a collection of items.
type Order struct {
    Items []Item
}

// Coefficients holds configuration values for calculators.
type Coefficients struct {
    WeightPerKg     float64 `json:"weight_per_kg"`
    DimensionPerM3  float64 `json:"dimension_per_m3"`
}


type FeeCalculator interface {
    Calculate(item Item) float64
}


