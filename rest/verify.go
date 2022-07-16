package rest

import (
	"net/http"
	"strings"
)

const verifyPath = "/verify"

type BadgeVerify struct {
	path string
}

func NewBadgeVerify() *BadgeVerify {
	return &BadgeVerify{
		path: verifyPath,
	}
}

func (verify *BadgeVerify) Serve(backend *Backend, w http.ResponseWriter, req *http.Request) {
	if ok := strings.HasPrefix(req.URL.Path, verify.Path()); !ok {
		http.Error(w, "", 0)
		return
	}
	switch req.Method {
	case http.MethodGet:
		go verify.Get(backend, w, req)
	case http.MethodPost:
		go verify.Post(backend, w, req)
	case http.MethodPut:
		go verify.Put(backend, w, req)
	case http.MethodDelete:
		go verify.Delete(backend, w, req)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
func (verify *BadgeVerify) Get(backend *Backend, w http.ResponseWriter, req *http.Request) {
	target := common.DefaultPathParser(req.URL.Path)
	if target == "" {
		http.Error(w, "", 0)
		return
	}
	enrollInfo, err := backend.EnrollInfo(target)
	if err != nil {
		http.Error(w, "", 0)
		return
	}
	ok, err := backend.EnrollVerify(enrollInfo)
	if !ok && err != nil {
		// error occured
		http.Error(w, "", 0)
		return
	}
	// return verify failed
}

func (verify *BadgeVerify) Post(backend *Backend, w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (verify *BadgeVerify) Put(backend *Backend, w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (verify *BadgeVerify) Delete(backend *Backend, w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (verify *BadgeVerify) Path() string {
	return verify.path
}
