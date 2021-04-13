package crontab

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/toolbox"
	"time"
)

func TestTask1() error {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func init() {
	tk1 := toolbox.NewTask("TestTask1", "*/5 * * * * *", TestTask1)
	toolbox.AddTask("TestTask1", tk1)

}
