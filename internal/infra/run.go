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
	countSseCh := make(chan post.PostCounts)

	cron, err := NewCron(
		db, mode,
		classifyServerAddr,
		posSseCh, negSseCh, countSseCh,
	)
	if err != nil {
		return err
	}

	server := NewFiberServer(db)
	go readPostCh(posSseCh, &server.CurrentPositivePosts)
	go readPostCh(negSseCh, &server.CurrentNegativePosts)
	go readCountCh(countSseCh, &server.PostsCount)

	err = cron.Start()
	if err != nil {
		return err
	}

	err = server.App.Listen(fmt.Sprintf(":%v", port))
	return err
}

func readPostCh(sseCh chan post.Post, post *post.Post) {
	for {
		*post = (<-sseCh)
	}
}

func readCountCh(countCh chan post.PostCounts, count *post.PostCounts) {
	for {
		*count = (<-countCh)
	}
}
