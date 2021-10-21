package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntity(t *testing.T) {
	// arrange
	e := NewID()
	// act
	a := NewID()
	// assert
	assert.IsType(t, e, a)
}
