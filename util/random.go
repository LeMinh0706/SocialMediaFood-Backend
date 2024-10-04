package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

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

func RandomType() int32 {
	types := []int32{1, 3}
	n := len(types)
	return types[rand.Intn(n)]
}

func RandomEmail() string {
	return RandomString(6) + fmt.Sprint(RandomInt(1000, 9999)) + "@gmail.com"
}

func RandomGender() int32 {
	genders := []int32{0, 1}
	n := len(genders)
	return genders[rand.Intn(n)]
}

func RandomTypeImage() string {
	image := []string{".png", ".jpg"}
	n := len(image)
	return image[rand.Intn(n)]
}

func RandomImage() string {
	return fmt.Sprintf("%v", RandomInt(1700000000, 1800000000)) + RandomTypeImage()
}
