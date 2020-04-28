/*
@Time : 2020/4/27 10:40 AM
*/
package service

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func InitCronJob() {
	c := cron.New()
	c.AddFunc("* * * * * *", RunEverySecond)
	c.AddFunc("1 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
	c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
	c.AddFunc("@hourly",      func() { fmt.Println("Every hour, starting an hour from now") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
	c.Start()
}
func RunEverySecond() {
	fmt.Printf("%v\n", time.Now())
}