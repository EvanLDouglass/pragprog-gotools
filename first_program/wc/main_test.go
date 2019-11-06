package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	want := 4
	got := count(b)
	if got != want {
		t.Errorf("Wanted %d, got %d.", want, got)
	}
}
