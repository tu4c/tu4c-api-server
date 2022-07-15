package rest

import "net/http"

const viewPath = "/badge"

type BadgeView struct {
	path string
}

func NewBadgeView() *BadgeView {
	return &BadgeView{
		path: viewPath,
	}
}

func (view *BadgeView) Serve(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		go view.Get(w, req)
	case http.MethodPost:
		view.Post(w, req)
	case http.MethodPut:
		view.Put(w, req)
	case http.MethodDelete:
		view.Delete(w, req)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (view *BadgeView) Get(w http.ResponseWriter, req *http.Request) {
	//
}

func (view *BadgeView) Post(w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (view *BadgeView) Put(w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (view *BadgeView) Delete(w http.ResponseWriter, req *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (view *BadgeView) Path() string {
	return view.path
}
