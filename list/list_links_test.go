package list

import "testing"

// TestHasAbsent checks that Has terminates and reports false for a value that is
// not in the list. Has walks Next() until it finds nil, which requires the list
// to be terminated at both ends.
func TestHasAbsent(t *testing.T) {
	list := New()
	list = Insert(1, list)
	list = Insert(2, list)

	if Has(99, list) {
		t.Error("[Error] Has(99, list) returned true for an absent value'")
	}
}

// TestTraversalTerminates checks that walking from First() through Next(), and
// from Last() through Prev(), both end, and that the two agree.
func TestTraversalTerminates(t *testing.T) {
	list := New()
	for i := 1; i <= 4; i++ {
		list = Insert(i, list)
	}

	var forward []int
	for item := list.First(); item != nil; item = item.Next() {
		forward = append(forward, item.Val.(int))
	}

	if len(forward) != 4 {
		t.Fatalf("[Error] forward traversal visited %d items, want 4", len(forward))
	}

	var backward []int
	for item := list.Last(); item != nil; item = item.Prev() {
		backward = append(backward, item.Val.(int))
	}

	if len(backward) != 4 {
		t.Fatalf("[Error] backward traversal visited %d items, want 4", len(backward))
	}

	for i := range forward {
		if forward[i] != backward[len(backward)-1-i] {
			t.Fatalf("[Error] traversal orders disagree: forward %v, backward %v", forward, backward)
		}
	}
}

// TestRemoveSingle checks that removing the only element leaves an empty list.
func TestRemoveSingle(t *testing.T) {
	list := New()
	list = Insert(1, list)

	list = Remove(1, list)

	if Has(1, list) {
		t.Error("[Error] Has(1, list) returned true after the only element was removed'")
	}
	if Length(list) != 0 {
		t.Errorf("[Error] Length(list) = %d after removing the only element, want 0", Length(list))
	}
	if list.First() != nil {
		t.Error("[Error] First() is not nil after the only element was removed'")
	}
	if list.Last() != nil {
		t.Error("[Error] Last() is not nil after the only element was removed'")
	}
}

// TestRemoveHeadAndLast checks that removing the first or the last element keeps
// both ends of the list correct.
func TestRemoveHeadAndLast(t *testing.T) {
	list := New()
	for i := 1; i <= 3; i++ {
		list = Insert(i, list)
	}

	list = Remove(3, list)
	if list.First().Val != 2 {
		t.Errorf("[Error] First() = %v after removing the first element, want 2", list.First().Val)
	}

	list = Remove(1, list)
	if list.Last().Val != 2 {
		t.Errorf("[Error] Last() = %v after removing the last element, want 2", list.Last().Val)
	}

	if Length(list) != 1 {
		t.Errorf("[Error] Length(list) = %d, want 1", Length(list))
	}
}

// TestInsertAfterEmptyRemove checks that Remove on an empty list releases the
// lock, so the list stays usable afterwards. Run with -timeout: the base
// implementation returns while still holding the read lock, so Insert blocks.
func TestInsertAfterEmptyRemove(t *testing.T) {
	list := New()

	list = Remove(1, list)
	list = Insert(7, list)

	if Length(list) != 1 {
		t.Errorf("[Error] Length(list) = %d after inserting into an emptied list, want 1", Length(list))
	}
	if list.First().Val != 7 {
		t.Errorf("[Error] First() = %v after inserting into an emptied list, want 7", list.First().Val)
	}
}
