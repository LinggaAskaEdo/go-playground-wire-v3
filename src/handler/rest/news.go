package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	httpresponse "github.com/linggaaskaedo/go-playground-wire-v3/lib/http/response"
)

func (h *RestHandlerImpl) FindNewsByID(w http.ResponseWriter, r *http.Request) {
	rawNewsID := chi.URLParam(r, "id")
	newsID, err := strconv.Atoi(rawNewsID)
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	news, err := h.NewsService.FindNewsByID(r.Context(), int64(newsID))
	if err != nil {
		log.Err(err)
		httpresponse.Err(w, err)
		return
	}

	httpresponse.Json(w, http.StatusOK, "", news)
}
