package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	linkLib = "https://binaries.tonlabs.io/"

	nameFileMac = "tonclient_%s_darwin.gz"
	wayFileMac  = "../gateway/client/lib/darwin/amd64/libton_client.dylib"

	nameFileMacArm = "tonclient_%s_darwin_arm64.gz"
	wayFileMacArm  = "../gateway/client/lib/darwin/arm64/libton_client.dylib"

	nameFileLinux = "tonclient_%s_linux.gz"
	wayFileLinux  = "../gateway/client/lib/linux/libton_client.so"

	nameFileWinDll = "tonclient_%s_win32_dll.gz"
	wayFileWinDll  = "../gateway/client/lib/windows/libton_client.dll"
)

var version string
var dataLinks []string
var nameLibs []string

func init() {
	if len(os.Args) > 0 {
		version = os.Args[1]
		if version == "" || version == "0" {
			version = "1"
		}
	} else {
		version = "1"
	}

	version = strings.ReplaceAll(version, ".", "_")
	getLinks(nameFileMac)
	getLinks(nameFileMacArm)
	getLinks(nameFileLinux)
	getLinks(nameFileWinDll)
}

func main() {
	for idx, val := range dataLinks {
		var currentWay string
		switch nameLib := nameLibs[idx]; nameLib {
		case addVersion(nameFileMac):
			currentWay = wayFileMac
		case addVersion(nameFileMacArm):
			currentWay = wayFileMacArm
		case addVersion(nameFileLinux):
			currentWay = wayFileLinux
		case addVersion(nameFileWinDll):
			currentWay = wayFileWinDll
		default:
			log.Fatal("Not supported type")
		}

		log.Println(currentWay)
		if _, err := os.Stat(currentWay); !os.IsNotExist(err) {
			err = os.Remove(currentWay)
			if err != nil {
				log.Fatal(currentWay, err)
			}
		}

		ff, err := os.Create(currentWay)
		if err != nil {
			log.Fatal("Error for writing file: ", err)
		}

		response, err := http.Get(val)
		if err != nil {
			log.Fatal("Error for get file: ", err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			log.Fatal("Error for get file status response not 200: ", err, response.StatusCode)
		}

		reader, err := gzip.NewReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(ff, reader)
		if err != nil {
			log.Fatal(err)
		}
		defer ff.Close()

		time.Sleep(100 * time.Millisecond)
	}
}

func getLinks(fileName string) {
	nameLib := addVersion(fileName)
	nameLibs = append(nameLibs, nameLib)
	dataLinks = append(dataLinks, linkLib+nameLib)
}

func addVersion(name string) string {
	return fmt.Sprintf(name, version)
}
