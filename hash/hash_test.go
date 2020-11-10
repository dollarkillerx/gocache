package hash

import (
	"fmt"
	"hash/crc32"
	"log"
	"testing"
)

func TestHash(t *testing.T) {
	ls := []string{
		"a", "v", "c", "d", "e", "f", "g", "h", "i", "j", "k",
	}
	for _, v := range ls {
		log.Println(SimpleHashExp(v, 10))
	}
}

func TestCrc32(t *testing.T) {
	p := []byte("abcdefg")
	ieee := crc32.ChecksumIEEE(p)
	log.Println(ieee)

	hash := SimpleHash(string(p))
	log.Println(hash)
}

func TestA(t *testing.T) {
	ser := "src"
	for i := 0; i < 100; i++ {
		ieee := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%s_%d", ser, i)))
		log.Println(ieee)
	}
}
