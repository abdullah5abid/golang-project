package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/spongeling/admin-api/internal/auth"
	"github.com/spongeling/admin-api/internal/errors"
	"github.com/spongeling/admin-api/internal/response"
	"github.com/spongeling/admin-api/internal/service"
)

type Category struct {
	svc *service.Service
}

func NewCategory(svc *service.Service) *Category {
	return &Category{svc: svc}
}

func (api *Category) Routes(r chi.Router) {
	r.Get("/category/top", auth.Authenticator(api.GetAllTopLevelCategories))
	r.Get("/category/{category_id}/subcategories", auth.Authenticator(api.GetSubCategories))
}

// GetAllTopLevelCategories is a handler for GET /category/top
func (api *Category) GetAllTopLevelCategories(w http.ResponseWriter, r *http.Request) {
	// fetch categories from category service
	categories, err := api.svc.GetAllTopLevelCategories(r.Context())
	if err != nil {
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build response
	var categoryResponses []*response.Category
	for _, category := range categories {
		categoryResponses = append(categoryResponses, &response.Category{
			Id:   category.Id,
			Name: category.Name,
		})
	}

	// respond
	respondOk(w, categoryResponses)
}

// GetSubCategories is a handler for GET /category/{category_id}/subcategories
func (api *Category) GetSubCategories(w http.ResponseWriter, r *http.Request) {
	// fetch url parameter
	categoryId := chi.URLParam(r, "category_id")
	if categoryId == "" {
		respondErrorMessage(w, http.StatusBadRequest, "url parameter `category_id` is required")
		return
	}

	// parse url parameter
	cId, err := strconv.ParseUint(categoryId, 10, 0)
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, "invalid parameter `category_id`")
		return
	}

	// fetch categories from category service
	categories, err := api.svc.GetSubCategories(r.Context(), uint(cId))
	if err != nil {
		log.Println(err)
		if errors.IsNotFound(err) {
			respondErrorMessage(w, http.StatusNotFound, "category not found")
			return
		}
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build response
	var res []*response.Category
	for _, category := range categories {
		res = append(res, &response.Category{
			Id:   category.Id,
			Name: category.Name,
		})
	}

	// respond
	respondOk(w, res)
}
