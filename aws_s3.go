package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	s3_region = "us-west-2"
	s3_bucket = "on-prem-releases"
)

// Function for error handling
func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	// Get Filename from user
	if len(os.Args) != 4 {
		exitErrorf("Filename required")
	}
	filename := os.Args[1]
	aws_access_key := os.Args[2]
	aws_secret_key := os.Args[3]
	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file", err)
	}
	defer file.Close()

	// Open AWS Session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3_region),
		Credentials: credentials.NewStaticCredentials(
			aws_access_key,
			aws_secret_key,
			"",
		),
	})

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3_bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		exitErrorf("Upload failed", err)
	}
	fmt.Printf("Successfully uploaded %q to %q\n", filename, s3_bucket)
}
