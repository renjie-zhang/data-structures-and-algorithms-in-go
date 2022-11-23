package search

func Initialise(length int) []int {
	var parent = make([]int, length)
	for i := 0; i < length; i++ {
		parent[i] = -1
	}
	return parent
}

func FindRoot(x int, parent []int) int {
	var tmp = x
	for parent[tmp] != -1 {
		tmp = parent[tmp]
	}
	return tmp
}

func UnionVertices(x, y int, parent []int) bool {
	var temp1 = FindRoot(x, parent)
	var temp2 = FindRoot(y, parent)
	if temp2 == temp1 {
		return true
	} else {
		parent[temp1] = temp2
		return false
	}
}
