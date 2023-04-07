package html

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedirect(t *testing.T) {
	// Arrange
	var buff bytes.Buffer
	params := RedirectPageParams{
		Title: "Title",
	}
	// Act
	err := Redirect(&buff, params)
	result := buff.String()

	// Assert
	assert.Nil(t, err)
	assert.Contains(t, result, "Title has been moved to")
}
