/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by my-gin.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package util

import "os"

// 目录 or 文件 是否存在
func FileExist(file string) bool {
	_, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

// 递归创建目录 or 文件
func CreateFile(filePath string) error {
	if !FileExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}

	return nil
}
