package hash

import (
	"unicode/utf8"
)

// 基础hash 选择节点实现
func SimpleHashExp(key string, sed uint32) uint32 {
	var out uint32
	len := len(key)
	p := 0
	for {
		if p < len {
			r, size := utf8.DecodeRuneInString(key[p:])
			p += size
			out += uint32(r)
			continue
		}
		break
	}

	return out % sed
}

// 基础hash功能
func SimpleHash(key string) uint32 {
	var out uint32
	len := len(key)
	p := 0
	for {
		if p < len {
			r, size := utf8.DecodeRuneInString(key[p:])
			p += size
			out += uint32(r)
			continue
		}
		break
	}

	return out
}
