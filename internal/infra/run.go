package infra

import (
	"fmt"
)

func Run(
	dburl string,
	classifyServerAddr string,
	port string,
) error {
	db, err := NewSQLDB(dburl)
	if err != nil {
		return err
	}

	sseCh := make(chan string)

	cron, err := NewCron(db, classifyServerAddr, sseCh)
	if err != nil {
		return err
	}

	server := NewFiberServer(db, sseCh)

	err = cron.Start()
	if err != nil {
		return err
	}

	err = server.Listen(fmt.Sprintf(":%v", port))
	return err
}
