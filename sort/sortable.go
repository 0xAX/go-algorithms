package sort

type Sortable interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	return false
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
