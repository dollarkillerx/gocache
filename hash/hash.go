package hash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"unicode/utf8"
)

// 基础hash 选择节点实现
func SimpleHashExp(key string, sed uint32) uint32 {
	var out uint32
	l := len(key)
	p := 0
	for {
		if p < l {
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
	l := len(key)
	p := 0
	for {
		if p < l {
			r, size := utf8.DecodeRuneInString(key[p:])
			p += size
			out += uint32(r)
			continue
		}
		break
	}

	return out
}

func Crc32Hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

type Hash func(key string) uint32

type DistributedMap struct {
	hash     Hash
	replicas int            // 每个真实节点对应的虚节点数量
	keys     []int          // hash环
	hashMap  map[int]string // 虚拟节点 于真实节点的关系  虚:真
}

func New(replicas int, hashFun Hash) *DistributedMap {
	if replicas == 0 {
		replicas = 1
	}
	if hashFun == nil {
		hashFun = Crc32Hash
	}
	return &DistributedMap{
		replicas: replicas,
		hash:     hashFun,
		keys:     []int{},
		hashMap:  map[int]string{},
	}
}

// 添加节点s
func (d *DistributedMap) AddNodes(nodes ...string) {
	for _, key := range nodes {
		for i := 0; i < d.replicas; i++ {
			hash := int(d.hash(fmt.Sprintf("%s_%d", key, i)))
			d.keys = append(d.keys, hash)
			d.hashMap[hash] = key
		}
	}

	sort.Ints(d.keys) // 排序
}

// 获取当前key在那台节点上
func (d *DistributedMap) GetNodeAddr(key string) (string, error) {
	if len(d.keys) == 0 || len(key) == 0 {
		return "", fmt.Errorf("not kyes or key == nil ")
	}

	hash := int(d.hash(key))
	idx := sort.Search(len(d.keys), func(i int) bool {
		return d.keys[i] >= hash
	})

	return d.hashMap[d.keys[idx%len(d.keys)]], nil
}
