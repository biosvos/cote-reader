package main

import (
	"github.com/pkg/errors"
	"os"
	"strings"
)

func GenerateTest(filename string) (string, error) {
	text, err := ReadInputText(filename)
	if err != nil {
		return "", errors.WithStack(err)
	}
	table := ParseTable(text)
	var buffer strings.Builder
	buffer.WriteString("func TestSolution(t *testing.T) {\n")
	// for
	buffer.WriteString("\tfor _, tt := range []struct {\n")
	buffer.WriteString("\t\tname string\n")
	for idx := range table.header {
		buffer.WriteString("\t\t")
		buffer.WriteString(table.header[idx])
		buffer.WriteString(" ")
		buffer.WriteString(table.cells[0][idx].Type())
		buffer.WriteString("\n")
	}
	buffer.WriteString("\t}{\n")

	// examples
	for _, row := range table.cells {
		buffer.WriteString("\t\t{")
		buffer.WriteString("\n")
		buffer.WriteString("\t\t\tname:\"\",\n")
		for idx, cell := range row {
			buffer.WriteString("\t\t\t")
			buffer.WriteString(table.header[idx])
			buffer.WriteString(":")
			buffer.WriteString(cell.Value())
			buffer.WriteString(",\n")
		}
		buffer.WriteString("\t\t},\n")
	}
	buffer.WriteString("\t}")
	buffer.WriteString(" {\n")

	buffer.WriteString("\t\tt.Run(tt.name, func(t *testing.T) {\n")
	buffer.WriteString("\t\t\tret := solution(")
	var arguments []string
	for idx := range table.header[:len(table.header)-1] {
		arguments = append(arguments, "tt."+table.header[idx])
	}
	arg := strings.Join(arguments, ", ")
	buffer.WriteString(arg)
	buffer.WriteString(")\n")
	buffer.WriteString("\t\t\trequire.Equal(t, tt.")
	buffer.WriteString(table.header[len(table.header)-1])
	buffer.WriteString(", ret)\n")
	buffer.WriteString("\t\t})\n")

	buffer.WriteString("\t}\n")
	buffer.WriteString("}\n")
	return buffer.String(), nil
}

func ReadInputText(filename string) (string, error) {
	body, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(body), nil
}

type Table struct {
	header []string
	cells  [][]Cell
}

func ParseTable(body string) *Table {
	sp := splitByLines(body)
	header := sp[0]
	var cells [][]Cell
	for _, row := range sp[1:] {
		records := splitByRecords(row)
		var rows []Cell
		for _, record := range records {
			rows = append(rows, NewCell(record))
		}
		cells = append(cells, rows)
	}

	return &Table{
		header: splitByRecords(header),
		cells:  cells,
	}
}

func splitByRecords(row string) []string {
	split := strings.Split(row, "\t")
	var ret []string
	for _, item := range split {
		trim := strings.TrimSpace(item)
		trim = ConvertSquareToCurlyBracket(trim)
		ret = append(ret, trim)
	}
	return ret
}

func splitByLines(body string) []string {
	return strings.Split(strings.ReplaceAll(body, "\r\n", "\n"), "\n")
}

func ConvertSquareToCurlyBracket(str string) string {
	str = strings.ReplaceAll(str, "[", "{")
	str = strings.ReplaceAll(str, "]", "}")
	return str
}
