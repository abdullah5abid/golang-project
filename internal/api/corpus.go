package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spongeling/admin-api/internal/auth"
	"github.com/spongeling/admin-api/internal/dao"
	"github.com/spongeling/admin-api/internal/pos"
	"github.com/spongeling/admin-api/internal/request"
	"github.com/spongeling/admin-api/internal/service"
)

type Corpus struct {
	svc *service.Service
}

func NewCorpus(svc *service.Service) *Corpus {
	return &Corpus{svc: svc}
}

func (api *Corpus) Routes(r chi.Router) {
	r.Get("/pos/{pos}/word", auth.Authenticator(api.GetWordsByPos))
	r.Post("/word/pos", auth.Authenticator(api.GetPosOfWord))
}

func (api *Corpus) GetWordsByPos(w http.ResponseWriter, r *http.Request) {
	posString := chi.URLParam(r, "pos")
	if posString == "" {
		respondErrorMessage(w, http.StatusBadRequest, "url parameter `pos` is required")
		return
	}

	posObj, err := dao.POSFromString(posString)
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	words, err := api.svc.GetWordsByPos(r.Context(), posObj)
	if err != nil {
		log.Println(err)
		respondError(w, http.StatusNotFound)
		return
	}

	var response []string
	for _, word := range words {
		response = append(response, word.Word)
	}

	respondOk(w, response)
}

// GetPosOfWord is handler for route POST /word/pos
func (api *Corpus) GetPosOfWord(w http.ResponseWriter, r *http.Request) {
	// parse and validate request
	var req request.GetPosOfWord
	err := parseRequest(r, &req)
	if err != nil {
		respondErrorMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	// call service: get pos of given word
	poss, err := api.svc.GetPosOfWord(r.Context(), req.Word)
	if err != nil {
		log.Println(err)
		respondErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}

	// build the response
	// response description: map containing the POS with POS details
	res := make(map[string]map[pos.Field]rune)
	for _, p := range poss {
		pos := p.ToPOS()
		vals := pos.Values
		vals["category"] = p.Category

		// map pos details against the pos
		res[pos.String()] = vals
	}

	// respond
	respondOk(w, res)
}
