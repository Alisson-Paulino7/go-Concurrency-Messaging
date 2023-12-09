// Serve para cuidar de uma necessidade do sistema

package usecase

import (
	"github.com/Alisson-Paulino7/go-concurrence-messaging/internal/entity"
)

// Dados que irão entrar no sistema
type OrderInput struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

// {"id": "1", "price": 10.0, "tax": 0.1}

// Dados de saída. Uma estrutura diferente
type OrderOutput struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"total"`
}

type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPrice(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPrice {
	return &CalculateFinalPrice{
		OrderRepository: orderRepository,
	}
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	output := &OrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
	// Poderia ser assim:
	// return &OrderOutput{
	// 	ID: order.ID,
	// 	Price: order.Price,
	// 	Tax: order.Tax,
	// 	FinalPrice: order.FinalPrice,
	// }, nil
	return output, nil
}
