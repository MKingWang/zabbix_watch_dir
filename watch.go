package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fsnotify/fsnotify"
)

/*
将监控主函数
*/
func watchDir(filename string) {
	watch, _ := fsnotify.NewWatcher()
	defer watch.Close()

	err := watch.Add(filename)
	if err != nil {
		fmt.Errorf("监听目录失败：%v", err.Error())
	}

	//把指定目录下的所有目录都加入监控
	for _, k := range listDir(filename) {
		watch.Add(fmt.Sprintf("%s/%s", filename, k))
	}

	for {
		select {
		case event := <-watch.Events:
			{
				if event.Op&fsnotify.Create == fsnotify.Create {
					//发现创建的文件是目录，则加入监控
					if isDir(event.Name) {
						watch.Add(event.Name)
						continue
					}
				}

				pushQueue(event.Name)
				fmt.Println(event.Name, " create")
			}
		case err := <-watch.Errors:
			{
				fmt.Errorf("Err:%v", err)
				return
			}
		}
	}

}

/*
检查是否是目录
*/
func isDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		fmt.Errorf("%v", err.Error())
		return false
	}
	return stat.IsDir()
}

/*
递归列出所有目录
*/
func listDir(path string) []string {
	var dirlist []string
	fileList, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Errorf("list %s err:%v", path, err.Error())
		return nil
	}
	for _, fileName := range fileList {
		//	fmt.Println(fileName.Name())
		if fileName.IsDir() {
			dirlist = append(dirlist, fmt.Sprintf("%s", fileName.Name()))
			for _, k := range listDir(fmt.Sprintf("%s/%s", path, fileName.Name())) {
				dirlist = append(dirlist, fmt.Sprintf("%s/%s", fileName.Name(), k))
			}
		}
	}
	return dirlist
}