package go_concurrency

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//func DoRdpCommand(clis []CliSpec, pipe chan CliSpec) {
//	for k, _ := range clis {
//		// 单个超时
//		ticker := time.NewTicker(time.Duration(1) * time.Second) // 主要考虑到批量的问题
//		go RdpConn(clis[k], pipe)
//		select {
//		case <-ticker.C:
//			pipe <- CliSpec{
//				MakeConfig: clis[k].MakeConfig,
//				AssetName:  clis[k].AssetName,
//				Alive:      false,
//			}
//		}
//	}
//}

func cb(args ...interface{}) error {
	fmt.Println("Start")
	time.Sleep(5 * time.Second)
	return nil
}

func TestRun(t *testing.T) {
	err := Run(cb, 3*time.Second, 3, "55")
	fmt.Println(err)
	fmt.Println("number of goroutines:", runtime.NumGoroutine())

}
