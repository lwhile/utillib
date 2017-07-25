package strs

// DelEscape :
func DelEscape(b, escape []byte) []byte {
	if len(b) == 0 {
		return b
	}

	var r []byte
	var c int
	for i := 0; i < len(b); i++ {
		t := b[i]
		if t != '\b' {
			r = append(r, t)
			c++
		} else {
			c--
			r = r[:c]
		}
	}
	return r
}
