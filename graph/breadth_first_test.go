package graph

import (
	"../common"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	var nod5 common.Node = common.Node{5, false, nil}
	var nod4 common.Node = common.Node{4, false, nil}
	var nod3 common.Node = common.Node{3, false, nil}
	var nod2 common.Node = common.Node{2, false, []*common.Node{&nod4, &nod5}}
	var nod1 common.Node = common.Node{1, false, []*common.Node{&nod2, &nod3}}

	println(breadthFirstSearch(nod1, 2))
}
