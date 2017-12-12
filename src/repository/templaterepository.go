package repository

import (
	"log"

	"github.com/MatthiasHertel/ws17_pp/src/connection"
	"github.com/MatthiasHertel/ws17_pp/src/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TemplatesRepo connection struct
type TemplatesRepo struct {
	Server   string
	Database string
}

var db_template *mgo.Database

// Represent the CollectionName
const (
	COLLECTION_TEMPLATE = "templates"
)

var config_template = connection.Config{}
var templatesRepo = TemplatesRepo{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config_template.Read()

	templatesRepo.Server = config_template.Server
	templatesRepo.Database = config_template.Database

	templatesRepo.Connect()
}

// Connect Establish a connection to database
func (m *TemplatesRepo) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db_template = session.DB(m.Database)
}

// FindAll Find list of Templates
func (m *TemplatesRepo) FindAll() ([]models.Template, error) {

	var Templates []models.Template
	err := db_template.C(COLLECTION_TEMPLATE).Find(bson.M{}).All(&Templates)
	return Templates, err
}

// FindByID Find a template by its id
func (m *TemplatesRepo) FindByID(id string) (models.Template, error) {
	var template models.Template
	err := db_template.C(COLLECTION_TEMPLATE).FindId(bson.ObjectIdHex(id)).One(&template)
	return template, err
}

// Insert a template into database
func (m *TemplatesRepo) Insert(template models.Template) error {
	err := db_template.C(COLLECTION_TEMPLATE).Insert(&template)
	return err
}

// Delete an existing template
func (m *TemplatesRepo) Delete(template models.Template) error {
	err := db_template.C(COLLECTION_TEMPLATE).RemoveId(template.ID)
	// .Remove(&template)
	return err
}

// Update an existing template
func (m *TemplatesRepo) Update(template models.Template) error {
	err := db_template.C(COLLECTION_TEMPLATE).UpdateId(template.ID, &template)
	return err
}

// FindTemplateByJobID find Templates from Job by jobID
func (m *TemplatesRepo) FindTemplateByJobID(id string) ([]models.Template, error) {
	var Template []models.Template
	err := db.C(COLLECTION_TEMPLATE).Find(bson.M{"jobid": id}).All(&Template)
	return Template, err
}
