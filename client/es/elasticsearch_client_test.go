package es

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	testcl := NewEsclient()
	assert.NotNil(t, testcl)
}

func TestIndexClient(t *testing.T) {
	testcl := NewEsclient()
	res, err := testcl.IndexClient("testindex", map[string]string{"test": "testing"})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
