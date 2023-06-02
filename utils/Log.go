package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// GetRootLogFile
// 生成Root日志文件
func GetRootLogFile() (f *os.File, err error) {

	_, err = os.Stat("./storage/log")
	if err != nil {
		os.MkdirAll("./storage/log", 0777)
	}
	filename := "root_" + time.Now().Format("2006-01-02") + ".log"
	filepath := "./storage/log/" + filename
	_, err = os.Stat(filepath)
	if err != nil {
		f, err = os.Create(filepath)
		if err != nil {
			log.Println("[os.Create ERROR]: ", err)
			return
		}
	} else {
		f, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Println("[os.OpenFile ERROR]: ", err)
			return
		}
	}
	return
}

// GetDebugFile
// 生成debug文件
func GetDebugFile() (f *os.File, err error) {
	_, err = os.Stat("./storage/log")
	if err != nil {
		os.MkdirAll("./storage/log", 0777)
	}
	filename := "debug_" + time.Now().Format("2006-01-02") + ".log"
	filepath := "./storage/log/" + filename
	_, err = os.Stat(filepath)
	if err != nil {
		f, err = os.Create(filepath)
		if err != nil {
			log.Println("[os.Create ERROR]: ", err)
			return
		}
	} else {
		f, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Println("[os.OpenFile ERROR]: ", err)
			return
		}
	}
	return
}

// GetErrorFile
// 生成Error文件
func GetErrorFile() (f *os.File, err error) {
	_, err = os.Stat("./storage/log")
	if err != nil {
		os.MkdirAll("./storage/log", 0777)
	}
	filename := "error_" + time.Now().Format("2006-01-02") + ".log"
	filepath := "./storage/log/" + filename
	_, err = os.Stat(filepath)
	if err != nil {
		f, err = os.Create(filepath)
		if err != nil {
			log.Println("[os.Create ERROR]: ", err)
			return
		}
	} else {
		f, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Println("[os.OpenFile ERROR]: ", err)
			return
		}
	}
	return
}

// GetInfoFile
// 生成Info文件
func GetInfoFile() (f *os.File, err error) {
	_, err = os.Stat("./storage/log")
	if err != nil {
		os.MkdirAll("./storage/log", 0777)
	}
	filename := "info_" + time.Now().Format("2006-01-02") + ".log"
	filepath := "./storage/log/" + filename
	_, err = os.Stat(filepath)
	if err != nil {
		f, err = os.Create(filepath)
		if err != nil {
			log.Println("[os.Create ERROR]: ", err)
			return
		}
	} else {
		f, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Println("[os.OpenFile ERROR]: ", err)
			return
		}
	}
	return
}

// WriteRootLog
// 写入Root日志
func WriteRootLog(data interface{}) error {
	f, err := GetRootLogFile()
	if err != nil {
		return err
	}
	defer f.Close()
	template := fmt.Sprintf("[ROOT-LOG] %s : %v \r\n", time.Now().Format("2006-01-02 15:04:05"), data)
	_, err = f.WriteString(template)
	if err != nil {
		return err
	}
	return nil
}

// WriteDebugLog
// 写入Debug日志
func WriteDebugLog(data interface{}) error {
	f, err := GetDebugFile()
	if err != nil {
		return err
	}
	defer f.Close()
	template := fmt.Sprintf("[DEBUG-LOG] %s : %v \r\n", time.Now().Format("2006-01-02 15:04:05"), data)
	_, err = f.WriteString(template)
	if err != nil {
		return err
	}
	return nil
}

// WriteErrorLog
// 写入Error日志
func WriteErrorLog(data interface{}) error {
	f, err := GetErrorFile()
	if err != nil {
		return err
	}
	defer f.Close()
	template := fmt.Sprintf("[ERROR-LOG] %s : %v \r\n", time.Now().Format("2006-01-02 15:04:05"), data)
	_, err = f.WriteString(template)
	if err != nil {
		return err
	}
	return nil
}

// WriteInfoLog
// 写入Info日志
func WriteInfoLog(data interface{}) error {
	f, err := GetInfoFile()
	if err != nil {
		return err
	}
	defer f.Close()
	template := fmt.Sprintf("[INFO-LOG] %s : %v \r\n", time.Now().Format("2006-01-02 15:04:05"), data)
	_, err = f.WriteString(template)
	if err != nil {
		return err
	}
	return nil
}
