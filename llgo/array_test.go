package main

import (
	"testing"
)

func TestArrayRange(t *testing.T) { checkOutputEqual(t, "arrays/range.go") }
func TestArrayIndex(t *testing.T) { checkOutputEqual(t, "arrays/index.go") }

// vim: set ft=go:
