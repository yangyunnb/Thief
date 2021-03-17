package protocol

type AddMovieReq struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}
