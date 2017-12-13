package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MatthiasHertel/ws17_pp/src/models"
	"github.com/MatthiasHertel/ws17_pp/src/repository"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var templateRepository = repository.TemplatesRepo{}

// AllTemplatesEndPoint GET list of templates
func AllTemplatesEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if !bson.IsObjectIdHex(params["jobID"]) {
		respondWithError(w, http.StatusBadRequest, "Invalid Job ID")
		return
	}
	_, err := jobRepository.FindByID(params["jobID"])
	if err != nil {
		respondWithJSON(w, http.StatusNotFound, "No Job found")
		return
	}
	templates, err := templateRepository.FindTemplateByJobID(params["jobID"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if templates == nil {
		respondWithJSON(w, http.StatusNotFound, "No Templates found")
		return
	}
	respondWithJSON(w, http.StatusOK, templates)
}

// FindTemplateEndpoint GET a template by ID
func FindTemplateEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if !bson.IsObjectIdHex(params["templateID"]) {
		respondWithError(w, http.StatusBadRequest, "Invalid Template ID")
		return
	}
	template, err := templateRepository.FindByID(params["templateID"])
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, template)
}

// CreateTemplateEndPoint POST a new template
func CreateTemplateEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var template models.Template
	params := mux.Vars(r)

	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	template.ID = bson.NewObjectId()
	template.JobID = params["jobID"]

	if err := templateRepository.Insert(template); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, template)
}

// UpdateTemplateEndPoint PUT update an existing template
func UpdateTemplateEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var template models.Template
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := templateRepository.Update(template); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteTemplateEndPoint DELETE an existing template
func DeleteTemplateEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var template models.Template
	if err := json.NewDecoder(r.Body).Decode(&template); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := templateRepository.Delete(template); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
