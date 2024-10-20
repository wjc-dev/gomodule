package vegetable

import "gomod/logprint"

func Make(s string) {
	name := "菜名: " + s
	logprint.Info(name)
}
