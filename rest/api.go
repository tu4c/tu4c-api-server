package rest

import "net/http"

type Service interface {
	Serve(*Backend, http.ResponseWriter, *http.Request)
	Path() string
}

type RestMethod interface {
	Get(*Backend, http.ResponseWriter, *http.Request)
	Post(*Backend, http.ResponseWriter, *http.Request)
	Put(*Backend, http.ResponseWriter, *http.Request)
	Delete(*Backend, http.ResponseWriter, *http.Request)
}

type Backend interface {
}

type API struct {
	backend  *Backend
	services map[string]Service
}

func NewAPI(backend *Backend) *API {
	var services map[string]Service

	view := NewBadgeView()
	verify := NewBadgeVerify()

	services[view.Path()] = view
	services[verify.Path()] = verify
	//
	return &API{
		backend:  backend,
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
	go service.Serve(api.backend, w, req)
}

// check required. can only increase complexity
func MethodRouter(run RestMethod, backend *Backend, w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		go run.Get(backend, w, req)
	case http.MethodPost:
		go run.Post(backend, w, req)
	case http.MethodPut:
		go run.Put(backend, w, req)
	case http.MethodDelete:
		go run.Delete(backend, w, req)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
