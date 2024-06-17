package infra

import (
	"database/sql"

	"github.com/ary82/balance/internal/classification"
	"github.com/ary82/balance/internal/cron"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCron(
	db *sql.DB,
	classifyServerAddr string,
	sseCh chan string,
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

	classifyService := classification.NewClassifyServiceClient(client)

	service := cron.NewCronService(repo, classifyService, sseCh)
	return service, nil
}
