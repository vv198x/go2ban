package osUtil

import (
	"fmt"
	"time"
)

func DateNow() string {
	y, m, d := time.Now().Date()
	return fmt.Sprintf("%d %v %d", y, m, d)
}

func TimeNow() string {
	h, m, _ := time.Now().Clock()
	return fmt.Sprintf(" %d:%d", h, m)
}
