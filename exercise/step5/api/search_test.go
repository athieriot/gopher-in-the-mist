package api

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestSearch(t *testing.T) {
	resp, err := Search("Wonder woman")
	if err != nil {
		t.Error("Expected search to return some movies")
	}

	assert.Equal(t, resp.TotalResults, 9)
}