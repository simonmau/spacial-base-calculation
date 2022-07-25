package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonVec2(t *testing.T) {
	jsonData := `{"X":754.2811822753854,"Y":474.46299706322134}`

	resVec := JsonVector2{}

	err := json.Unmarshal([]byte(jsonData), &resVec)

	assert.Nil(t, err)

	goVec := resVec.GenVec()

	assert.InDelta(t, 754.2811822753854, goVec[0], 0.00000001)
	assert.InDelta(t, 474.46299706322134, goVec[1], 0.00000001)
}
