package challenge02

import "fmt"

var ErrInputNotSupported = fmt.Errorf("input not supported")

type MiniCompiler struct {
	Value  int
	Output string
}

func NewMiniCompiler() *MiniCompiler {
	return &MiniCompiler{
		Value:  0,
		Output: "",
	}
}

func (m *MiniCompiler) ParseInput(input rune) error {
	switch input {
	case '#':
		m.Value++
	case '@':
		m.Value--
	case '*':
		m.Value *= m.Value
	case '&':
		m.Output = fmt.Sprintf("%s%d", m.Output, m.Value)
	default:
		return ErrInputNotSupported
	}

	return nil
}

func Run(input string) string {
	compiler := NewMiniCompiler()
	for _, ch := range input {
		compiler.ParseInput(ch)
	}
	return compiler.Output
}
