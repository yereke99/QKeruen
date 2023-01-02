package help

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomID(l int) (lens int) {
	s := ""
	rand.Seed(time.Now().Unix())

	for i := 0; i < l; i++ {
		s += (string)(rand.Intn(10) + 48)
	}

	lens, _ = strconv.Atoi(s)

	return
}
