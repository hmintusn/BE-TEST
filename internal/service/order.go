package service

import "letsgo/internal/domain"


type OrderService struct {
    Shipping ShippingService
}

// Gross price of all items of order.
func (o OrderService) GrossPrice(order domain.Order) float64 {
    var total float64
    for _, it := range order.Items {
        total += o.Shipping.ShippingFee(it) + it.AmazonPrice
    }
    return total
}


