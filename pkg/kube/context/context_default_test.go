package context

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestDefaultContext(t *testing.T) {
	ctx := DefaultContext()
	ctx1 := DefaultContext()

	assert.Equal(t, &ctx, &ctx1)
}
