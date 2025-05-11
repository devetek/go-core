package helper

import (
	"testing"
)

func TestGenerateRandStr(t *testing.T) {
	var expectLen int = 5

	randStr := GenerateRandStr(expectLen)
	realLen := len(randStr)

	if realLen != expectLen {
		t.Errorf("GenerateRandStr got %d want %d", realLen, expectLen)
	}
}
