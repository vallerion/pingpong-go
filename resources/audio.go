package resources

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type resourceManager struct {
	list map[string]*os.File
}

func (r *resourceManager) Get(resource string) *os.File {
	return r.list[resource]
}

var Resources resourceManager

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	audioPath := pwd + "/assets/audio/"

	files, err := ioutil.ReadDir(audioPath)
	if err != nil {
		panic(err)
	}

	Resources = resourceManager{make(map[string]*os.File, len(files))}

	for _, file := range files {
		f, err := os.Open(audioPath + file.Name())
		if err != nil {
			panic(err)
		}

		Resources.list[fileNameWithoutExtension(file.Name())] = f
	}
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
