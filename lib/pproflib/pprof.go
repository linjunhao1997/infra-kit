package pproflib

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/pprof"
	"time"
)

const (
	cpuProfile  = "_cpuprofile"
	heapProfile = "_heapprofile"
)

func StartProfile(dir string) error {
	ticker := time.NewTicker(1 * time.Minute)
	// 生产场景建议把文件输出到s3等文件系统
	for {
		if err := checkAndDeleteExpiredFiles(dir); err != nil {
			return err
		}
		cpuFileName := dir + "/" + time.Now().Format("2006-01-02_15-04-05") + cpuProfile
		heapProfile := dir + "/" + time.Now().Format("2006-01-02_15-04-05") + heapProfile
		cpuf, err := openFile(cpuFileName)
		if err != nil {
			return err
		}
		defer cpuf.Close()
		heapf, err := openFile(heapProfile)
		if err != nil {
			return err
		}
		defer heapf.Close()
		pprof.StartCPUProfile(cpuf)
		pprof.WriteHeapProfile(heapf)

		<-ticker.C
		pprof.StopCPUProfile()

	}
}

func openFile(fileName string) (*os.File, error) {
	// 获取文件所在的目录
	dir := filepath.Dir(fileName)

	// 如果目录不存在，则创建目录
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func checkAndDeleteExpiredFiles(dir string) error {
	files, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		err = fmt.Errorf("Error reading directory, err=%v", err)
		return err
	}

	// 当前时间
	now := time.Now()

	// 遍历文件并检查删除过期文件
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// 使用 os.Stat 获取文件信息
		fileInfo, err := os.Stat(dir + "/" + file.Name())
		if err != nil {
			err = fmt.Errorf("Error getting file info, fiiename=%v, err=%v", file.Name(), err)
			return err
		}

		// 获取文件创建时间
		createTime := fileInfo.ModTime()

		// 如果文件创建时间超过一小时，则删除文件
		if now.Sub(createTime) > 1*time.Hour {
			err := os.Remove(dir + "/" + file.Name())
			if err != nil {
				err = fmt.Errorf("Error deleting file, fiiename=%v, err=%v", file.Name(), err)
				return err
			} else {
				fmt.Println("File deleted:", file.Name())
			}
		}
	}
	return nil
}
