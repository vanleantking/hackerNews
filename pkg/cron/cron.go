package cron

import (
	"errors"
	"strings"
	"time"

	"github.com/go-co-op/gocron/v2"
)

var (
	defaultLocationGoCron = "Asia/Ho_Chi_Minh"
)

type GoCron interface {
	GetCron() gocron.Scheduler
	AddJobDurationName(duration time.Duration, handler func(), name string) error
	AddJobSpecName(spec string, handler func(), name string) error
	AddJobOneName(datetime time.Time, handler func(), name string) error
	UpdateJobOneByName(name string, datetime time.Time, handler func()) error

	AddJobOneHandleParamName(datetime time.Time, task gocron.Task, name string) error
}

type goCron struct {
	cron gocron.Scheduler
}

func NewGoCron(id, loc string) (GoCron, error) {
	if loc == "" {
		loc = defaultLocationGoCron
	}

	l, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}
	optLoc := gocron.WithLocation(l)
	s, err := gocron.NewScheduler(optLoc)
	if err != nil {
		return nil, err
	}
	return &goCron{
		cron: s,
	}, nil
}

func (c *goCron) Start() {
	c.cron.Start()
}

func (c *goCron) Stop() error {
	return c.cron.Shutdown()
}

func (c *goCron) GetCron() gocron.Scheduler {
	return c.cron
}

func (c *goCron) AddJobDurationName(duration time.Duration, handler func(), name string) error {
	optName := gocron.WithName(name)
	_, err := c.cron.NewJob(gocron.DurationJob(duration), gocron.NewTask(handler), optName)
	return err
}

func (c *goCron) AddJobSpecName(spec string, handler func(), name string) error {
	spec = strings.TrimSpace(spec)
	lenSpec := len(strings.Split(spec, " "))
	if lenSpec < 5 || lenSpec > 6 {
		return errors.New("spec expected exactly 5 fields or 6 fields (second)")
	}
	withSecond := false
	if lenSpec == 6 {
		withSecond = true
	}
	optName := gocron.WithName(name)
	_, err := c.cron.NewJob(gocron.CronJob(spec, withSecond), gocron.NewTask(handler), optName)
	return err
}

func (c *goCron) AddJobOneName(datetime time.Time, handler func(), name string) error {
	optName := gocron.WithName(name)
	_, err := c.cron.NewJob(gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(datetime)), gocron.NewTask(handler), optName)
	return err
}

func (c *goCron) AddJobOneHandleParamName(datetime time.Time, task gocron.Task, name string) error {
	optName := gocron.WithName(name)
	_, err := c.cron.NewJob(gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(datetime)), task, optName)
	return err
}

/**
 * Update JobOne
 * 1 get list jobs
 * 2 find job by name
 * 3 create jon if name not found
 */
func (c *goCron) UpdateJobOneByName(name string, datetime time.Time, handler func()) error {
	list := c.cron.Jobs()
	for _, job := range list {
		if job.Name() == name {
			optName := gocron.WithName(name)
			_, err := c.cron.Update(job.ID(), gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(datetime)), gocron.NewTask(handler), optName)
			return err
		}
	}

	//create
	return c.AddJobOneName(datetime, handler, name)
}
