package rest

import (
	"net/http"
	"strings"
)

const viewPath = "/badge"

type BadgeView struct {
	path string
}

func NewBadgeView() *BadgeView {
	return &BadgeView{
		path: viewPath,
	}
}

func (view *BadgeView) Serve(backend *Backend, w http.ResponseWriter, req *http.Request) {
	if ok := strings.HasPrefix(req.URL.Path, view.Path()); !ok {
		http.Error(w, "", 0)
		return
	}
	switch req.Method {
	case http.MethodGet:
		go view.Get(backend, w, req)
	case http.MethodPost:
		view.Post(backend, w, req)
	case http.MethodPut:
		view.Put(backend, w, req)
	case http.MethodDelete:
		view.Delete(backend, w, req)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (view *BadgeView) Get(backend *Backend, w http.ResponseWriter, req *http.Request) {
	target := common.DefaultPathParser(req.URL.Path)
	if target == "" {
		http.Error(w, "", 0)
		return
	}
	// backend
	// w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode()
}

func (view *BadgeView) Post(backend *Backend, w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (view *BadgeView) Put(backend *Backend, w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (view *BadgeView) Delete(backend *Backend, w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (view *BadgeView) Path() string {
	return view.path
}
