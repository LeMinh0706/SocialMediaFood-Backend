package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz "

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomDescription() string {
	return RandomString(int(RandomInt(2, 5))) + " " + RandomString(int(RandomInt(2, 5))) + " " + RandomString(int(RandomInt(2, 5))) + " " + RandomString(int(RandomInt(2, 5)))
}

func RandomType() int64 {
	types := []int64{1, 3}
	n := int64(len(types))
	return types[rand.Int63n(n)]
}
