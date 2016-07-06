package lib

func IsDigit(c int32) bool {
	return '0' <= c && c <= '9'
}

func IsSpace(c int32) bool {
	switch c {
		case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:	// 0xA0 : No-Break Space(U+00A0), 0x85 : NEXT LINE(NEL)(U+0085)
			return true
	}
	return false
}
