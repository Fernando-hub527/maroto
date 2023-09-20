// nolint: dupl
package code

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type matrixCode struct {
	code   string
	prop   props.Rect
	config *config.Config
}

func NewMatrix(code string, barcodeProps ...props.Rect) core.Component {
	prop := props.Rect{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &matrixCode{
		code: code,
		prop: prop,
	}
}

func NewMatrixCol(size int, code string, ps ...props.Rect) core.Col {
	matrixCode := NewMatrix(code, ps...)
	return col.New(size).Add(matrixCode)
}

func NewMatrixRow(height float64, code string, ps ...props.Rect) core.Row {
	matrixCode := NewMatrix(code, ps...)
	c := col.New().Add(matrixCode)
	return row.New(height).Add(c)
}

func (m *matrixCode) Render(provider core.Provider, cell core.Cell) {
	provider.AddMatrixCode(m.code, cell, m.prop)
}

func (m *matrixCode) GetStructure() *tree.Node[core.Structure] {
	str := core.Structure{
		Type:  "matrixcode",
		Value: m.code,
	}

	return tree.NewNode(str)
}

func (m *matrixCode) SetConfig(config *config.Config) {
	m.config = config
}