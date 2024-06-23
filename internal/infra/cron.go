package infra

import (
	"crypto/tls"
	"database/sql"

	"github.com/ary82/balance/internal/cron"
	"github.com/ary82/balance/internal/post"
	"github.com/ary82/balance/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCron(
	db *sql.DB,
	mode string,
	classifyServerAddr string,
	posSseCh chan post.Post,
	negSseCh chan post.Post,
) (cron.CronService, error) {
	repo := cron.NewCronRepository(db)

	creds := credentials.NewTLS(&tls.Config{})
	if mode != "prod" {
		creds = insecure.NewCredentials()
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			creds,
		),
	}

	client, err := grpc.NewClient(classifyServerAddr, opts...)
	if err != nil {
		return nil, err
	}

	classifyService := proto.NewClassifyServiceClient(client)

	service := cron.NewCronService(repo, classifyService, posSseCh, negSseCh)
	return service, nil
}
