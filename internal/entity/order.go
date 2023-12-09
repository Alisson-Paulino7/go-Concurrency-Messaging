package entity

import "errors"

type Order struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

// Função com ponteiro para a estrutura que recebe valores
// mas, antes de retornar a estrutura, valida se algum dos dados está vazio
// executando a função validate()
// Se todos os dados forem preenchidos, retorna order sem problemas
// Caso contrário, a estrutura se mantém vazia, mostrada na linha 26

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.Validate()
	if err != nil {
		return nil, err
	}
	return order, nil
}

// Pega o valor no endereço de memória da estrutura
// define o atributo derivativo a partir de 2 atributos fixos

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.Validate()
	if err != nil {
		return err
	}
	return nil
}

// Valida se os dados estão preenchidos
// Retorna um error quando executada, seja nulo ou não

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("ID is required")
	}

	if o.Price <= 0 {
		return errors.New("preço deve ser maior que zero")
	}

	if o.Tax <= 0 {
		return errors.New("invalid Tax")
	}
	return nil
}
