package cron

import (
	"context"
	"fmt"
	batchprocessing "hackerNewsApi/internal/delivery/cli/batchProcessing"
	"hackerNewsApi/internal/delivery/cli/subscriber"
	"hackerNewsApi/internal/domains/pubsub"
	pubsubImpl "hackerNewsApi/internal/infrastructure/pubsub"
	"hackerNewsApi/internal/infrastructure/repository/postgre"
	service "hackerNewsApi/internal/infrastructure/service/hn_api"
	"hackerNewsApi/internal/infrastructure/service/hn_api/common"
	"hackerNewsApi/internal/infrastructure/usecase"
	subImpl "hackerNewsApi/internal/infrastructure/usecase/subscriber"
	"hackerNewsApi/pkg/cron"
	"hackerNewsApi/pkg/redis"
	"hackerNewsApi/pkg/request"
	"log"

	"github.com/spf13/cobra"
)

var (
	scheduleService = "schedule-service"
	versionCore     = "1.0.0"
)

var (
	ScheduleCmd = &cobra.Command{
		Use:     "schedule",
		Short:   "schedule run cron call hn api to up-date items hn item api",
		Long:    `schedule batch call hn api to publisher redis`,
		Version: versionCore,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("--->> schedule run <<---")
			// start new cron
			cr, er := cron.NewGoCron("")
			if er != nil {
				log.Fatalf(fmt.Sprintf("Oops, something happend on start cron, %s", er.Error()))
			}
			cr.GetCron().Start()
			app := AppConfig()
			itemDetailRepo := postgre.NewItemDetailRepository(&app.Logger, app.DB.GetDb())
			listItemRepo := postgre.NewListItemRepository(&app.Logger, app.DB.GetDb())
			apiClientSrv := service.NewHNAPIClient(
				request.NewClient(),
				app.Config.HNBaseURL,
				app.Config.HNAPIVersion,
				app.Config.HNAPIFormat,
			)
			itemDetailSubHandl := subImpl.NewSubscriberHNItemUsc(itemDetailRepo)
			listItemUsc := usecase.NewListItemUsercase(listItemRepo)

			redClient, err := redis.NewRedisClient(app.Config, app.Logger)
			if err != nil {
				app.Logger.Printf("error, %s", err.Error())
				panic("z")
			}
			pubRedisClient, subRedisClient := pubsubImpl.NewRedisClient(redClient)

			subscriberProcessing := subscriber.NewSubscriberHandler(subRedisClient)

			// declare subscriber handler function to subscriber messages from batch processing has published
			handlers := map[string]pubsub.HandlerFunc{
				common.PB_EVENT_HN_ITEM_DETAIL: itemDetailSubHandl.SubscriberUpdateDetailTopStory(),
			}
			go subscriberProcessing.Subscribes(context.Background(), handlers)

			batchProcess := batchprocessing.NewBatchProcessTopStories(
				listItemUsc,
				apiClientSrv,
				pubRedisClient,
			)

			er = cr.AddJobSpecName("*/5 * * * *", func() {
				err = batchProcess.RunBatchProcess()
				if err != nil {
					fmt.Println("Process batch cron error, ", err.Error())
				}
			}, "report tiktok video")
			fmt.Println("zzzzzzzzzzzzz, AddJobSpecName ", er, cr.GetCron())

			var forever chan struct{}
			<-forever

		},
	}
)
