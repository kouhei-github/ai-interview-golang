package amazon

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

type SessionAWS struct {
	Config *session.Session
}

func NewSessionAWS() (*SessionAWS, error) {
	if os.Getenv("ENVIRONMENT") == "local" {
		value := credentials.Value{AccessKeyID: os.Getenv("AWS_KEY"), SecretAccessKey: os.Getenv("AWS_SECRET")}
		credential := credentials.NewStaticCredentialsFromCreds(value)
		conf, err := session.NewSession(&aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credential,
		})
		if err != nil {
			return &SessionAWS{}, err
		}
		return &SessionAWS{Config: conf}, nil
	} else {
		conf, err := session.NewSession(&aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
		})
		if err != nil {
			return &SessionAWS{}, err
		}
		return &SessionAWS{Config: conf}, nil
	}
}
