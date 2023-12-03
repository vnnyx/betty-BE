package utils

import (
	"bytes"
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/vnnyx/betty-BE/internal/config"
	"github.com/vnnyx/betty-BE/internal/enums"
)

func InArrayStr(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func InArrayInt(array []int, value int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func InArrayScope(array []enums.Scope, value enums.Scope) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func UploadToS3(key string, file string) (string, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return "", err
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(conf.AWS.Region),
		Credentials: credentials.NewStaticCredentials(
			conf.AWS.AccessKey,
			conf.AWS.SecretKey,
			"",
		),
	})
	if err != nil {
		return "", err
	}

	fileBytes, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		return "", err
	}
	svc := s3.New(sess)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(conf.AWS.S3Bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(fileBytes),
	})
	if err != nil {
		return "", err
	}
	return key, nil
}
