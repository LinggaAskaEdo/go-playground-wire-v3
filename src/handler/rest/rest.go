package rest

import (
	"github.com/go-chi/chi/v5"

	newssvc "github.com/linggaaskaedo/go-playground-wire-v3/src/module/news/service"
	productsvc "github.com/linggaaskaedo/go-playground-wire-v3/src/module/product/service"
)

type RestHandlerImpl struct {
	newssvc.NewsService
	productsvc.ProductService
}

func NewRestHandler(
	newsService newssvc.NewsService,
	productService productsvc.ProductService,
) *RestHandlerImpl {
	return &RestHandlerImpl{
		NewsService:    newsService,
		ProductService: productService,
	}
}

func (h *RestHandlerImpl) Router(r *chi.Mux) {
	// news
	r.Get("/api/news/{id}", h.FindNewsByID)

	// product
	r.Get("/api/product/{id}", h.FindProductByID)

	// user
	r.Post("/api/user/", h.CreateUser)
}
