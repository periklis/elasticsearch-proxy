package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// func NewTinkerHandler() http.Handler {
//	mux := http.NewServeMux()

//	mux.HandleFunc("/", defaultHandler)

//	return mux
// }

type TinkerHandler struct{}

func (t *TinkerHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("handle %s, read %d bytes", r.URL.Path, len(b))

	rw.WriteHeader(http.StatusOK)
}
