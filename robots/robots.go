package robots

import (
	"cloud.google.com/go/iot/apiv1"
	"context"
	"google.golang.org/api/option"
)

type Robot struct {
	cli *iot.DeviceManagerClient
}

func New(ctx context.Context, opts ...option.ClientOption) (*Robot, error) {
	r := &Robot{}
	var err error
	r.cli, err = iot.NewDeviceManagerClient(ctx, opts...)
	return r, err
}

func (r *Robot) Client() *iot.DeviceManagerClient {
	return r.cli
}

func (r *Robot) Close() {
	_ = r.cli.Close()
}
