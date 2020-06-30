package strutil

// SubBytesWithRuneCheck safely get sub bytes from byte array
func SubBytesWithRuneCheck(bs []byte, lengthLimit int) []byte {
	if lengthLimit <= 0 {
		return nil
	}
	bsLen := len(bs)
	max := min(bsLen,lengthLimit)
	right := max
getLast:
	for i := 1; i <= 4; i++ {
		l := getCharLength(bs[right-1])
		switch l {
		case 1:
			break getLast
		case -1:
			right--
		default:
			right--
			if right+l == max {
				right = max
			}
			break getLast
		}
	}
	return bs[:right]
}

// SubStringWithByteLimit safely get sub string with byte length limit
func SubStringWithByteLimit(s string, byteLengthLimit int) string {
	bs := SubBytesWithRuneCheck([]byte(s), byteLengthLimit)
	if bs == nil {
		return ""
	}
	return string(bs)
}

// getCharLength get utf8 word length
// reference https://en.wikipedia.org/wiki/UTF-8#Description
func getCharLength(b byte) int {
	switch {
	case b < 0b10000000:
		return 1
	case b < 0b11000000:
		return -1
	case b < 0b11100000:
		return 2
	case b < 0b11110000:
		return 3
	default:
		return 4
	}
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
