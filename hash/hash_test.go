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
		fmt.Println()
		ieee := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%s_%d", ser, i)))
		log.Println("CRC32: ", ieee)

		hash := SimpleHash(fmt.Sprintf("%s_%d", ser, i))
		log.Println("SimpleHash: ", hash)
	}
}

func TestHashC(t *testing.T) {
	dm := New(100, nil)
	dm.AddNodes("s1", "s2", "s3", "s4", "s5", "s6", "s7", "s9")
	addr, err := dm.GetNodeAddr("spsps")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(addr)

	for i := 0; i < 999; i++ {
		addr, err := dm.GetNodeAddr(fmt.Sprintf("spsps:%d", i))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(addr)
	}
}

func TestHashC2(t *testing.T) {
	dm := New(1000, SimpleHash)
	dm.AddNodes("s1", "s2", "s3", "s4", "s5", "s6", "s7", "s9")
	addr, err := dm.GetNodeAddr("spsps")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(addr)

	for i := 0; i < 999; i++ {
		addr, err := dm.GetNodeAddr(fmt.Sprintf("spsps:%d", i))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(addr)
	}
}
