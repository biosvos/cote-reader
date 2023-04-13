package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateTest(t *testing.T) {
	ret, _ := GenerateTest("input.txt")
	fmt.Println(ret)
}

func TestReadAll(t *testing.T) {
	body, err := ReadInputText("input.txt")
	require.NoError(t, err)
	require.NotEmpty(t, body)
}

func TestParseTable(t *testing.T) {
	body, _ := ReadInputText("input.txt")

	table := ParseTable(body)

	require.Len(t, table.header, 5)
	require.Len(t, table.cells, 4)
}

func TestConvertSquareToCurlyBracket(t *testing.T) {
	ret := ConvertSquareToCurlyBracket("[[]]")
	require.Equal(t, "{{}}", ret)
}

func TestNewCell(t *testing.T) {
	t.Run("NumberCell", func(t *testing.T) {
		cell := NewCell("6")
		require.Equal(t, NumberCell, cell.Kind())
	})

	t.Run("StringCell", func(t *testing.T) {
		cell := NewCell(`"rest"`)
		require.Equal(t, StringCell, cell.Kind())
	})

	t.Run("ArrayStringCell", func(t *testing.T) {
		cell := NewCell(`{"rest"}`)
		require.Equal(t, ArrayCell, cell.Kind())
	})
}

func TestNewArrayCell(t *testing.T) {
	t.Run("one item array", func(t *testing.T) {
		cell := newArrayCell("{1}")

		require.Equal(t, ArrayCell, cell.Kind())
		require.Len(t, cell.innerCells, 1)
		require.Equal(t, NumberCell, cell.innerCells[0].Kind())
	})

	t.Run("one item double array", func(t *testing.T) {
		cell := newArrayCell("{{1}}")

		require.Equal(t, ArrayCell, cell.Kind())
		require.Len(t, cell.innerCells, 1)
		require.Equal(t, ArrayCell, cell.innerCells[0].Kind())
	})

	t.Run("two item array", func(t *testing.T) {
		cell := newArrayCell("{1, 2}")

		require.Equal(t, ArrayCell, cell.Kind())
		require.Len(t, cell.innerCells, 2)
		require.Equal(t, NumberCell, cell.innerCells[0].Kind())
		require.Equal(t, NumberCell, cell.innerCells[1].Kind())
	})

	t.Run("two item double array", func(t *testing.T) {
		cell := newArrayCell("{{1, 2}, {3, 4}}")

		require.Equal(t, ArrayCell, cell.Kind())
		require.Len(t, cell.innerCells, 2)
		require.Equal(t, ArrayCell, cell.innerCells[0].Kind())
		require.Equal(t, ArrayCell, cell.innerCells[1].Kind())
	})
}

// func TestSolution(t *testing.T) {
//	for _, tt := range []struct {
//		n        int
//		paths    [][]int
//		gates    []int
//		summits  []int
//		expected []int
//	}{
//		{
//			n:        6,
//			paths:    [][]int{[1, 2, 3], [2, 3, 5], [2, 4, 2], [2, 5, 4], [3, 4, 4], [4, 5, 3], [4, 6, 1], [5, 6, 1]},
//			gates:    []int{},
//			summits:  []int{},
//			expected: []int{5, 3},
//		},
//	} {
//		t.Run("", func(t *testing.T) {
//			ret := solution(tt.n, tt.paths, tt.gates, tt.summits)
//			require.Equal(t, tt.expected, ret)
//		})
//	}
// }
