package mlib

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed")
	}
	if mm.len() != 0 {
		t.Error("Not Empty")
	}

}
