package repository

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/MatthiasHertel/ws17_pp/src/connection"
	"github.com/MatthiasHertel/ws17_pp/src/models"
	minio "github.com/minio/minio-go"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// JobsRepo connection struct
type JobsRepo struct {
	Server   string
	Database string
}

type JobTemplate struct {
	Param string
	JobID string
}

var db *mgo.Database

// Represent the CollectionName
const (
	COLLECTION = "jobs"
)

var config = connection.Config{}
var jobsRepo = JobsRepo{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	jobsRepo.Server = config.Server
	jobsRepo.Database = config.Database

	jobsRepo.Connect()
}

// Connect Establish a connection to database
func (m *JobsRepo) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAll Find list of Jobs
func (m *JobsRepo) FindAll() ([]models.Job, error) {

	var Jobs []models.Job
	err := db.C(COLLECTION).Find(bson.M{}).All(&Jobs)
	return Jobs, err
}

// FindByID Find a job by its id
func (m *JobsRepo) FindByID(id string) (models.Job, error) {
	var job models.Job
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&job)
	return job, err
}

// Insert a job into database
func (m *JobsRepo) Insert(job models.Job) error {
	//TODO here render template with params

	var jobTmpl JobTemplate

	jobTmpl.Param = job.Param
	jobTmpl.JobID = job.ID.Hex()

	// >>> template
	tmpl, _ := template.New("job.tmpl").ParseFiles("files/templates/job.tmpl")
	filename := fmt.Sprintf("%v.json", jobTmpl.JobID)
	f, _ := os.Create(filename)

	_ = tmpl.Execute(f, jobTmpl)

	// >>> objectstore

	// save to minio

	endpoint := "localhost:9001"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	bucketName := jobTmpl.JobID

	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created new Bucket with Name:  %s\n", bucketName)

	// Upload the text file
	objectName := filename
	filePath := filename
	contentType := "application/json"

	// Upload the text file with FPutObject

	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n into Bucket: %s", objectName, n, bucketName)

	os.Remove(filename)
	// >>> database
	err_db := db.C(COLLECTION).Insert(&job)
	return err_db
}

// Delete an existing job
func (m *JobsRepo) Delete(job models.Job) error {
	err := db.C(COLLECTION).RemoveId(job.ID)
	// .Remove(&job)
	return err
}

// Update an existing job
func (m *JobsRepo) Update(job models.Job) error {
	err := db.C(COLLECTION).UpdateId(job.ID, &job)
	return err
}
