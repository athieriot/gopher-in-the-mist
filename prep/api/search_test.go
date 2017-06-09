package api

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestSearch(t *testing.T) {
	result, err := Search("Wonder Woman")
	if err != nil {
		t.Error("Oops", err.Error())
	}

	assert.Equal(t, result.TotalResults, 9)
}