package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/MatthiasHertel/ws17_pp/src/models"
	"github.com/MatthiasHertel/ws17_pp/src/repository"
	"github.com/gorilla/mux"
	"github.com/levigross/grequests"
	minio "github.com/minio/minio-go"
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

// Submit an existing Job
func SubmitJobEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("TEST")
	job, err := jobRepository.FindByID(params["jobID"])

	if err != nil {
		fmt.Println("TEST")
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	var buffer bytes.Buffer
	buffer.WriteString("successesfully submit job with id: ")
	buffer.WriteString(job.ID.Hex())

	// TODO fetch Template json from minio
	endpoint := "localhost:9001"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	bucketName := job.ID.Hex()

	// _, err = minioClient.FGetObject(bucketName, job.ID.Hex(), "/5a7079f19ebea47c0be898a9/5a7079f19ebea47c0be898a9", minio.GetObjectOptions{})
	reader, err := minioClient.GetObject(bucketName, "5a71ec91ce2236616cd3a044.json", minio.GetObjectOptions{})
	if err != nil {
		fmt.Println("TEST")
		fmt.Println(err)
		return
	}

	localFile, err := os.Create("job-temp.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer localFile.Close()

	stat, err := reader.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := io.CopyN(localFile, reader, stat.Size); err != nil {
		log.Fatalln(err)
	}

	// TODO submit Job to nomad via put request
	// curl -X PUT -d @fib-ex.nomad.json http://127.0.0.1:4646/v1/jobs

	fd, err := grequests.FileUploadFromDisk("job-temp.json")

	if err != nil {
		log.Println("Unable to open file: ", err)
	}

	// This will upload the file as a multipart mime request
	resp, err := grequests.Put("http://127.0.0.1:4646/v1/jobs",
		&grequests.RequestOptions{
			Files: fd,
		})

	if err != nil {
		log.Println("Unable to make request", resp.Error)
	}

	if resp.Ok != true {
		log.Println("Request did not return OK")
	}
	fmt.Println(resp.String())

	os.Remove("job-temp.json")

	respondWithJSON(w, http.StatusOK, map[string]string{"result": buffer.String()})
}
