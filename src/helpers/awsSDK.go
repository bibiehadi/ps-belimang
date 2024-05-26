package helpers

import (
	"context"
	"log"
	"os"

	// "github.com/aws/aws-sdk-go-v2/aws/credentials"
	// "github.com/aws/aws-sdk-go-v2/aws/session"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// func newSession() (*session.Session, error) {
// 	sess, err := session.NewSession(&aws.Config{
// 		Region: os.Getenv("AWS_REGION"),
// 		Credentials: credentials.NewStaticCredentials(
// 			os.Getenv("AWS_ACCESS_KEY_ID"),
// 			os.Getenv("AWS_SECRET_ACCESS_KEY"),
// 			"",
// 		),
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return sess, nil
// }

func S3Client() *s3.Client {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return s3.NewFromConfig(awsCfg)
}
