package raiFile

import (
	"io/ioutil"
	"os"
)

//readFile dosyayı okur string döner
func readFile(fileName string) (string, error) { // CleanCode kurallarına uyalım dost
	icerik, err := ioutil.ReadFile(fileName)
	if !(err == nil) {
		return "", err
	}
	return string(icerik), nil
}

//writeFile Dosyaya yazar ama sıfırlayarak .
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

//writeFileLines Dosyaya yazıyor ama sıfırlamadan .
func writeFileLines(fileName, text string) error {
	icerik, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModePerm) // dosya adı - string , flag - int , perm - os.FileMode
	defer icerik.Close()
	if !(err == nil) {
		return err
	}
	_, err = icerik.WriteString(text + "\n") // Alt satıra da biz indirelim
	if !(err == nil) {
		return err
	}
	return nil
}
