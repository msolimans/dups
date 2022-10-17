package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	exists := FileExists(workingDir+"/tests/test.json")
	assert.True(t, exists)	
}