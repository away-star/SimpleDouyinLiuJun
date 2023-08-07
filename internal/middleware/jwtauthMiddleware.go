package middleware

import (
	"net/http"
)

type JwtAuthMiddleware struct {
}

func NewJwtAuthMiddleware() *JwtAuthMiddleware {
	return &JwtAuthMiddleware{}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		//token := r.URL.Query().Get("token")
		//_, err := util.ParseToken(token)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//}
		// Pass through to next handler
		next(w, r)
	}
}
