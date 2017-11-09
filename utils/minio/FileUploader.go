package main

import (
	"log"

	minio "github.com/minio/minio-go"
)

func main() {
	// endpoint := "play.minio.io:9000"
	// accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	// secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	endpoint := "localhost:9000"
	accessKeyID := "7SAPC6M1QZPXS1T0CRNW"
	secretAccessKey := "mUjTsz+zjzcWN6yYdcaD8YurvVa7mAKd/iaphqlu"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "cpu247testbucket"
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
	objectName := "test.txt"
	filePath := "tmp/test.txt"
	contentType := "text/plain"

	// Upload the text file with FPutObject
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
