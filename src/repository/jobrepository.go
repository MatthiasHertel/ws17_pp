package repository

import (
	"log"

	. "github.com/MatthiasHertel/ws17_pp/src/config"
	. "github.com/MatthiasHertel/ws17_pp/src/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type JobsRepo struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "jobs"
)

var config = Config{}
var dao = JobsRepo{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database

	dao.Connect()
}

// Establish a connection to database
func (m *JobsRepo) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of Jobs
func (m *JobsRepo) FindAll() ([]Job, error) {
	var Jobs []Job
	err := db.C(COLLECTION).Find(bson.M{}).All(&Jobs)
	return Jobs, err
}

// Find a job by its id
func (m *JobsRepo) FindById(id string) (Job, error) {
	var job Job
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&job)
	return job, err
}

// Insert a job into database
func (m *JobsRepo) Insert(job Job) error {
	err := db.C(COLLECTION).Insert(&job)
	return err
}

// Delete an existing job
func (m *JobsRepo) Delete(job Job) error {
	err := db.C(COLLECTION).RemoveId(job.ID)
	// .Remove(&job)
	return err
}

// Update an existing job
func (m *JobsRepo) Update(job Job) error {
	err := db.C(COLLECTION).UpdateId(job.ID, &job)
	return err
}
