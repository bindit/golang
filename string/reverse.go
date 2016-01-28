package string

func Reverse(s string) string {
	sArr := []byte(s)
	sLength := int(len(sArr))
	for i, j := 0, sLength - 1; i < sLength / 2; i, j = i+1, j-1 {
		sArr[i], sArr[j] = sArr[j], sArr[i]
	}

	return string(sArr)
}