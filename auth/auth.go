package auth

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/option"
)

type Auth struct {
	iAM *IAM
	kys *Keys
}

func New(ctx context.Context, opts ...option.ClientOption) (*Auth, error) {
	var err error
	var newErr error
	a := &Auth{}
	a.iAM, newErr = NewIAM(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	a.kys, newErr = NewKeys(ctx, opts...)
	if err != nil {
		err = errors.Wrap(err, newErr.Error())
	}
	return a, err
}

func (a *Auth) Close() {
	a.iAM.Close()
	a.kys.Close()
}

func (a *Auth) Keys() *Keys {
	return a.kys
}

func (a *Auth) IAM() *IAM {
	return a.iAM
}
