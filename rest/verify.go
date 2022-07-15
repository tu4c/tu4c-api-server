package rest

import "net/http"

const verifyPath = "/verify"

type BadgeVerify struct {
	path string
}

func NewBadgeVerify() *BadgeVerify {
	return &BadgeVerify{
		path: verifyPath,
	}
}

func (verify *BadgeVerify) Serve(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		go verify.Get(w, req)
	case http.MethodPost:
		verify.Post(w, req)
	case http.MethodPut:
		verify.Put(w, req)
	case http.MethodDelete:
		verify.Delete(w, req)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
func (verify *BadgeVerify) Get(w http.ResponseWriter, req *http.Request) {
	//
}

func (verify *BadgeVerify) Post(w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (verify *BadgeVerify) Put(w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (verify *BadgeVerify) Delete(w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (verify *BadgeVerify) Path() string {
	return verify.path
}
