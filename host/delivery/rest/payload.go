package rest

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/lexicforlxd/backend-reloaded/lexicError"
	"github.com/lexicforlxd/backend-reloaded/models"
)

type HostPayload struct {
	*models.Host
}

type HostListResponse []*HostPayload

// Bind will run after the unmarshalling is complete
func (h *HostPayload) Bind(r *http.Request) error {
	if h.Name == "" {
		return lexicError.NewWrongInputError(errors.New("missing required name field"))
	}

	if h.Address == "" {
		return lexicError.NewWrongInputError(errors.New("missing required address field"))
	}

	return nil
}

// Render will run before the marshalling
func (h *HostPayload) Render(w http.ResponseWriter, r *http.Request) error {
	if h != nil {
		h.Password = ""
		if h.Containers == nil {
			h.Containers = []models.Container{}
		}
	}
	return nil
}

func NewHostResponse(host *models.Host) *HostPayload {
	resp := &HostPayload{Host: host}

	return resp
}

func NewHostListResponse(hosts []*models.Host) []render.Renderer {
	list := []render.Renderer{}
	for _, host := range hosts {
		list = append(list, NewHostResponse(host))
	}
	return list
}
