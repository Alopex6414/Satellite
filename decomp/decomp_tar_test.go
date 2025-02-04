package decomp

import "testing"

func TestDeCompressTarGz(t *testing.T) {
	src := "../test/data/decomp/file.tar.gz"
	dest := "../test/data/decomp/"
	err := DeCompressTarGz(src, dest)
	if err != nil {
		t.Fatal("Error DeCompress Tar Gz:", err)
	}
}

func BenchmarkDeCompressTarGz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := "../test/data/decomp/file.tar.gz"
		dest := "../test/data/decomp/"
		err := DeCompressTarGz(src, dest)
		if err != nil {
			b.Fatal("Error DeCompress Tar Gz:", err)
		}
	}
}
