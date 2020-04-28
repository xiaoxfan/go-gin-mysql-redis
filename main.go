/*
@Time : 2020/4/26 5:41 PM
*/

package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"time"
)

//func init() {
//	service.InitCronJob()
//}
func main() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 */1 * * * *", RunEverySecond)
	//c.AddFunc("2 * * * * *", RunEvery2Second)
	c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
func RunEverySecond() {
	fmt.Printf("%v\n", time.Now())
}
func RunEvery2Second() {
	fmt.Printf("2%v\n", time.Now())
}
func TryLock() {
	// redis.SETNX(key,expireTime)
	//
}