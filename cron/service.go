package cron

import (
	"time"

	"go-base-structure/cron/jobs"
	"go-base-structure/pkg/constant"
	"go-base-structure/pkg/util/env"

	"github.com/go-co-op/gocron"
)

func Init() {
	if !env.NewEnv().GetBool(constant.CRON_ENABLED) {
		return
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	cron := gocron.NewScheduler(loc)

	cron.Every(1).Day().At("09:00").Do(jobs.Welcome)

	cron.StartAsync()
}
