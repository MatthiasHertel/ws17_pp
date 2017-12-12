package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MatthiasHertel/ws17_pp/src/models"
	"github.com/MatthiasHertel/ws17_pp/src/repository"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var dao = repository.JobsRepo{}

// AllJobsEndPoint GET list of jobs
func AllJobsEndPoint(w http.ResponseWriter, r *http.Request) {
	jobs, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, jobs)
}

// FindJobEndpoint GET a job by ID
func FindJobEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	job, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Job ID")
		return
	}
	respondWithJSON(w, http.StatusOK, job)
}

// CreateJobEndPoint POST a new job
func CreateJobEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var job models.Job

	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	job.ID = bson.NewObjectId()

	if err := dao.Insert(job); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, job)
}

// UpdateJobEndPoint PUT update an existing job
func UpdateJobEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var job models.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(job); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteJobEndPoint DELETE an existing job
func DeleteJobEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var job models.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(job); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
