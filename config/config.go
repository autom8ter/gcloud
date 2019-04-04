package config

import(
	"context"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/runtimeconfig/v1beta1"
	"cloud.google.com/go/speech/apiv1"

)
func Dial(ctx context.Context, opts ...option.ClientOption) (pb.RuntimeConfigManagerClient, func(), error) {
	conn, err := grpc.DialContext(ctx, endPoint,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(oauth.TokenSource{TokenSource: ts}),
		useragent.GRPCDialOption("runtimevar"),
	)
	if err != nil {
		return nil, nil, err
	}

	return pb.NewRuntimeConfigManagerClient(conn), func() { conn.Close() }, nil
}
