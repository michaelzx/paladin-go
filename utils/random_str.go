package utils

import (
	"math/rand"
	"time"
)

const (
	randomTypeNum   = iota // 纯数字
	randomTypeLower        // 小写字母
	randomTypeUpper        // 大写字母
	randomTypeAll          // 数字、大小写字母
)

func RandomNumStr(site int) string {
	return string(randomStr(site, randomTypeNum))
}

// 随机字符串
func randomStr(size int, randomType int) []byte {
	kind, kinds, result := randomType, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := randomType > 2 || randomType < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random kind
			kind = rand.Intn(3)
		}
		scope, base := kinds[kind][0], kinds[kind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}
