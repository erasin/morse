package morse

import (
	"testing"
)

func TestEncode(t *testing.T) {
	en := "-- --- .-. ... .  -.-. --- -.. ."
	de := "MORSE CODE"
	m := Morse{Dict: MOEN}
	if got := m.Encode(de); got != en {
		t.Errorf("de='%s', expected='%s'\n", got, en)
	}
}

func TestDecode(t *testing.T) {
	en := "-- --- .-. ... .  -.-. --- -.. ."
	de := "MORSE CODE"
	m := Morse{Dict: MOEN}
	if got := m.Decode(en); got != de {
		t.Errorf("de='%s', expected='%s'\n", got, de)
	}
}
