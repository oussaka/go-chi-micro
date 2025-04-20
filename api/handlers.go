package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetRecordSetPost(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "data")
	services, err := store.GetRecordSetPost(data)
	if err != nil {
		log.Errorf("Unable To Fetch stats ", httphandler.Error(err).Code, services, err)
		httphandler.ErrInvalidRequest(err, "Unable To Fetch Services ")
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, httphandler.NewSuccessResponse(http.StatusOK, services))
}
