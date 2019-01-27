package gcalauth

import (
	"cloud.google.com/go/storage"
	"encoding/json"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"io/ioutil"

	"github.com/pkg/errors"

	dga "github.com/go-zen-chu/gae-fitbit-go/pkg/domain/gcalauth"
)

type cloudStorageStore struct {
	bucketName   string
	bucketHandle *storage.BucketHandle
}

func NewCloudStorageStore(bucketName string) (dga.Store, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &cloudStorageStore{
		bucketName:   bucketName,
		bucketHandle: client.Bucket(bucketName),
	}, nil
}

func (css *cloudStorageStore) WriteGCalToken(token *oauth2.Token) error {
	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return errors.Wrap(err, "Error while marshaling")
	}
	w := css.bucketHandle.Object("gcal_oauth_token.json").NewWriter(context.Background())
	defer w.Close()
	w.ContentType = "application/json"
	_, err = w.Write(tokenBytes)
	if err != nil {
		return errors.Wrap(err, "Error while uploading to cloud storage")
	}
	return nil
}

func (css *cloudStorageStore) FetchGCalToken() (*oauth2.Token, error) {
	r, err := css.bucketHandle.Object("gcal_oauth_token.json").NewReader(context.Background())
	defer r.Close()
	if err != nil {
		return nil, errors.Wrap(err, "Error while creating reader of cloud storage")
	}
	tokenBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "Error while getting token from cloud storage")
	}

	var token oauth2.Token
	err = json.Unmarshal(tokenBytes, &token)
	if err != nil {
		return nil, errors.Wrap(err, "Error while unmarshaling token from file")
	}
	return &token, nil
}