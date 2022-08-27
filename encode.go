package hebr

import (
	"errors"
)

var encodetable = map[string]string{
	"00000": "א",
	"00001": "ב",
	"00010": "ג",
	"00011": "ד",
	"00100": "ה",
	"00101": "ו",
	"00110": "ז",
	"00111": "ח",
	"01000": "ט",
	"01001": "י",
	"01010": "ך",
	"01011": "כ",
	"01100": "ל",
	"01101": "ם",
	"01110": "מ",
	"01111": "ן",
	"10000": "נ",
	"10001": "ס",
	"10010": "ע",
	"10011": "ף",
	"10100": "פ",
	"10101": "ץ",
	"10110": "צ",
	"10111": "ק",
	"11000": "ר",
	"11001": "ש",
	"11010": "ת",
	"11011": ",",
	"11100": ".",
	"11101": "!",
	"11110": "?",
	"11111": "-",
}

func Encode(src []byte) ([]byte, error) {
	bits := tobits(src)
	blocks := toblocks(bits, 5, true)

	bytes := []byte{}
	for _, block := range blocks {
		char, ok := bitsToEncoded(block)
		if !ok {
			return nil, errors.New("encoding failed")
		}
		bytes = append(bytes, []byte(char)...)
	}

	return bytes, nil
}

func bitsToEncoded(bits []bool) (string, bool) {
	s, ok := encodetable[bitstostr(bits)]
	return s, ok
}

func bitstostr(bits []bool) string {
	b := []byte{}
	for _, bit := range bits {
		if bit {
			b = append(b, '1')
		} else {
			b = append(b, '0')
		}
	}
	return string(b)
}

func tobits(src []byte) []bool {
	bits := make([]bool, 0, len(src)*8)

	for _, c := range src {
		for i := 0; i < 8; i++ {
			bits = append(bits, (c>>(8-(i+1))&1) == 1)
		}
	}

	return bits
}
