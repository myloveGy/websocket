package utils

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var id int32 = 0

func Unique(prefix string) string {
	atomic.AddInt32(&id, 1)
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s%s%d%03d", prefix, time.Now().Format("200601021505"), id, rand.Intn(999))
}
