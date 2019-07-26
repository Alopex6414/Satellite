package unpack

import (
	"crypto/cipher"
	"crypto/des"
	"log"
	"sync"
)

func DESDecryptGo(src, key []byte, dest *[]byte, wg *sync.WaitGroup) (err error) {
	*dest, err = DESDecrypt(src, key)
	if err != nil {
		log.Println("Error DES Decrypt data:", err)
		wg.Done()
		return err
	}
	wg.Done()
	return err
}

func DESDecrypt(src, key []byte) (dest []byte, err error) {
	// key length should be 8
	block, err := des.NewCipher(key)
	if err != nil {
		log.Println("Error key length:", err)
		return dest, err
	}
	// calculate block size
	blockSize := block.BlockSize()
	// encrypt mode
	blockMode := cipher.NewCBCDecrypter(block,key[:blockSize])
	// create slice
	dest = make([]byte, len(src))
	// decrypt
	blockMode.CryptBlocks(dest, src)
	// delete block data
	dest = PKCS5UnPadding(dest)
	return dest, err
}

func PKCS5UnPadding(src []byte) []byte {
	// don't need do anything
	/*size := len(src)
	unpadding := int(src[size-1])
	return src[:(size - unpadding)]*/
	return src
}
