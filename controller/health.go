package controller

import "time"

type Health struct {
}

func (ctl *Health) GetPing() string {
	time.Sleep(5 * time.Second)
	return "pong"
}
