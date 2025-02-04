package unpack

import (
	"bytes"
	"io/ioutil"
	. "satellite/utils"
	"sync"
	"testing"
)

func TestUnpackAES(t *testing.T) {
	srcfile := "../test/data/unpack/file_aes.txt"
	destpath := "../test/data/unpack/"
	err := UnpackAES(srcfile, destpath)
	if err != nil {
		t.Fatal("Error Unpack AES:", err)
	}
}

func TestUnpackAESOneGo(t *testing.T) {
	var wg sync.WaitGroup
	src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x72, 0x37, 0x3D, 0x2E, 0x54, 0xDD, 0xFB, 0x78, 0x8D, 0x16, 0x54, 0xDE, 0xC9, 0x4C, 0xE2, 0x49,
		0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80, 0x87, 0x83, 0xF9, 0x0B, 0x06, 0xC8, 0x05, 0x1D,
		0x81, 0x3C, 0x53, 0x06, 0x1D, 0x13, 0xCD, 0xCB, 0xEB, 0xF3, 0x88, 0xA3, 0xB8, 0x66, 0xC7, 0x0C,
		0x86, 0x06, 0xA6, 0x9B, 0xF4, 0xD9, 0x82, 0x50, 0x92, 0xE5, 0x45, 0x09, 0x47, 0x8C, 0x4D, 0xF7,
		0x75, 0x5A, 0x2B, 0x0C, 0xD3, 0xD5, 0x35, 0x47, 0xCA, 0xD7, 0x9B, 0xD3, 0xB9, 0x25, 0xB4, 0xDA,
		0xC3, 0xDD, 0x20, 0x4C, 0x49, 0x1B, 0x67, 0x6A, 0x85, 0xFE, 0xC3, 0xB8, 0x58, 0xB0, 0x95, 0x0B,
		0xDA, 0x36, 0x07, 0xA0, 0x2B, 0x9B, 0x4F, 0x61, 0x6A, 0xBC, 0x1F, 0x37, 0x57, 0xB6, 0xAF, 0x97,
		0x44, 0x3C, 0xF9, 0x14, 0x73, 0x9D, 0x61, 0xE4, 0x21, 0x69, 0x86, 0x8E, 0xAC, 0xF9, 0x22, 0x02,
		0xCC, 0x85, 0x22, 0xE0, 0x0F, 0x21, 0xA9, 0xDA, 0x4D, 0x74, 0x6D, 0x26, 0x54, 0x05, 0x55, 0xF8,
		0x6D, 0x12, 0x27, 0x0C, 0xC3, 0x61, 0x36, 0xDF}
	destpath := "../test/data/unpack/"
	hh := TUnpackAESOne{}
	hh.Name = make([]byte, 32)
	hh.Key = make([]byte, 16)
	hh.OriginSize = make([]byte, 4)
	hh.CryptSize = make([]byte, 4)

	rd := bytes.NewReader(src)
	_, err := rd.Read(hh.Name)
	if err != nil {
		t.Fatal("Error read header name:", err)
	}
	_, err = rd.Read(hh.Key)
	if err != nil {
		t.Fatal("Error read header key:", err)
	}
	_, err = rd.Read(hh.OriginSize)
	if err != nil {
		t.Fatal("Error read header origin size:", err)
	}
	_, err = rd.Read(hh.CryptSize)
	if err != nil {
		t.Fatal("Error read header crypt size:", err)
	}
	s := make([]byte, BytesToInt(hh.CryptSize))
	n, err := rd.Read(s)
	if n <= 0 {
		t.Fatal("Error read body:", err)
	}

	wg.Add(1)
	go UnpackAESOneGo(s, hh, destpath, &wg)
	wg.Wait()
}

func TestUnpackAESOne(t *testing.T) {
	src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x72, 0x37, 0x3D, 0x2E, 0x54, 0xDD, 0xFB, 0x78, 0x8D, 0x16, 0x54, 0xDE, 0xC9, 0x4C, 0xE2, 0x49,
		0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80, 0x87, 0x83, 0xF9, 0x0B, 0x06, 0xC8, 0x05, 0x1D,
		0x81, 0x3C, 0x53, 0x06, 0x1D, 0x13, 0xCD, 0xCB, 0xEB, 0xF3, 0x88, 0xA3, 0xB8, 0x66, 0xC7, 0x0C,
		0x86, 0x06, 0xA6, 0x9B, 0xF4, 0xD9, 0x82, 0x50, 0x92, 0xE5, 0x45, 0x09, 0x47, 0x8C, 0x4D, 0xF7,
		0x75, 0x5A, 0x2B, 0x0C, 0xD3, 0xD5, 0x35, 0x47, 0xCA, 0xD7, 0x9B, 0xD3, 0xB9, 0x25, 0xB4, 0xDA,
		0xC3, 0xDD, 0x20, 0x4C, 0x49, 0x1B, 0x67, 0x6A, 0x85, 0xFE, 0xC3, 0xB8, 0x58, 0xB0, 0x95, 0x0B,
		0xDA, 0x36, 0x07, 0xA0, 0x2B, 0x9B, 0x4F, 0x61, 0x6A, 0xBC, 0x1F, 0x37, 0x57, 0xB6, 0xAF, 0x97,
		0x44, 0x3C, 0xF9, 0x14, 0x73, 0x9D, 0x61, 0xE4, 0x21, 0x69, 0x86, 0x8E, 0xAC, 0xF9, 0x22, 0x02,
		0xCC, 0x85, 0x22, 0xE0, 0x0F, 0x21, 0xA9, 0xDA, 0x4D, 0x74, 0x6D, 0x26, 0x54, 0x05, 0x55, 0xF8,
		0x6D, 0x12, 0x27, 0x0C, 0xC3, 0x61, 0x36, 0xDF}
	destpath := "../test/data/unpack/"
	hh := TUnpackAESOne{}
	hh.Name = make([]byte, 32)
	hh.Key = make([]byte, 16)
	hh.OriginSize = make([]byte, 4)
	hh.CryptSize = make([]byte, 4)

	rd := bytes.NewReader(src)
	_, err := rd.Read(hh.Name)
	if err != nil {
		t.Fatal("Error read header name:", err)
	}
	_, err = rd.Read(hh.Key)
	if err != nil {
		t.Fatal("Error read header key:", err)
	}
	_, err = rd.Read(hh.OriginSize)
	if err != nil {
		t.Fatal("Error read header origin size:", err)
	}
	_, err = rd.Read(hh.CryptSize)
	if err != nil {
		t.Fatal("Error read header crypt size:", err)
	}
	s := make([]byte, BytesToInt(hh.CryptSize))
	n, err := rd.Read(s)
	if n <= 0 {
		t.Fatal("Error read body:", err)
	}

	err = UnpackAESOne(s, hh, destpath)
	if err != nil {
		t.Fatal("Error unpack crypt file:", err)
	}
}

func TestAESDecryptGo(t *testing.T) {
	var dest []byte
	var wg sync.WaitGroup
	src := []byte{0x3B, 0x1B, 0x63, 0x41, 0x08, 0xC7, 0x8B, 0x97, 0xEC, 0x0D, 0xA3, 0xE4, 0xD2, 0xCD, 0x39, 0x84}
	key := []byte("Satellite-266414")
	wg.Add(1)
	go AESDecryptGo(src, key, &dest, &wg)
	wg.Wait()
	err := ioutil.WriteFile("../test/data/unpack/file.txt", dest, 0644)
	if err != nil {
		t.Fatal("Error Write AES One:", err)
	}
}

func TestAESDecrypt(t *testing.T) {
	src := []byte{0x3B, 0x1B, 0x63, 0x41, 0x08, 0xC7, 0x8B, 0x97, 0xEC, 0x0D, 0xA3, 0xE4, 0xD2, 0xCD, 0x39, 0x84}
	key := []byte("Satellite-266414")
	r, err := AESDecrypt(src, key)
	if err != nil {
		t.Fatal("Error AES Decrypt:", err)
	}
	err = ioutil.WriteFile("../test/data/unpack/file.txt", r, 0644)
	if err != nil {
		t.Fatal("Error Write AES One:", err)
	}
}

func BenchmarkUnpackAES(b *testing.B) {
	for i := 0; i < b.N; i++ {
		srcfile := "../test/data/unpack/file_aes.txt"
		destpath := "../test/data/unpack/"
		err := UnpackAES(srcfile, destpath)
		if err != nil {
			b.Fatal("Error Unpack AES:", err)
		}
	}
}

func BenchmarkUnpackAESOneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x72, 0x37, 0x3D, 0x2E, 0x54, 0xDD, 0xFB, 0x78, 0x8D, 0x16, 0x54, 0xDE, 0xC9, 0x4C, 0xE2, 0x49,
			0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80, 0x87, 0x83, 0xF9, 0x0B, 0x06, 0xC8, 0x05, 0x1D,
			0x81, 0x3C, 0x53, 0x06, 0x1D, 0x13, 0xCD, 0xCB, 0xEB, 0xF3, 0x88, 0xA3, 0xB8, 0x66, 0xC7, 0x0C,
			0x86, 0x06, 0xA6, 0x9B, 0xF4, 0xD9, 0x82, 0x50, 0x92, 0xE5, 0x45, 0x09, 0x47, 0x8C, 0x4D, 0xF7,
			0x75, 0x5A, 0x2B, 0x0C, 0xD3, 0xD5, 0x35, 0x47, 0xCA, 0xD7, 0x9B, 0xD3, 0xB9, 0x25, 0xB4, 0xDA,
			0xC3, 0xDD, 0x20, 0x4C, 0x49, 0x1B, 0x67, 0x6A, 0x85, 0xFE, 0xC3, 0xB8, 0x58, 0xB0, 0x95, 0x0B,
			0xDA, 0x36, 0x07, 0xA0, 0x2B, 0x9B, 0x4F, 0x61, 0x6A, 0xBC, 0x1F, 0x37, 0x57, 0xB6, 0xAF, 0x97,
			0x44, 0x3C, 0xF9, 0x14, 0x73, 0x9D, 0x61, 0xE4, 0x21, 0x69, 0x86, 0x8E, 0xAC, 0xF9, 0x22, 0x02,
			0xCC, 0x85, 0x22, 0xE0, 0x0F, 0x21, 0xA9, 0xDA, 0x4D, 0x74, 0x6D, 0x26, 0x54, 0x05, 0x55, 0xF8,
			0x6D, 0x12, 0x27, 0x0C, 0xC3, 0x61, 0x36, 0xDF}
		destpath := "../test/data/unpack/"
		hh := TUnpackAESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 16)
		hh.OriginSize = make([]byte, 4)
		hh.CryptSize = make([]byte, 4)

		rd := bytes.NewReader(src)
		_, err := rd.Read(hh.Name)
		if err != nil {
			b.Fatal("Error read header name:", err)
		}
		_, err = rd.Read(hh.Key)
		if err != nil {
			b.Fatal("Error read header key:", err)
		}
		_, err = rd.Read(hh.OriginSize)
		if err != nil {
			b.Fatal("Error read header origin size:", err)
		}
		_, err = rd.Read(hh.CryptSize)
		if err != nil {
			b.Fatal("Error read header crypt size:", err)
		}
		s := make([]byte, BytesToInt(hh.CryptSize))
		n, err := rd.Read(s)
		if n <= 0 {
			b.Fatal("Error read body:", err)
		}

		wg.Add(1)
		go UnpackAESOneGo(s, hh, destpath, &wg)
		wg.Wait()
	}
}

func BenchmarkUnpackAESOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte{0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x72, 0x37, 0x3D, 0x2E, 0x54, 0xDD, 0xFB, 0x78, 0x8D, 0x16, 0x54, 0xDE, 0xC9, 0x4C, 0xE2, 0x49,
			0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x80, 0x87, 0x83, 0xF9, 0x0B, 0x06, 0xC8, 0x05, 0x1D,
			0x81, 0x3C, 0x53, 0x06, 0x1D, 0x13, 0xCD, 0xCB, 0xEB, 0xF3, 0x88, 0xA3, 0xB8, 0x66, 0xC7, 0x0C,
			0x86, 0x06, 0xA6, 0x9B, 0xF4, 0xD9, 0x82, 0x50, 0x92, 0xE5, 0x45, 0x09, 0x47, 0x8C, 0x4D, 0xF7,
			0x75, 0x5A, 0x2B, 0x0C, 0xD3, 0xD5, 0x35, 0x47, 0xCA, 0xD7, 0x9B, 0xD3, 0xB9, 0x25, 0xB4, 0xDA,
			0xC3, 0xDD, 0x20, 0x4C, 0x49, 0x1B, 0x67, 0x6A, 0x85, 0xFE, 0xC3, 0xB8, 0x58, 0xB0, 0x95, 0x0B,
			0xDA, 0x36, 0x07, 0xA0, 0x2B, 0x9B, 0x4F, 0x61, 0x6A, 0xBC, 0x1F, 0x37, 0x57, 0xB6, 0xAF, 0x97,
			0x44, 0x3C, 0xF9, 0x14, 0x73, 0x9D, 0x61, 0xE4, 0x21, 0x69, 0x86, 0x8E, 0xAC, 0xF9, 0x22, 0x02,
			0xCC, 0x85, 0x22, 0xE0, 0x0F, 0x21, 0xA9, 0xDA, 0x4D, 0x74, 0x6D, 0x26, 0x54, 0x05, 0x55, 0xF8,
			0x6D, 0x12, 0x27, 0x0C, 0xC3, 0x61, 0x36, 0xDF}
		destpath := "../test/data/unpack/"
		hh := TUnpackAESOne{}
		hh.Name = make([]byte, 32)
		hh.Key = make([]byte, 16)
		hh.OriginSize = make([]byte, 4)
		hh.CryptSize = make([]byte, 4)

		rd := bytes.NewReader(src)
		_, err := rd.Read(hh.Name)
		if err != nil {
			b.Fatal("Error read header name:", err)
		}
		_, err = rd.Read(hh.Key)
		if err != nil {
			b.Fatal("Error read header key:", err)
		}
		_, err = rd.Read(hh.OriginSize)
		if err != nil {
			b.Fatal("Error read header origin size:", err)
		}
		_, err = rd.Read(hh.CryptSize)
		if err != nil {
			b.Fatal("Error read header crypt size:", err)
		}
		s := make([]byte, BytesToInt(hh.CryptSize))
		n, err := rd.Read(s)
		if n <= 0 {
			b.Fatal("Error read body:", err)
		}

		err = UnpackAESOne(s, hh, destpath)
		if err != nil {
			b.Fatal("Error unpack crypt file:", err)
		}
	}
}

func BenchmarkAESDecryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dest []byte
		var wg sync.WaitGroup
		src := []byte{0x3B, 0x1B, 0x63, 0x41, 0x08, 0xC7, 0x8B, 0x97, 0xEC, 0x0D, 0xA3, 0xE4, 0xD2, 0xCD, 0x39, 0x84}
		key := []byte("Satellite-266414")
		wg.Add(1)
		go AESDecryptGo(src, key, &dest, &wg)
		wg.Wait()
		err := ioutil.WriteFile("../test/data/unpack/file.txt", dest, 0644)
		if err != nil {
			b.Fatal("Error Write AES One:", err)
		}
	}
}

func BenchmarkAESDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte{0x3B, 0x1B, 0x63, 0x41, 0x08, 0xC7, 0x8B, 0x97, 0xEC, 0x0D, 0xA3, 0xE4, 0xD2, 0xCD, 0x39, 0x84}
		key := []byte("Satellite-266414")
		r, err := AESDecrypt(src, key)
		if err != nil {
			b.Fatal("Error AES Decrypt:", err)
		}
		err = ioutil.WriteFile("../test/data/unpack/file.txt", r, 0644)
		if err != nil {
			b.Fatal("Error Write AES One:", err)
		}
	}
}
