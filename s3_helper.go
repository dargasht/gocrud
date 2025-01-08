package gocrud

import (
	"bytes"
	"context"
	"io"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// This is for setting up the db
// call this in the main file and set it to gocrud.GoCRUDConfig Object
func SetupS3Client(accessKey, secretKey, endpoint string) *s3.Client {

	config, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("iran"))

	if err != nil {
		log.Println(err)
	}

	config.Credentials = aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
		}, nil
	})

	config.BaseEndpoint = aws.String(endpoint)

	return s3.NewFromConfig(config)
}

// UploadToS3 uploads a file to S3
func UploadToS3(
	ctx context.Context,
	client *s3.Client,
	file multipart.File,
	fileName string,
	folderName string,
) error {

	// Read the file into a byte buffer
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	fileContent := bytes.NewReader(buf.Bytes())

	// Specify the destination key in the bucket
	destinationKey := GoCRUDConfig.appName + "/" + folderName + "/" + fileName

	// Use the S3 client to upload the file
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(GoCRUDConfig.bucketName),
		Key:         aws.String(destinationKey),
		Body:        fileContent,
		ContentType: aws.String("image/jpg, image/png, image/jpeg, image/webp"),
	})
	if err != nil {
		return err
	}

	return nil
}
