package binary_search_tree

import "testing"

var tests = []struct {
	input  int
	output bool
}{
	{6, true},
	{16, false},
	{3, true},
}

func TestNode_Search(t *testing.T) {
	//   6
	//  /
	//3
	tree := &Node{Key: 6, Left: &Node{Key: 3}}

	for i, test := range tests {
		if res := tree.Search(test.input); res != test.output {
			t.Errorf("%d：got %v ,expected %v", i, res, test.output)
		}
	}
}

func BenchmarkNode_Search(b *testing.B) {
	tree := &Node{Key: 6}
	for i := 0; i < b.N; i++ {
		tree.Search(6)
	}
}
