package unpack

import (
	"bytes"
	"io/ioutil"
	. "satellite/utils"
	"sync"
	"testing"
)

func TestUnpackBase64(t *testing.T) {
	srcfile := "../test/data/unpack/file_base64.txt"
	destpath := "../test/data/unpack/"
	err := UnpackBase64(srcfile, destpath)
	if err != nil {
		t.Fatal("Error Unpack Base64:", err)
	}
}

func TestUnpackBase64OneGo(t *testing.T) {
	var wg sync.WaitGroup
	src := []byte{
		0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
		0x62, 0x47, 0x51, 0x68,
	}
	destpath := "../test/data/unpack/"
	h := TUnpackBase64One{}
	h.Name = make([]byte, 32)
	h.Size = make([]byte, 4)
	rd := bytes.NewReader(src)
	_, err := rd.Read(h.Name)
	if err != nil {
		t.Fatal("Error read header name:", err)
	}
	_, err = rd.Read(h.Size)
	if err != nil {
		t.Fatal("Error read header size:", err)
	}
	s := make([]byte, BytesToInt(h.Size))
	n, err := rd.Read(s)
	if n <= 0 {
		t.Fatal("Error read body:", err)
	}
	wg.Add(1)
	go UnpackBase64OneGo(s, h, destpath, &wg)
	wg.Wait()
}

func TestUnpackBase64One(t *testing.T) {
	src := []byte{
		0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
		0x62, 0x47, 0x51, 0x68,
	}
	destpath := "../test/data/unpack/"
	h := TUnpackBase64One{}
	h.Name = make([]byte, 32)
	h.Size = make([]byte, 4)
	rd := bytes.NewReader(src)
	_, err := rd.Read(h.Name)
	if err != nil {
		t.Fatal("Error read header name:", err)
	}
	_, err = rd.Read(h.Size)
	if err != nil {
		t.Fatal("Error read header size:", err)
	}
	s := make([]byte, BytesToInt(h.Size))
	n, err := rd.Read(s)
	if n <= 0 {
		t.Fatal("Error read body:", err)
	}
	err = UnpackBase64One(s, h, destpath)
	if err != nil {
		t.Fatal("Error unpack crypt file:", err)
	}
}

func TestBase64DecryptGo(t *testing.T) {
	var wg sync.WaitGroup
	var r string
	src := "aGVsbG8sd29ybGQh"
	dest := "hello,world!"
	wg.Add(1)
	go Base64DecryptGo(src, &r, &wg)
	wg.Wait()
	if r != dest {
		t.Errorf("Error Decrypt Base64.")
	}
	err := ioutil.WriteFile("../test/data/unpack/file.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write Base64 One:", err)
	}
}

func TestBase64Decrypt(t *testing.T) {
	src := "aGVsbG8sd29ybGQh"
	dest := "hello,world!"
	r := Base64Decrypt(src)
	if r != dest {
		t.Errorf("Error Decrypt Base64.")
	}
	err := ioutil.WriteFile("../test/data/unpack/file.txt", []byte(r), 0644)
	if err != nil {
		t.Fatal("Error Write Base64 One:", err)
	}
}

func BenchmarkUnpackBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		srcfile := "../test/data/unpack/file_base64.txt"
		destpath := "../test/data/unpack/"
		err := UnpackBase64(srcfile, destpath)
		if err != nil {
			b.Fatal("Error Unpack Base64:", err)
		}
	}
}

func BenchmarkUnpackBase64OneGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		src := []byte{
			0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
			0x62, 0x47, 0x51, 0x68,
		}
		destpath := "../test/data/unpack/"
		h := TUnpackBase64One{}
		h.Name = make([]byte, 32)
		h.Size = make([]byte, 4)
		rd := bytes.NewReader(src)
		_, err := rd.Read(h.Name)
		if err != nil {
			b.Fatal("Error read header name:", err)
		}
		_, err = rd.Read(h.Size)
		if err != nil {
			b.Fatal("Error read header size:", err)
		}
		s := make([]byte, BytesToInt(h.Size))
		n, err := rd.Read(s)
		if n <= 0 {
			b.Fatal("Error read body:", err)
		}
		wg.Add(1)
		go UnpackBase64OneGo(s, h, destpath, &wg)
		wg.Wait()
	}
}

func BenchmarkUnpackBase64One(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := []byte{
			0x66, 0x69, 0x6C, 0x65, 0x2E, 0x74, 0x78, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x10, 0x61, 0x47, 0x56, 0x73, 0x62, 0x47, 0x38, 0x73, 0x64, 0x32, 0x39, 0x79,
			0x62, 0x47, 0x51, 0x68,
		}
		destpath := "../test/data/unpack/"
		h := TUnpackBase64One{}
		h.Name = make([]byte, 32)
		h.Size = make([]byte, 4)
		rd := bytes.NewReader(src)
		_, err := rd.Read(h.Name)
		if err != nil {
			b.Fatal("Error read header name:", err)
		}
		_, err = rd.Read(h.Size)
		if err != nil {
			b.Fatal("Error read header size:", err)
		}
		s := make([]byte, BytesToInt(h.Size))
		n, err := rd.Read(s)
		if n <= 0 {
			b.Fatal("Error read body:", err)
		}
		err = UnpackBase64One(s, h, destpath)
		if err != nil {
			b.Fatal("Error unpack crypt file:", err)
		}
	}
}

func BenchmarkBase64DecryptGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		var r string
		src := "aGVsbG8sd29ybGQh"
		dest := "hello,world!"
		wg.Add(1)
		go Base64DecryptGo(src, &r, &wg)
		wg.Wait()
		if r != dest {
			b.Errorf("Error Decrypt Base64.")
		}
		err := ioutil.WriteFile("../test/data/unpack/file.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write Base64 One:", err)
		}
	}
}

func BenchmarkBase64Decrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "aGVsbG8sd29ybGQh"
		dest := "hello,world!"
		r := Base64Decrypt(src)
		if r != dest {
			b.Errorf("Error Decrypt Base64.")
		}
		err := ioutil.WriteFile("../test/data/unpack/file.txt", []byte(r), 0644)
		if err != nil {
			b.Fatal("Error Write Base64 One:", err)
		}
	}
}
