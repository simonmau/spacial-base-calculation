package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonVec3(t *testing.T) {
	jsonData := `{"x":754.2811822753854,"y":474.46299706322134,"z":113.872134}`

	resVec := JsonVector3{}

	err := json.Unmarshal([]byte(jsonData), &resVec)

	assert.Nil(t, err)

	goVec := resVec.GenVec()

	assert.InDelta(t, 754.2811822753854, goVec[0], 0.0000001)
	assert.InDelta(t, 474.46299706322134, goVec[1], 0.0000001)
	assert.InDelta(t, 113.872134, goVec[2], 0.0000001)
}
