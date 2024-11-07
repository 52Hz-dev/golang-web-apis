// middlewares/check-jwt.go
package middlewares

import (
	"errors"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/52Hz-dev/golang-web-apis/auth"
	"github.com/52Hz-dev/golang-web-apis/utils"
)

func CheckJwt(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := jwt.Verify(r)

		if err != nil {
			res.ERROR(w, 401, errors.New("Unauthorized"))
			return
		}

		next(w, r, ps)
	}
}
