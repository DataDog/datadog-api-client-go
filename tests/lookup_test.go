package tests

import (
	"context"
	"testing"
)

var testMap = map[string]interface{}{
	"foo": map[string]interface{}{
		"bar": 1,
	},
	"spam": []map[string]interface{}{
		{"egg": 1},
		{"egg": 2},
	},
}

func TestLookupStringI_Map(t *testing.T) {
	assert := Assert(context.Background(), t)
	value, err := LookupStringI(testMap, "foo.bar")
	assert.Equal(nil, err)
	assert.Equal(1, value.Interface())

	_, err = LookupStringI(testMap, "foo.spam")
	assert.Equal(ErrNotFound, err)

	value, err = LookupStringI(testMap, "spam[1].egg")
	assert.Equal(nil, err)
	assert.Equal(2, value.Interface())

	_, err = LookupStringI(testMap, "spam[2].egg")
	assert.Equal(ErrNotFound, err)
}
