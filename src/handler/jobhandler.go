package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MatthiasHertel/ws17_pp/src/models"
	"github.com/MatthiasHertel/ws17_pp/src/repository"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var jobRepository = repository.JobsRepo{}

// AllJobsEndPoint GET list of jobs
func AllJobsEndPoint(w http.ResponseWriter, r *http.Request) {
	jobs, err := jobRepository.FindAll()
	for i := range jobs {
		jobs[i].Templates, err = templateRepository.FindTemplateByJobID(jobs[i].ID.Hex())
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, jobs)
}

// FindJobEndpoint GET a job by ID
func FindJobEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if !bson.IsObjectIdHex(params["jobID"]) {
		respondWithError(w, http.StatusBadRequest, "Invalid Job ID")
		return
	}
	job, err := jobRepository.FindByID(params["jobID"])
	job.Templates, _ = templateRepository.FindTemplateByJobID(params["jobID"])
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
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
	if err := jobRepository.Insert(job); err != nil {
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
	if err := jobRepository.Update(job); err != nil {
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
	if err := jobRepository.Delete(job); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
