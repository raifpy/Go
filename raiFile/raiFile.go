package raiFile

import (
	"io/ioutil"
	"os"
)

//readFile , return file's value (string)
func readFile(fileName string) (string, error) { // CleanCode kurallarına uyalım dost
	icerik, err := ioutil.ReadFile(fileName)
	if !(err == nil) {
		return "", err
	}
	return string(icerik), nil
}

//writeFile write string in file (0) *
func writeFile(fileName string, text string) error {
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

//writeFileLines write string in file but lines :D *
func writeFileLines(fileName, text string) error {
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
