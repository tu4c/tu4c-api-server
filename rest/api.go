package rest

import "net/http"

type Service interface {
	Serve(http.ResponseWriter, *http.Request)
	Path() string
}

type RestFunc interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Put(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type API struct {
	services map[string]Service
}

func NewAPI() *API {
	var services map[string]Service

	view := NewBadgeView()
	verify := NewBadgeVerify()

	services[view.Path()] = view
	services[verify.Path()] = verify
	//
	return &API{
		services: services,
	}
}

func (api *API) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// parse first path
	firstUrl := "/badge" // temp

	var service Service
	var ok bool
	if service, ok = api.services[firstUrl]; !ok {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	go service.Serve(w, req)
}
