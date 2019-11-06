package main

import (
	"bytes"
	"testing"
)

func checkError(t *testing.T, want, got int) {
	if got != want {
		t.Errorf("Wanted %d, got %d.", want, got)
	}
}

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	want := 4
	got := count(b, "words")
	checkError(t, want, got)
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2\nline2\nline3 word3")
	want := 3
	got := count(b, "lines")
	checkError(t, want, got)
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("123456789")
	want := 9
	got := count(b, "bytes")
	checkError(t, want, got)
}

func TestCountRunes(t *testing.T) {
	b := bytes.NewBufferString("ΑΒΓΔΕΖΗΘ")
	want := 8
	got := count(b, "runes")
	checkError(t, want, got)
}
