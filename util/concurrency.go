package util

import (
	"context"
	"log"
	"sync"
)

// Task 并发执行任务结构体
type Task struct {
	Name   string                                                         `json:"name"` // 任务名称
	Action func(ctx context.Context, name string, data interface{}) error // 任务执行器, ctx 取消返回nil
	BizCtx interface{}                                                    // 执行任务参数结果
}

func ConcurrentTaskExec(tasks []*Task) error {
	var wg sync.WaitGroup
	ctx, ctxCancel := context.WithCancelCause(context.Background())
	for index, task := range tasks {
		wg.Add(1)
		go func(index int, task *Task) {
			defer wg.Done()
			err := task.Action(ctx, task.Name, task.BizCtx)
			if err != nil {
				log.Printf("task %v exec error: %v", task.Name, err)
				ctxCancel(err)
			}
		}(index, task)
	}
	wg.Wait()
	return ctx.Err()
}
