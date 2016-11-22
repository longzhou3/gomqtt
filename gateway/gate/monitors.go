package gate

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/uber-go/zap"
)

func monitorLeaking() {
	for {
		r := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -n |awk '{print $2}'|sort|uniq -c | grep %v", os.Getpid()))
		v, _ := r.Output()
		fds := strings.Split(string(v), " ")[2]
		Logger.Debug("goroutine和fd数目", zap.Int("goroutine", runtime.NumGoroutine()), zap.String("fd", fds))
		Logger.Debug("主要数据结构", zap.Object("conns", cons), zap.Object("mutexs", accounts))

		time.Sleep(20 * time.Second)

	}

}

func monitorsStart() {
	// monitor the goroutine and file descriptor leaking
	go monitorLeaking()
}
