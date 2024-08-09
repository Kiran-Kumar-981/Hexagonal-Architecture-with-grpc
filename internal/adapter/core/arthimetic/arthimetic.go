package arthimetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}
func (art *Adapter) Addition(a, b int) (int, error) {
	return a + b, nil
}
func (art *Adapter) Subtraction(a, b int) (int, error) {
	return a - b, nil
}
func (art *Adapter) Multiplication(a, b int) (int, error) {
	return a * b, nil
}
func (art *Adapter) Divition(a, b int) (int, error) {
	return a / b, nil
}
