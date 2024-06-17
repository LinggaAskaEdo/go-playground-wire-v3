package rest

import (
	"net/http"

	httpresponse "github.com/linggaaskaedo/go-playground-wire-v3/lib/http/response"
)

func (h *RestHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	// rawProductID := chi.URLParam(r, "id")
	// productID, err := strconv.Atoi(rawProductID)
	// if err != nil {
	// 	log.Err(err)
	// 	httpresponse.Err(w, err)
	// 	return
	// }

	// user, err := h.ProductService.FindProductByID(r.Context(), int64(productID))
	// if err != nil {
	// 	log.Err(err)
	// 	httpresponse.Err(w, err)
	// 	return
	// }

	httpresponse.Json(w, http.StatusCreated, "", nil)
}
