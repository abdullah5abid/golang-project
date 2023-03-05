package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/spongeling/admin-api/internal/auth"
	"github.com/spongeling/admin-api/internal/dao"
	"github.com/spongeling/admin-api/internal/request"
	"github.com/spongeling/admin-api/internal/response"
	"github.com/spongeling/admin-api/internal/service"
)

type WordClass struct {
	svc *service.Service
}

func NewWordClass(svc *service.Service) *WordClass {
	return &WordClass{svc: svc}
}

func (api *WordClass) Routes(r chi.Router) {
	r.Route("/word/class", func(r chi.Router) {
		r.Get("/", auth.Authenticator(api.GetWordClass))
		r.Post("/", auth.Authenticator(api.AddWordClass))
		r.Route("/{word_class_id}", func(r chi.Router) {
			r.Patch("/", auth.Authenticator(api.UpdateWordClass))
			r.Delete("/", auth.Authenticator(api.DeleteWordClass))
		})
	})
}

// GetWordClass is handler for route GET /word/class
func (api *WordClass) GetWordClass(w http.ResponseWriter, r *http.Request) {
	// fetch the word classes
	wordClasses, err := api.svc.GetWordClass(r.Context())
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build response
	var res []response.WordClass
	for _, wc := range wordClasses {
		res = append(res, response.WordClass{
			Name:        wc.Name,
			Description: wc.Description,
			Words:       wc.Words,
		})
	}

	// respond
	respondOk(w, res)
}

// AddWordClass is handler for route POST /word/class
func (api *WordClass) AddWordClass(w http.ResponseWriter, r *http.Request) {
	// parse and validate request
	var req request.WordClass
	err := parseRequest(r, &req)
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// build dao object
	var class = dao.Class{
		Name:        req.Name,
		Description: req.Description,
	}

	// save word class
	cId, err := api.svc.AddWordClass(r.Context(), class, req.Words)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build response
	var res = response.Response{
		Id:      uint64(cId),
		Status:  http.StatusOK,
		Message: "record created successfully",
	}

	// respond
	respondOk(w, res)
}

// UpdateWordClass is handler for route UPDATE /word/class/{word_class_id}
func (api *WordClass) UpdateWordClass(w http.ResponseWriter, r *http.Request) {
	// fetch url parameter
	WordClassId := chi.URLParam(r, "word_class_id")
	if WordClassId == "" {
		respondErrorMessage(w, http.StatusBadRequest, "url parameter `word_class_id` is required")
		return
	}

	// parse url parameter
	wcId, err := strconv.ParseUint(WordClassId, 10, 0)
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, "invalid parameter `word_class_id`")
		return
	}

	// parse and validate request
	var req request.WordClass
	err = parseRequest(r, &req)
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// build dao object
	var class = dao.Class{
		Name:        req.Name,
		Description: req.Description,
	}

	// update word class
	resId, err := api.svc.UpdateWordClass(r.Context(), uint(wcId), class, req.Words)
	if err != nil {
		log.Println(err)
		return
	}

	// build response
	var res = response.Response{
		Id:      uint64(resId),
		Status:  http.StatusOK,
		Message: "record updated successfully",
	}

	// respond
	respondOk(w, res)
}

// DeleteWordClass is handler for route DELETE /word/class/{word_class_id}
func (api *WordClass) DeleteWordClass(w http.ResponseWriter, r *http.Request) {
	// fetch url parameter
	WordClassId := chi.URLParam(r, "word_class_id")
	if WordClassId == "" {
		respondErrorMessage(w, http.StatusBadRequest, "url parameter `word_class_id` is required")
		return
	}

	// parse url parameter
	wcId, err := strconv.ParseUint(WordClassId, 10, 0)
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, "invalid parameter `word_class_id`")
		return
	}

	// delete word class
	resId, err := api.svc.DeleteWordClass(r.Context(), uint(wcId))
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build response
	var res = response.Response{
		Id:      uint64(resId),
		Status:  http.StatusOK,
		Message: "record deleted successfully",
	}

	// respond
	respondOk(w, res)
}
