package common

import (
	"bufio"
	"html/template"
	"io/ioutil"
	"os"
)

//LoadFile - load a file
func LoadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	records := make([]string, 0)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		records = append(records, scanner.Text())
	}
	return records, scanner.Err()
}

//LoadTemplate - load trmplate
func LoadTemplate(file string) *template.Template {
	fs, _ := os.Open(file)
	defer fs.Close()
	data, _ := ioutil.ReadAll(fs)
	t := template.Must(template.New("").Parse(string(data)))
	return t
}
