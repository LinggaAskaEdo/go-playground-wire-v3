package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	httpresponse "github.com/linggaaskaedo/go-playground-wire-v3/lib/http/response"
	"github.com/rs/zerolog/log"
)

func (h *RestHandlerImpl) FindProductByID(w http.ResponseWriter, r *http.Request) {
	rawProductID := chi.URLParam(r, "id")
	productID, err := strconv.Atoi(rawProductID)
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	user, err := h.ProductService.FindProductByID(r.Context(), int64(productID))
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "", user)
}
