// 包名最好用小写,不要带下户线
package logprint

import (
	"fmt"
	"time"
)

func Info(msg interface{}) {
	now := time.Now()
	fmt.Printf("[info] , 当前时间: %s , 内容: %s \n", now.Format("2006-01-02 15:04:05"), msg)
}
