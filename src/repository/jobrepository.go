package repository

import (
	"log"

	"github.com/MatthiasHertel/ws17_pp/src/connection"
	"github.com/MatthiasHertel/ws17_pp/src/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// JobsRepo connection struct
type JobsRepo struct {
	Server   string
	Database string
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
	err := db.C(COLLECTION).Insert(&job)
	return err
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
