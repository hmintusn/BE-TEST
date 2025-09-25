# Shipping Service Calculator

A flexible and extensible Go application that calculates shipping costs for Amazon products purchased through a Vietnamese shipping service. Built following Clean Architecture principles and SOLID design patterns.

## Architecture Overview

This project implements Clean Architecture with clear separation of concerns:

- **Domain Layer** (`internal/domain`): Core business entities and interfaces
- **Service Layer** (`internal/service`): Business logic and orchestration  
- **Infrastructure Layer** (`internal/config`): External concerns and implementations

## Key Design Principles

### SOLID Principles Applied

- **Single Responsibility**: Each calculator handles one fee type (weight/dimension)
- **Open/Closed**: New fee calculators can be added without modifying existing code
- **Liskov Substitution**: All calculators implement the same `FeeCalculator` interface
- **Interface Segregation**: Small, focused interfaces like `FeeCalculator`
- **Dependency Inversion**: Services depend on abstractions (`FeeCalculator` interface), not concrete implementations

### Flexible Shipping Fee Architecture

The shipping fee calculation is designed for extensibility through the `FeeCalculator` interface:

```go
type FeeCalculator interface {
    Calculate(item Item) float64
}
```

This design allows adding new fee types (product-specific, region-based, etc.) without changing the core `ShippingService` logic. The service simply iterates through all registered calculators and takes the maximum fee.

## Project Structure

```
letsgo/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── domain/
│   │   └── models.go        # Core entities and interfaces
│   ├── service/
│   │   ├── order.go         # Order processing logic
│   │   └── shipping.go      # Shipping fee orchestration
│   ├── fee/
│   │   ├── weight.go        # Weight-based fee calculator
│   │   └── dimension.go     # Dimension-based fee calculator
│   └── config/
│       └── config.go        # Configuration loading
├── config.json              # Shipping coefficients
└── README.md
```

## How It Works

1. **Load Configuration**: Coefficients loaded from JSON config file
2. **Initialize Calculators**: Weight and dimension calculators created with coefficients
3. **Register Calculators**: All calculators registered with `ShippingService`
4. **Calculate Fees**: For each item, all calculators run and maximum fee is selected
5. **Compute Gross Price**: Amazon price + shipping fee for each item

## Extensibility Example

To add product-type-based fees:

```go
// 1. Create new calculator
type ProductTypeCalculator struct {
    ProductRates map[string]float64
}

func (p ProductTypeCalculator) Calculate(item domain.Item) float64 {
    return p.ProductRates[item.ProductType]
}

// 2. Register with service (no code changes needed in ShippingService)
productCalc := ProductTypeCalculator{ProductRates: productRates}
ship := service.ShippingService{
    Calculators: []domain.FeeCalculator{weightCalc, dimCalc, productCalc},
}
```

## Running the Application

```bash

go run cmd/shipping/main.go

```

## Benefits of This Architecture

- **Testability**: Each component can be unit tested in isolation
- **Maintainability**: Clear separation of concerns makes code easy to understand
- **Extensibility**: New fee types can be added without modifying existing code
- **Configuration**: Coefficients are externalized and easily adjustable
- **Dependency Inversion**: High-level modules don't depend on low-level details

The use of interfaces and dependency injection makes this system highly flexible and future-proof for business requirement changes.