package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"

	minio "github.com/minio/minio-go"
)

type JobTemplate struct {
	Param *string
	JobID *int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("templating!")

	var a JobTemplate

	a.Param = flag.String("Param", "1234", "help message for Param")
	a.JobID = flag.Int("JobID", 1234, "help message for JobID")

	flag.Parse()

	fmt.Println("Param has value: " + *a.Param)
	fmt.Println("JobID has value ", *a.JobID)

	// >>> template
	tmpl, _ := template.New("job.tmpl").ParseFiles("job.tmpl")

	// SAVE TO FILE
	// f, _ := os.Create("./job.json")
	filename := fmt.Sprintf("%v.json", *a.JobID)
	f, _ := os.Create(filename)
	_ = tmpl.Execute(f, a)

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

	// Make a new bucket called mymusic.
	// TODO fetch the consumer with uuid from kong here
	bucketName := strconv.Itoa(*a.JobID)

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
	log.Printf("Successfully created %s\n", bucketName)

	// Upload the text file
	objectName := filename
	filePath := filename
	contentType := "application/json"

	// Upload the text file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
