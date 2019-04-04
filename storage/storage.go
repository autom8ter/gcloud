package storage

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

type Storage struct {
	docs *Document
	blob *Blob
	sQL  *SQL
}

func New(ctx context.Context, opts ...option.ClientOption) (*Storage, error) {
	s := &Storage{}
	var err error
	var newErr error
	s.blob, newErr = NewBlob(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	s.docs, newErr = NewDocument(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	s.sQL, newErr = NewSQL(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	return s, err
}

func (s *Storage) Close() {
	s.docs.Close()
	s.blob.Close()
	s.sQL.Close()
}

func (s *Storage) Document() *Document {
	return s.docs
}

func (s *Storage) Blob() *Blob {
	return s.blob
}

func (s *Storage) SQL() *SQL {
	return s.sQL
}
