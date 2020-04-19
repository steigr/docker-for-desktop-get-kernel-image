package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/hooklift/iso9660"
	"gopkg.in/yaml.v2"
)

type Linuxkit struct {
	Kernel struct {
		Image string
	}
}

func main() {
	var (
		isoFilePath = os.Getenv("ISO_FILE_PATH")
		linuxkitFilePath = os.Getenv("LINUXKIT_FILE_PATH")
	)

	if isoFilePath == "" {
		isoFilePath = "/dev/sr0"
	}

	if linuxkitFilePath == "" {
		linuxkitFilePath = "/etc/linuxkit.yml"
	}


	file, err := os.Open(isoFilePath)
	if err != nil {
		panic(err)
	}

	r, err := iso9660.NewReader(file)
	if err != nil {
		panic(err)
	}

    var linuxkit Linuxkit

	for {
		f, err := r.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

        if f.Name() != linuxkitFilePath {
            continue
        }

		fileReader := f.Sys().(io.Reader)
		buf := new(bytes.Buffer)
        buf.ReadFrom(fileReader)
		err = yaml.Unmarshal(buf.Bytes(), &linuxkit)
		if err != nil {
			panic(err)
		}
		break
	}
	fmt.Println(linuxkit.Kernel.Image)
}
