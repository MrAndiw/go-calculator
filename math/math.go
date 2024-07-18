package math

type (
	ModuleMath interface {
		SumNumber(a, b int) int
		MinNumber(a, b int) int
		TimesNumber(a, b int) int
		DevideNumber(a, b int) int
	}

	moduleMath struct {
	}
)

func NewMath() ModuleMath {
	return &moduleMath{}
}

func (mod *moduleMath) SumNumber(a, b int) int {

	return a + b
}

func (mod *moduleMath) MinNumber(a, b int) int {

	return a - b
}

func (mod *moduleMath) TimesNumber(a, b int) int {

	return a * b
}

func (mod *moduleMath) DevideNumber(a, b int) int {

	return a / b
}
