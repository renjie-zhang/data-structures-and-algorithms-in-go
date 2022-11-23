package search

import (
	"testing"
)

// TestUnionVertices 存在环的情况
func TestUnionVertices(t *testing.T) {
	initialise := Initialise(6)
	var data = [][]int{{0, 1}, {1, 2}, {1, 3},
		{2, 4}, {3, 4}, {2, 5}}
	for i := 0; i < len(data); i++ {
		var x = data[i][0]
		var y = data[i][1]
		if UnionVertices(x, y, initialise) {
			t.Log("cycle detected !")
			return
		}
	}
	t.Log("no cycle !")
}

// TestUnionVertices2  不存在环的情况
func TestUnionVertices2(t *testing.T) {
	initialise := Initialise(6)
	var data = [][]int{{0, 1}, {1, 2}, {1, 3},
		{3, 4}, {2, 5}}
	for i := 0; i < len(data); i++ {
		var x = data[i][0]
		var y = data[i][1]
		if UnionVertices(x, y, initialise) {
			t.Log("cycle detected !")
			return
		}
	}
	t.Log("no cycle !")
}
