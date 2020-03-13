package main

import "testing"

func TestNegritoSqrt(t *testing.T) {
	result := NegritoSqrt("Code.education rocks!")
	correto := "<b>Code.education rocks!</b>"
	if result != correto {
	   t.Errorf("resultado incorreto: %s obtido, %s seria correto", result, correto)
	}
}
