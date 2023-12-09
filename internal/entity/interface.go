package entity

//Padrão que vai falar os métodos que tem que ter
// para persistir dados num banco de dados
type OrderRepositoryInterface interface {
	//Método que vai salvar um pedido
	Save(order *Order) error
	GetTotal() (int, error)
}
