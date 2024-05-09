package db

import (
	"context"
	"errors"
	config2 "github.com/arung-agamani/mitsukeru-server-go/config"
	"github.com/arung-agamani/mitsukeru-server-go/utils/image"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
	"log"
)
import "github.com/aws/aws-sdk-go-v2/config"

var s3Client *s3.Client

func InitS3Client() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				viper.GetString("AWS_ACCESS_KEY_ID"),
				viper.GetString("AWS_SECRET_ACCESS_KEY"),
				"",
			),
		),
		config.WithRegion("ap-southeast-1"),
	)
	//cfg, err := config.load
	if err != nil {
		log.Fatal(err)
	}
	client := s3.NewFromConfig(cfg)
	s3Client = client
}

func GetS3Client() *s3.Client {
	return s3Client
}

func S3Upload(imageData, objectId string) (*s3.PutObjectOutput, error) {
	imgReader, err := image.Base64ToPngBuffer(imageData)
	if err != nil {
		return nil, err
	}
	out, err := GetS3Client().PutObject(context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String(config2.AppConfig.AWSBucket),
			Key:    aws.String(objectId),
			Body:   imgReader,
		})
	if err != nil {
		return nil, errors.New("couldn't upload file to S3")
	}
	return out, nil
}
