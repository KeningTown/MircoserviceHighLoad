package commonerr

type Error struct {
	Err string `json:"err"`
}

func New(msg string) Error {
	return Error{msg}
}
