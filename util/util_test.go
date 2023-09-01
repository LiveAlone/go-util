package util

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestMakeDir(t *testing.T) {
	err := CreateAllParentDirs("dest/a/c/active.java")
	if err != nil {
		log.Fatal(err)
	}
}

func TestExcelRead(t *testing.T) {
	excelFile := "demo.xlsx"
	sheetIndex := 0
	data, err := ReadExcelData(excelFile, sheetIndex)
	if err != nil {
		t.Error(err)
	}
	for index, row := range data {
		t.Logf("row %d: %v", index, row)
	}
}

func TestTask(t *testing.T) {
	action := func(ctx context.Context, name string, data interface{}) error {
		fmt.Println("task: ", name)
		for i := 0; i < 100; i++ {
			if name == "task-3" && i == 5 {
				return fmt.Errorf("task %v error", name)
			}

			time.Sleep(time.Second)
			fmt.Println(name, "finish ct", i)
			select {
			case <-ctx.Done():
				fmt.Println(name, "gain channel cancel")
				return nil
			default:
			}
		}
		return nil
	}

	tasks := make([]*Task, 0)
	for i := 0; i < 5; i++ {
		tasks = append(tasks, &Task{
			Name:   fmt.Sprintf("task-%d", i),
			Action: action,
			BizCtx: map[string]interface{}{},
		})
	}
	err := ConcurrentTaskExec(tasks)
	fmt.Println(err)
}
