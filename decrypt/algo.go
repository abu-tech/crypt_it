package decrypt

func Dcrypt(str string) string {
	decryptedStr := ""

	for _, c := range str {
		ascii := int(c)
		newChar := string(rune(ascii - 3))
		decryptedStr += newChar
	}

	return decryptedStr
}
