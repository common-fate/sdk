package uid

import "cmp"

func SortFunc(a, b UID) int {
	if a.Type == b.Type {
		return cmp.Compare(a.ID, b.ID)
	}
	return cmp.Compare(a.Type, b.Type)
}
