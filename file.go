package tools

import "os"

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func Save(fileName string, b []byte) (int, error) {
	if PathExists(fileName) {
		_ = os.Remove(fileName)
	}
	f, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = f.Close()
	}()
	return f.Write(b)
}
