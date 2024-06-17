package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	httpresponse "github.com/linggaaskaedo/go-playground-wire-v3/lib/http/response"
)

func (h *RestHandlerImpl) FindProductByID(w http.ResponseWriter, r *http.Request) {
	rawProductID := chi.URLParam(r, "id")
	productID, err := strconv.Atoi(rawProductID)
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	product, err := h.ProductService.FindProductByID(r.Context(), int64(productID))
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "", product)
}
