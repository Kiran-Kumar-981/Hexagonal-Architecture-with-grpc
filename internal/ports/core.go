package ports

type ArthimeticPort interface {
	Addition(a, b int) (int, error)
	Subtraction(a, b int) (int, error)
	Multiplication(a, b int) (int, error)
	Divition(a, b int) (int, error)
}
