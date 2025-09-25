package main

import (
    "fmt"
    "os"
    "path/filepath"
    "letsgo/internal/config"
    "letsgo/internal/domain"
    "letsgo/internal/fee"
    "letsgo/internal/service"
)

func main() {
    // Determine config path 
    cwd, _ := os.Getwd()
    defaultPath := filepath.Join(cwd, "config.json")
    path := defaultPath
    if len(os.Args) > 1 && os.Args[1] != "" {
        path = os.Args[1]
    }

    cfg, err := config.Load(path)
    if err != nil {
        fmt.Println("failed to load config:", err)
        os.Exit(1)
    }

    weightCalc := fee.WeightCalculator{Coefficients: cfg.Coefficients}
    dimCalc := fee.DimensionCalculator{Coefficients: cfg.Coefficients}
    ship := service.ShippingService{Calculators: []domain.FeeCalculator{weightCalc, dimCalc}}
    orderSvc := service.OrderService{Shipping: ship}

    // Sample items
    items := []domain.Item{
        {AmazonPrice: 100, WeightKg: 0.3, WidthM: 0.07, HeightM: 0.015, DepthM: 0.15, ProductType: "smartphone"},
        {AmazonPrice: 500, WeightKg: 0.01, WidthM: 0.03, HeightM: 0.03, DepthM: 0.03, ProductType: "ring"},
    }
    order := domain.Order{Items: items}

    gross := orderSvc.GrossPrice(order)
    fmt.Printf("Gross price: $%.2f\n", gross)
}


