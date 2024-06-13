package main

import (
	"fmt"

	"github.com/abu-tech/crypt_it/decrypt"
	"github.com/abu-tech/crypt_it/encrypt"
)

func main() {
	encryptedStr := encrypt.Ecrypt("hello")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Dcrypt(encryptedStr))
}
