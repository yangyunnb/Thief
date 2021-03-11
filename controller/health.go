package controller

type Health struct {
}

func (ctl *Health) GetPing() string {
	return "pong"
}
