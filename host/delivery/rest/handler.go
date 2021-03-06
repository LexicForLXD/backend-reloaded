package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/lexicforlxd/backend-reloaded/host"
	"github.com/lexicforlxd/backend-reloaded/lexicError"
	"github.com/lexicforlxd/backend-reloaded/models"
	_restUtil "github.com/lexicforlxd/backend-reloaded/util/delivery/rest"
)

type HostHandler struct {
	HostUsecase host.Usecase
}

func NewHostHandler(h host.Usecase) *chi.Mux {
	handler := &HostHandler{
		HostUsecase: h,
	}

	r := chi.NewRouter()
	r.Post("/", handler.CreateHost)
	r.With(_restUtil.Paginate).Get("/", handler.FetchHosts)

	r.Route("/{hostID}", func(r chi.Router) {
		r.Use(handler.HostCtx)
		r.Get("/", handler.ShowHost)
		r.Put("/", handler.UpdateHost)    // PUT /articles/123
		r.Delete("/", handler.DeleteHost) // DELETE /articles/123
	})
	return r
}

func (h *HostHandler) HostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var host *models.Host
		var err error
		ctx := r.Context()
		if hostID := chi.URLParam(r, "hostID"); hostID != "" {
			if host, err = h.HostUsecase.Show(ctx, hostID); err != nil {
				render.Render(w, r, _restUtil.NewErrorResponse(err))
				return
			}
		} else {
			render.Render(w, r, _restUtil.NewErrorResponse(lexicError.NewNotFoundError(errors.New("HostID not specified"))))
			return
		}

		ctx = context.WithValue(r.Context(), "host", host)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *HostHandler) CreateHost(w http.ResponseWriter, r *http.Request) {
	data := &HostPayload{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, _restUtil.NewErrorResponse(err))
		return
	}

	host := data.Host

	host, err := h.HostUsecase.Store(r.Context(), host)
	if err != nil {
		render.Render(w, r, _restUtil.NewErrorResponse(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewHostResponse(host))
}

func (h *HostHandler) FetchHosts(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	offset, _ := strconv.Atoi(r.FormValue("offset"))
	hosts, err := h.HostUsecase.Fetch(r.Context(), limit, offset)
	if err != nil {
		render.Render(w, r, _restUtil.ErrRender(err))
		return
	}

	if err := render.RenderList(w, r, NewHostListResponse(hosts)); err != nil {
		render.Render(w, r, _restUtil.ErrRender(err))
		return
	}
}

func (h *HostHandler) ShowHost(w http.ResponseWriter, r *http.Request) {
	host := r.Context().Value("host").(*models.Host)

	if err := render.Render(w, r, NewHostResponse(host)); err != nil {
		render.Render(w, r, _restUtil.ErrRender(err))
		return
	}
}

func (h *HostHandler) UpdateHost(w http.ResponseWriter, r *http.Request) {
	host := r.Context().Value("host").(*models.Host)

	//TODO Check for Host with same IP
	data := &HostPayload{Host: host}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, _restUtil.NewErrorResponse(err))
		return
	}
	host = data.Host
	host, err := h.HostUsecase.Update(r.Context(), host)
	if err != nil {
		render.Render(w, r, _restUtil.NewErrorResponse(err))
		return
	}

	if err := render.Render(w, r, NewHostResponse(host)); err != nil {
		render.Render(w, r, _restUtil.ErrRender(err))
		return
	}
}

// DeleteArticle removes an existing Article from our persistent store.
func (h *HostHandler) DeleteHost(w http.ResponseWriter, r *http.Request) {

	// Assume if we've reach this far, we can access the article
	// context because this handler is a child of the ArticleCtx
	// middleware. The worst case, the recoverer middleware will save us.
	host := r.Context().Value("host").(*models.Host)

	if err := h.HostUsecase.Delete(r.Context(), host.ID); err != nil {
		render.Render(w, r, _restUtil.NewErrorResponse(err))
		return
	}

	w.Write([]byte(fmt.Sprintf("host deleted with id %v", host.ID)))

}
