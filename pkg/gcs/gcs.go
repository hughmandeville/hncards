package gcs

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"cloud.google.com/go/storage"
)

const (
	contentType  = "application/json; charset=UTF-8"
	cacheControl = "max-age=60"
)

// GCSClient is a client For GCS
type GCSClient struct {
	Bucket       string
	ContentType  string
	CacheControl string
}

func NewGCSClient(bucket string) *GCSClient {
	return &GCSClient{
		Bucket:       bucket,
		ContentType:  contentType,
		CacheControl: cacheControl,
	}
}

// Write file to GCS.
func (c *GCSClient) Store(data []byte, fileName string) error {
	log.Printf("About to write %s to GCS: https://storage.cloud.google.com/%s/%[1]s", fileName, c.Bucket)
	// Create GCS connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	// Connect to bucket
	bucket := client.Bucket(c.Bucket)
	// Setup the GCS object with the filename to write to
	obj := bucket.Object(fileName)

	// w implements io.Writer.
	w := obj.NewWriter(ctx)
	defer w.Close()

	w.ContentType = c.ContentType
	//w.ACL = aclList
	w.CacheControl = c.CacheControl

	// Copy file into GCS
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("failed to copy to bucket: %s", err)
	}

	// Close writer
	if err := w.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %s", err)
	}

	// Need to find GCS url
	log.Printf("Wrote %s to GCS: https://storage.cloud.google.com/%s/%[1]s", fileName, c.Bucket)

	return nil
}

// Read file contents from GCS.
func (c *GCSClient) Read(fileName string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	// Connect to bucket
	bucket := client.Bucket(c.Bucket)
	// Setup the GCS object with the filename
	obj := bucket.Object(fileName)

	// Read it
	r, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}
