package rest

import "net/http"

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}
