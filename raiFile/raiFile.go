package raiFile

import (
	"io/ioutil"
	"os"
)

//ReadFile , return file's value (string)
func ReadFile(fileName string) (string, error) { // CleanCode kurallarına uyalım dost
	icerik, err := ioutil.ReadFile(fileName)
	if !(err == nil) {
		return "", err
	}
	return string(icerik), nil
}

//WriteFile write string in file (0) *
func WriteFile(fileName string, text string) error {
	dosya, err := os.Create(fileName)
	defer dosya.Close()
	if !(err == nil) {
		return err
	}
	_, err = dosya.WriteString(text)
	if !(err == nil) {
		return err
	}
	return nil
}

//WriteFileLines write string in file but lines :D *
func WriteFileLines(fileName, text string) error {
	icerik, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModePerm) // dosya adı - string , flag - int , perm - os.FileMode | fileName , flag , perm
	defer icerik.Close()
	if !(err == nil) {
		return err
	}
	_, err = icerik.WriteString(text + "\n") // 
	if !(err == nil) {
		return err
	}
	return nil
}
