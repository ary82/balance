package infra

import (
	"fmt"

	"github.com/ary82/balance/internal/post"
)

func Run(
	dburl string,
	mode string,
	classifyServerAddr string,
	port string,
) error {
	db, err := NewSQLDB(dburl)
	if err != nil {
		return err
	}

	posSseCh := make(chan post.Post)
	negSseCh := make(chan post.Post)

	cron, err := NewCron(
		db, mode,
		classifyServerAddr,
		posSseCh, negSseCh,
	)
	if err != nil {
		return err
	}

	server := NewFiberServer(db)
	go readCh(posSseCh, server.CurrentPositivePosts)
	go readCh(negSseCh, server.CurrentNegativePosts)

	err = cron.Start()
	if err != nil {
		return err
	}

	err = server.App.Listen(fmt.Sprintf(":%v", port))
	return err
}

func readCh(sseCh chan post.Post, post *post.Post) {
	for {
		*post = (<-sseCh)
	}
}
