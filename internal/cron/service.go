package cron

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/ary82/balance/internal/classification"
	"github.com/ary82/balance/internal/post"
	"github.com/go-co-op/gocron/v2"
)

type cronService struct {
	cronRepository  CronRepository
	classifyService classification.ClassifyServiceClient
}

func NewCronService(cronRepo CronRepository, classifyService classification.ClassifyServiceClient) CronService {
	return &cronService{
		cronRepository:  cronRepo,
		classifyService: classifyService,
	}
}

func (s *cronService) Start() error {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	_, err = scheduler.NewJob(
		gocron.DurationJob(10*time.Second),
		gocron.NewTask(s.job),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)
	if err != nil {
		return err
	}

	scheduler.Start()
	return nil
}

func (s *cronService) job() {
	start := time.Now()
	posts, err := s.cronRepository.SelectPosts(post.POST_MAPPING_NOT_ANALYSED)
	if err != nil {
		log.Println("err:", err)
		return
	}

	if len(posts) == 0 {
		log.Println("no posts to analyze")
		return
	}

	strSlice := []string{}
	for _, v := range posts {
		cleanBody := strings.ReplaceAll(v.Body, ",", "")
		cleanBody = strings.ReplaceAll(cleanBody, "\n", "")
		strSlice = append(strSlice, cleanBody)
	}

	log.Println("grpc req:", strSlice)
	res, err := s.classifyService.Classify(
		context.Background(),
		&classification.ClassifyRequest{
			Query: strSlice,
		},
	)
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Println("grpc res:", res.Result)

	resSlice := []int{}
	err = json.Unmarshal([]byte(res.Result), &resSlice)
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Println("unmarshalled res:", resSlice)

	if len(posts) != len(resSlice) {
		log.Println("err:", "len mismatch in grpc classify")
		return
	}

	for i, v := range posts {
		v.Type = resSlice[i]
	}

	err = s.cronRepository.UpdateTypesInPosts(posts)
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Println("completed, TOOK:", time.Since(start))
}
