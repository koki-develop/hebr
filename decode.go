package hebr

import (
	"errors"
)

var decodetable = map[string][]bool{
	"א": {false, false, false, false, false},
	"ב": {false, false, false, false, true},
	"ג": {false, false, false, true, false},
	"ד": {false, false, false, true, true},
	"ה": {false, false, true, false, false},
	"ו": {false, false, true, false, true},
	"ז": {false, false, true, true, false},
	"ח": {false, false, true, true, true},
	"ט": {false, true, false, false, false},
	"י": {false, true, false, false, true},
	"ך": {false, true, false, true, false},
	"כ": {false, true, false, true, true},
	"ל": {false, true, true, false, false},
	"ם": {false, true, true, false, true},
	"מ": {false, true, true, true, false},
	"ן": {false, true, true, true, true},
	"נ": {true, false, false, false, false},
	"ס": {true, false, false, false, true},
	"ע": {true, false, false, true, false},
	"ף": {true, false, false, true, true},
	"פ": {true, false, true, false, false},
	"ץ": {true, false, true, false, true},
	"צ": {true, false, true, true, false},
	"ק": {true, false, true, true, true},
	"ר": {true, true, false, false, false},
	"ש": {true, true, false, false, true},
	"ת": {true, true, false, true, false},
	",": {true, true, false, true, true},
	".": {true, true, true, false, false},
	"!": {true, true, true, false, true},
	"?": {true, true, true, true, false},
	"-": {true, true, true, true, true},
}

func Decode(src []byte) ([]byte, error) {
	bits := []bool{}
	for _, c := range string(src) {
		b, ok := decodetable[string(c)]
		if !ok {
			return nil, errors.New("invalid character in data")
		}
		bits = append(bits, b...)
	}

	blocks := toblocks(bits, 8, false)

	bytes := make([]byte, 0, len(blocks))
	for _, block := range blocks {
		bytes = append(bytes, bitsToByte(block))
	}

	return bytes, nil
}

func bitsToByte(bits []bool) byte {
	b := byte(0)
	for i, bit := range bits {
		if bit {
			if i == 7 {
				b += 1
			} else {
				b += 2 << (6 - i)
			}
		}
	}
	return b
}
