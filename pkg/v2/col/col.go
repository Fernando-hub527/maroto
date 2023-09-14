package col

import (
	"fmt"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

const (
	defaultGridSize = 12.0
)

type col struct {
	size       int
	_type      types.DocumentType
	components []domain.Component
	rows       []domain.Row
}

func New(size int) domain.Col {
	return &col{
		_type: types.Col,
		size:  size,
	}
}

func (c *col) GetType() string {
	return c._type.String()
}

func (c *col) Add(components ...domain.Component) domain.Col {
	c.components = append(c.components, components...)
	return c
}

func (c *col) AddInner(rows ...domain.Row) domain.Col {
	c.rows = append(c.rows, rows...)
	return c
}

func (c *col) GetSize() int {
	return c.size
}

func (c *col) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(c._type),
		Value: fmt.Sprintf("%d", c.size),
	}

	node := tree.NewNode(str)

	for _, c := range c.components {
		inner := c.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (c *col) Render(provider domain.Provider, cell internal.Cell) {
	c.render(provider, cell)

	for _, component := range c.components {
		component.Render(provider, cell)
	}

	return
}

func (c *col) render(provider domain.Provider, cell internal.Cell) {
	provider.CreateCol(cell.Width, cell.Height)
}
