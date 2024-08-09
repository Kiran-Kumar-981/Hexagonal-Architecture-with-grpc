package ports

type APIPorts interface {
	GetAddition(a, b int) (int, error)
	GetSubtraction(a, b int) (int, error)
	GetMultiplication(a, b int) (int, error)
	GetDivition(a, b int) (int, error)
}
