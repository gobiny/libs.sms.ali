package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetUid() string {
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(9999)+10000)
}
