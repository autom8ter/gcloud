package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"io"
)

type Blob struct {
	strg *storage.Client
}

func NewBlob(ctx context.Context, opts ...option.ClientOption) (*Blob, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Blob{
		strg: client,
	}, nil
}

func (b *Blob) Close() {
	_ = b.strg.Close()
}

func (b *Blob) Client() *storage.Client {
	return b.strg
}

func (b *Blob) CreateBucket(ctx context.Context, name, project string, attributes *storage.BucketAttrs) error {
	// Creates a Bucket instance.
	bucket := b.strg.Bucket(name)

	// Creates the new bucket.
	if err := bucket.Create(ctx, project, attributes); err != nil {
		return err
	}
	return nil
}

func (b *Blob) DeleteBucket(ctx context.Context, name string) error {
	// Creates a Bucket instance.
	bucket := b.strg.Bucket(name)
	// Creates the new bucket.
	if err := bucket.Delete(ctx); err != nil {
		return err
	}
	return nil
}

func (b *Blob) UpdateBucket(ctx context.Context, name string, attributes storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error) {
	// Creates a Bucket instance.
	bucket := b.strg.Bucket(name)
	// Creates the new bucket.
	return bucket.Update(ctx, attributes)
}

func (b *Blob) BucketAttributes(ctx context.Context, name string, attributes storage.BucketAttrsToUpdate) (*storage.BucketAttrs, error) {
	// Creates a Bucket instance.
	bucket := b.strg.Bucket(name)
	// Creates the new bucket.
	return bucket.Attrs(ctx)
}

func (b *Blob) ObjectURL(objAttrs *storage.ObjectAttrs) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", objAttrs.Bucket, objAttrs.Name)
}

func (b *Blob) UploadObject(ctx context.Context, r io.Reader, projectID, bucket, name string, public bool) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {
	bh := b.strg.Bucket(bucket)
	// Next check if the bucket exists
	if _, err := bh.Attrs(ctx); err != nil {
		return nil, nil, err
	}

	obj := bh.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, r); err != nil {
		return nil, nil, err
	}
	if err := w.Close(); err != nil {
		return nil, nil, err
	}

	if public {
		if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return nil, nil, err
		}
	}

	attrs, err := obj.Attrs(ctx)
	return obj, attrs, err
}

func (b *Blob) AddBucketACL(ctx context.Context, bucket string, e storage.ACLEntity, role storage.ACLRole) error {
	acl := b.strg.Bucket(bucket).ACL()
	if err := acl.Set(ctx, e, role); err != nil {
		return err
	}
	return nil
}

func (b *Blob) BucketRules(ctx context.Context, bucket string) ([]storage.ACLRule, error) {
	return b.strg.Bucket(bucket).ACL().List(ctx)
}

type SignedUrlFunc func(options *storage.SignedURLOptions)

func (b *Blob) SignedURL(bucket, object string, opts ...SignedUrlFunc) (string, error) {
	r := &storage.SignedURLOptions{}
	for _, o := range opts {
		o(r)
	}
	return storage.SignedURL(bucket, object, r)
}
