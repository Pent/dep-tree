package dep_tree

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gabotechs/dep-tree/internal/graph"
)

type TestParser struct {
	Spec [][]int
}

var _ NodeParser[[]int] = &TestParser{}

func (t *TestParser) Node(id string) (*graph.Node[[]int], error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	if idInt >= len(t.Spec) {
		return nil, fmt.Errorf("%d not present in spec", idInt)
	} else {
		return graph.MakeNode(id, t.Spec[idInt]), nil
	}
}

func (t *TestParser) Deps(n *graph.Node[[]int]) ([]*graph.Node[[]int], error) {
	result := make([]*graph.Node[[]int], 0)
	for _, child := range n.Data {
		if child < 0 {
			return nil, errors.New("no negative children")
		}
		c, err := t.Node(strconv.Itoa(child))
		if err != nil {
			n.Errors = append(n.Errors, err)
		} else {
			result = append(result, c)
		}
	}
	return result, nil
}

func (t *TestParser) Display(n *graph.Node[[]int]) DisplayResult {
	return DisplayResult{Name: n.Id}
}
