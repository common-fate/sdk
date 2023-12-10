package eid

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortingTypes(t *testing.T) {
	eids := []EID{
		New("BBB", "xyz"),
		New("AAA", "xyz"),
	}

	want := []EID{
		New("AAA", "xyz"),
		New("BBB", "xyz"),
	}

	slices.SortStableFunc(eids, SortFunc)

	assert.Equal(t, want, eids)
}

func TestSortingIDs(t *testing.T) {
	eids := []EID{
		New("AAA", "xyz"),
		New("AAA", "aaa"),
	}

	want := []EID{
		New("AAA", "aaa"),
		New("AAA", "xyz"),
	}

	slices.SortStableFunc(eids, SortFunc)

	assert.Equal(t, want, eids)
}

func TestSortingIDsAndTypes(t *testing.T) {
	eids := []EID{
		New("BBB", "aaa"),
		New("AAA", "xyz"),
		New("BBB", "xyz"),
		New("AAA", "aaa"),
	}

	want := []EID{
		New("AAA", "aaa"),
		New("AAA", "xyz"),
		New("BBB", "aaa"),
		New("BBB", "xyz"),
	}

	slices.SortStableFunc(eids, SortFunc)

	assert.Equal(t, want, eids)
}
