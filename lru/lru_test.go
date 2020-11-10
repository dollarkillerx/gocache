package lru

import (
	"fmt"
	"log"
	"testing"
)

func TestBase(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	lru := New(1)
	resp, err := lru.Get("axax")
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)

	lru.Set("a","b")
	lru.Set("a1","b")
	lru.Set("a2","b")

	log.Println(lru.Get("a1"))
	log.Println(lru.Len())
	lru.Del("a")
	log.Println(lru.Len())

	for i:=0;i<300;i++{
		lru.Set(fmt.Sprintf("d_%d",i),"xaxxaxa")
	}
	log.Println(lru.Len())
}

