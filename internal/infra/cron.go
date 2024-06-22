package infra

import (
	"database/sql"

	"github.com/ary82/balance/proto"
	"github.com/ary82/balance/internal/cron"
	"github.com/ary82/balance/internal/post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCron(
	db *sql.DB,
	classifyServerAddr string,
	posSseCh chan post.Post,
	negSseCh chan post.Post,
) (cron.CronService, error) {
	repo := cron.NewCronRepository(db)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
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
