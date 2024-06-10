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
	r.Get("/api/news/{id}", h.FindNewsByID)
	r.Get("/api/product/{id}", h.FindProductByID)
	// r.Get("/products", h.GetProducts)
	// r.Get("/products/{productId}", h.GetProductByID)
	// r.Post("/products", h.CreateNewProduct)
	// r.Delete("/products/{productId}", h.DeleteProductByID)
	// r.Get("/users/{userId}", h.GetUserById)
	// r.Post("/users/login", h.UserLogin)
	// r.With(middleware.JwtVerifyRefreshToken).Post("/users/refresh", h.UserRefreshToken)
}
