package encrypt

func Ecrypt(str string) string {
	encryptedStr := ""

	for _, c := range str {
		ascii := int(c)
		newChar := string(rune(ascii + 3))
		encryptedStr += newChar
	}

	return encryptedStr
}
