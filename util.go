package hebr

func toblocks(bits []bool, size int, pad bool) [][]bool {
	l := len(bits)/size + 1
	blocks := make([][]bool, 0, l)

	for i := 0; i < len(bits); i += size {
		end := i + size
		if len(bits) < end {
			end = len(bits)
		}
		blocks = append(blocks, bits[i:end])
	}

	if len(blocks[l-1]) < size {
		if pad {
			blocks[l-1] = padding(blocks[l-1], size)
		} else {
			return blocks[0 : l-1]
		}
	}

	return blocks
}

func padding(bits []bool, size int) []bool {
	padd := size - len(bits)
	for i := 0; i < padd; i++ {
		bits = append(bits, false)
	}
	return bits
}
