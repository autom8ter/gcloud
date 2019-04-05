package storage

import (
	"cloud.google.com/go/spanner"
	"cloud.google.com/go/spanner/admin/database/apiv1"
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
	"os"
	"time"
)

type TableMap map[string]map[string]interface{}
type TableStruct map[string]interface{}

type SQLCreate func(request *adminpb.CreateDatabaseRequest)

type SQL struct {
	admin *database.DatabaseAdminClient
	span  *spanner.Client
}

//Must set ""GCLOUD_SPANNER_DB"" in environmental variables
func NewSQL(ctx context.Context, opts ...option.ClientOption) (*SQL, error) {
	s := &SQL{}
	var err error
	var newErr error
	// Connect to the Spanner Admin API.
	s.admin, newErr = database.NewDatabaseAdminClient(ctx, opts...)
	if newErr != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	s.span, newErr = spanner.NewClient(ctx, os.Getenv("GCLOUD_SPANNER_DB"), opts...)
	if newErr != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	return s, err
}

func (s *SQL) Close() {
	_ = s.span.Close
	_ = s.admin.Close()
}

func (s *SQL) Admin() *database.DatabaseAdminClient {
	return s.admin
}

func (s *SQL) Spanner() *spanner.Client {
	return s.span
}

func (s *SQL) CreateDatabase(ctx context.Context, opts ...SQLCreate) (*database.CreateDatabaseOperation, error) {
	r := &adminpb.CreateDatabaseRequest{}
	for _, o := range opts {
		o(r)
	}
	return s.admin.CreateDatabase(ctx, r)
}

func (s *SQL) InsertTables(ctx context.Context, tables ...TableMap) (time.Time, error) {
	mx := []*spanner.Mutation{}
	for _, table := range tables {
		for k, v := range table {
			mx = append(mx, spanner.InsertOrUpdateMap(k, v))
		}
	}
	up, err := s.span.Apply(ctx, mx)
	if err != nil {
		return up, err
	}
	return up, nil
}

func (s *SQL) InsertStructs(ctx context.Context, tables ...TableStruct) (time.Time, error) {
	mx := []*spanner.Mutation{}
	var newmx *spanner.Mutation
	var err error
	var newErr error
	for _, table := range tables {
		for k, v := range table {
			newmx, newErr = spanner.InsertOrUpdateStruct(k, v)
			if newErr != nil {
				err = errors.Wrap(err, newErr.Error())
			}
			mx = append(mx, newmx)
		}
	}
	up, err := s.span.Apply(ctx, mx)
	if err != nil {
		return up, err
	}
	return up, nil
}
