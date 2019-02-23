package config

import (
//	"../crypto"
//	"errors"
//	"log"

//	"strconv"

  "io/ioutil"
//  "path/filepath"
  "strings"
)


const up_path string = "/home/yuichi/github.com/Tsukiyama-do/qso/uploads"

type DummyFiles struct {
	files map[string]interface{}
}

var fls DummyFiles

func init() {
	fls.files = map[string]interface{}{}
}

func DummyF() *DummyFiles {
	return &fls
}


// The func returns slice of file name list
func dirwalk(dir string) ([]string, error) {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        return nil, err
    }

    var paths []string
    for _, file := range files {
  /*      if file.IsDir() {
            paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
            continue
        }
  */
        paths = append(paths, file.Name())
    }

    return paths, err
}


func (df *DummyFiles) UploadExists() ( string, error) {

    sl, err := dirwalk(up_path)
    if err != nil { return "", err }


    var sl_json string   // json format string

    sl_json = "["

    for _ , item := range sl {
      sl_json = sl_json + "{ \"filename\" : \"" + item + "\" },"
    }

    if len(sl_json) > 2  {
      sl_json = strings.TrimRight(sl_json, ",")
    }
    sl_json = sl_json + "]"

  	return sl_json, nil

}
