package uid

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortingTypes(t *testing.T) {
	uids := []UID{
		New("BBB", "xyz"),
		New("AAA", "xyz"),
	}

	want := []UID{
		New("AAA", "xyz"),
		New("BBB", "xyz"),
	}

	slices.SortStableFunc(uids, SortFunc)

	assert.Equal(t, want, uids)
}

func TestSortingIDs(t *testing.T) {
	uids := []UID{
		New("AAA", "xyz"),
		New("AAA", "aaa"),
	}

	want := []UID{
		New("AAA", "aaa"),
		New("AAA", "xyz"),
	}

	slices.SortStableFunc(uids, SortFunc)

	assert.Equal(t, want, uids)
}

func TestSortingIDsAndTypes(t *testing.T) {
	uids := []UID{
		New("BBB", "aaa"),
		New("AAA", "xyz"),
		New("BBB", "xyz"),
		New("AAA", "aaa"),
	}

	want := []UID{
		New("AAA", "aaa"),
		New("AAA", "xyz"),
		New("BBB", "aaa"),
		New("BBB", "xyz"),
	}

	slices.SortStableFunc(uids, SortFunc)

	assert.Equal(t, want, uids)
}
