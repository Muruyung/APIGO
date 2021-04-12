package model

import (
	"time"
)

func CreateID() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
