// Copyright 2020 Cray Inc. All Rights Reserved.

package hms_s3

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"os"
	"time"
)

type ConnectionInfo struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
	Region    string
}

func (obj *ConnectionInfo) Equals(other ConnectionInfo) (equals bool) {
	if obj.Region == other.Region &&
		obj.Bucket == other.Bucket &&
		obj.Endpoint == other.Endpoint &&
		obj.SecretKey == other.SecretKey &&
		obj.AccessKey == other.AccessKey {
		equals = true
	}
	return equals
}

func NewConnectionInfo(AccessKey string, SecretKey string, Endpoint string, Bucket string,
	Region string) (c ConnectionInfo) {
	c = ConnectionInfo{
		AccessKey: AccessKey,
		SecretKey: SecretKey,
		Endpoint:  Endpoint,
		Bucket:    Bucket,
		Region:    Region,
	}
	return c
}

func LoadConnectionInfoFromEnvVars() (info ConnectionInfo, err error) {
	// There is no static default for the access and secret keys.
	info.AccessKey = os.Getenv("S3_ACCESS_KEY")
	if info.AccessKey == "" {
		err = errors.New("access key cannot be empty")
	}
	info.SecretKey = os.Getenv("S3_SECRET_KEY")
	if info.SecretKey == "" {
		err = errors.New("secret key cannot be empty")
	}
	info.Endpoint = os.Getenv("S3_ENDPOINT")
	if info.Endpoint == "" {
		err = errors.New("endpoint cannot be empty")
	}
	info.Bucket = os.Getenv("S3_BUCKET")
	if info.Bucket == "" {
		info.Bucket = "default"
	}
	// Default to "" if there is no region defined in the environment.
	info.Region = os.Getenv("S3_REGION")
	if info.Region == "" {
		info.Region = "default"
	}
	return info, err
}

type S3Client struct {
	Session *session.Session
	S3      *s3.S3
	//Service service
	ConnInfo ConnectionInfo
}

// NewS3Client only sets up the connection to S3, it *does not* test that connection. For that, call PingBucket().
func NewS3Client(info ConnectionInfo, httpClient *http.Client) (*S3Client, error) {
	var client S3Client
	var err error

	client.Session, err = session.NewSession(aws.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(info.AccessKey, info.SecretKey, "")).
		WithEndpoint(info.Endpoint).
		WithHTTPClient(httpClient).
		WithRegion(info.Region).
		WithS3ForcePathStyle(true),
	)

	if err != nil {
		return nil, fmt.Errorf("failed setting up new session: %w", err)
	}

	if httpClient != nil {
		client.Session.Config.HTTPClient = httpClient
	}

	client.S3 = s3.New(client.Session)
	client.ConnInfo = info

	return &client, nil
}

// PingBucket will test the connection to S3 a single time. If you're using this as a measure for whether S3 is
// responsive, call it in a loop looking for nil err returned.
func (client *S3Client) PingBucket() error {
	// Test connection to S3
	_, err := client.S3.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(client.ConnInfo.Bucket),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case s3.ErrCodeNoSuchBucket:
				err := fmt.Errorf("bucket %s does not exist at %s",
					client.ConnInfo.Bucket, client.ConnInfo.Endpoint)
				return err
			}
			err := fmt.Errorf("encountered error during head_bucket operation for bucket %s at %s: %w",
				client.ConnInfo.Bucket, client.ConnInfo.Endpoint, err)
			return err
		}
	}

	return nil
}

// SetBucket updates the connection info to use the newly passed in bucket name.
func (client *S3Client) SetBucket(newBucket string) {
	client.ConnInfo.Bucket = newBucket
}

// GET

func (client *S3Client) GetObjectInput(key string) *s3.GetObjectInput {
	return &s3.GetObjectInput{
		Bucket: aws.String(client.ConnInfo.Bucket),
		Key:    aws.String(key),
	}
}

func (client *S3Client) GetObject(key string) (*s3.GetObjectOutput, error) {
	return client.S3.GetObject(client.GetObjectInput(key))
}

func (client *S3Client) GetURL(key string, expire time.Duration) (string, error) {
	req, _ := client.S3.GetObjectRequest(client.GetObjectInput(key))
	urlStr, err := req.Presign(expire)

	return urlStr, err
}

// PUT

func (client *S3Client) PutObjectInput(key string, payloadBytes []byte) *s3.PutObjectInput {
	r := bytes.NewReader(payloadBytes)

	return &s3.PutObjectInput{
		Bucket: aws.String(client.ConnInfo.Bucket),
		Key:    aws.String(key),
		Body:   r,
	}
}

func (client *S3Client) PutObject(key string, payloadBytes []byte) (*s3.PutObjectOutput, error) {
	return client.S3.PutObject(client.PutObjectInput(key, payloadBytes))
}
