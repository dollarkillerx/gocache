package list

import (
	"log"
	"testing"
	"unsafe"
)

func TestList(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	list := New()

	list.LAppend("data")
	list.RAppend("data2")
	pop, err := list.LPop()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pop)

	pop, err = list.RPop()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pop)

	pop, err = list.LPop()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(pop)
}

func TestBytes(t *testing.T) {
	a := "asdsadsad"
	log.Println(unsafe.Sizeof(a))

	a = "asdsadsadasdsadsad"
	log.Println(unsafe.Sizeof(a))

	var b []string
	log.Println(unsafe.Sizeof(b))

	b = append(b, a, "asdsa", "sadsad")
	log.Println(unsafe.Sizeof(b))
}

func TestListD(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	list := New()
	list.RAppend("a")
	list.RAppend("b")
	log.Println(list)
	list.RAppend("c")
	list.RAppend("d")
	list.RAppend("e")
	list.RAppend("f")
	log.Println(list)
	list.MoveToFront("a")
	log.Println(list)

	list.DelByIdx(1)
	list.DelByIdx(1)
	list.DelByIdx(1)
	list.DelByIdx(1)
	log.Println(list)
}
