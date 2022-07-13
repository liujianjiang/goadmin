package antspool

import (
	"context"

	"log"

	"github.com/liujianjiang/goadmin/server/global"
	"github.com/panjf2000/ants/v2"
)

var taskPool *ants.Pool

//初始化协程池
func InitPools() (err error) {
	taskPool, err = ants.NewPool(global.GVA_CONFIG.AntsPool.Size, ants.WithMaxBlockingTasks(global.GVA_CONFIG.AntsPool.MaxBlockingTask), ants.WithNonblocking(false))
	if err != nil {
		log.Panicf("initPools ERROR! err:%v", err)
	}
	return nil
}

//协程池添加task
func SubmitTask(ctx context.Context, task func()) (err error) {
	if err = taskPool.Submit(task); err != nil {
		log.Println(ctx, "taskPool SubmitTask ERROR! err:%v", err)
		return err
	}
	return nil
}
