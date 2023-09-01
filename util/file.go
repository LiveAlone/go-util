package util

import (
	"bufio"
	"os"
	"path/filepath"
)

func ReadFileLines(filePath string) (fileLines []string, err error) {
	readFile, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return
}

func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func WriteFileLines(path string, lines []string) error {
	data := make([]byte, 0)
	for _, line := range lines {
		data = append(data, line...)
		data = append(data, '\n')
	}
	return os.WriteFile(path, data, 0644)
}

func WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

// CreateAllParentDirs 判断文件父级是否存在， 不存在创建
func CreateAllParentDirs(filePath string) error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := CreateAllParentDirs(dir); err != nil {
			return err
		}
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
