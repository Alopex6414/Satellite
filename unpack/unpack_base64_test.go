package unpack

import (
	"io/ioutil"
	"sync"
	"testing"
)

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