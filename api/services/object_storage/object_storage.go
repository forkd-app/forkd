package object_storage

import (
	"context"
	"forkd/util"
	"net/http"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/tags"
)

type ObjectStorageService interface {
	GetUploadUrl(ctx context.Context, filename string, expiry time.Duration) (string, error)
	SetTags(ctx context.Context, name string, otags *tags.Tags) error
}

type objectStorageService struct {
	storage *minio.Client
	bucket  string
}

func (o objectStorageService) GetUploadUrl(ctx context.Context, filename string, expiry time.Duration) (string, error) {
	info, err := o.storage.PresignedPutObject(ctx, o.bucket, filename, expiry)
	if err != nil {
		return "", err
	}

	return info.String(), nil
}

func (o objectStorageService) SetTags(ctx context.Context, name string, otags *tags.Tags) error {
	return o.storage.PutObjectTagging(ctx, o.bucket, name, otags, minio.PutObjectTaggingOptions{})
}

func New(bucket string) ObjectStorageService {
	env := util.GetEnv()
	var client *minio.Client
	var err error
	if env.GetEnvironment() == util.DEV_ENV {
		client, err = minio.New("localhost:9000", &minio.Options{
			Creds: credentials.NewStaticV4(env.GetObjectStorageAccessKey(), env.GetObjectStorageSecretKey(), ""),
			Transport: &http.Transport{
				Proxy: func(r *http.Request) (*url.URL, error) {
					r.URL.Host = env.GetObjectStorageHost()
					return r.URL, nil
				},
			},
		})
	} else {
		client, err = minio.New(env.GetObjectStorageHost(), &minio.Options{
			Creds: credentials.NewStaticV4(env.GetObjectStorageAccessKey(), env.GetObjectStorageSecretKey(), ""),
		})
	}
	if err != nil {
		panic(err)
	}
	return objectStorageService{
		storage: client,
		bucket:  bucket,
	}
}
