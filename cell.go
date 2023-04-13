package main

import (
	"strings"
)

type CellKind int

const (
	NumberCell = CellKind(iota + 1)
	StringCell
	ArrayCell
)

type Cell interface {
	Kind() CellKind
	Value() string
	Type() string
}

var _ Cell = &numberCell{}
var _ Cell = &stringCell{}
var _ Cell = &arrayCell{}

type numberCell struct {
	value string
}

func (n *numberCell) Type() string {
	return "int"
}

func (n *numberCell) Kind() CellKind {
	return NumberCell
}

func (n *numberCell) Value() string {
	return n.value
}

type stringCell struct {
	value string
}

func (s *stringCell) Type() string {
	return "string"
}

func (s *stringCell) Kind() CellKind {
	return StringCell
}

func (s *stringCell) Value() string {
	return s.value
}

type arrayCell struct {
	innerCells []Cell
}

func (a *arrayCell) Type() string {
	return "[]" + a.innerCells[0].Type()
}

func newArrayCell(str string) *arrayCell {
	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")
	depth := 0
	var cells []Cell
	var builder strings.Builder
	for _, s := range str {
		switch s {
		case '{':
			builder.WriteRune(s)
			depth++
		case '}':
			builder.WriteRune(s)
			depth--
			if depth == 0 && builder.Len() > 0 {
				sub := strings.TrimSpace(builder.String())
				builder.Reset()
				cells = append(cells, NewCell(sub))
			}
		case ',':
			if depth > 0 {
				builder.WriteRune(s)
			}
			if depth == 0 && builder.Len() > 0 {
				sub := strings.TrimSpace(builder.String())
				builder.Reset()
				cells = append(cells, NewCell(sub))
			}
		default:
			builder.WriteRune(s)
		}
	}
	if builder.Len() > 0 {
		sub := strings.TrimSpace(builder.String())
		cells = append(cells, NewCell(sub))
	}
	return &arrayCell{
		innerCells: cells,
	}
}

func (a *arrayCell) Kind() CellKind {
	return ArrayCell
}

func (a *arrayCell) Value() string {
	var ret []string
	for _, cell := range a.innerCells {
		ret = append(ret, cell.Value())
	}
	return a.Type() + "{" + strings.Join(ret, ", ") + "}"
}

func NewCell(str string) Cell {
	switch str[0] {
	case '{':
		return newArrayCell(str)
	case '"':
		return &stringCell{
			value: str,
		}
	default:
		return &numberCell{
			value: str,
		}
	}
}
