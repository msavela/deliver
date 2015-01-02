package deliver

import (
	"net/http"
)

type Deliver struct {
	*Router
}

func New() *Deliver {
	d := &Deliver{}
	d.Router = NewRouter()
	return d
}

func (d *Deliver) Run(addr string) {
	if err := http.ListenAndServe(addr, d); err != nil {
		panic(err)
	}
}