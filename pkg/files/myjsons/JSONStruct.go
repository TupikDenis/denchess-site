package myjsons

import (
	"io"
	"net/http"
	"os"
)

type Data struct {
	URL      string
	filename string
}

func DownloadJSON(number string) string {
	data := &Data{
		URL:      "https://lichess.org/api/tournament/" + number + "/results",
		filename: "pkg/files/myjsons/" + number + ".json",
	}

	data.download()

	return data.filename
}

func (data *Data) download() error {
	response, err := http.Get(data.URL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	file, err := os.Create(data.filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		panic(err)
	}

	return nil
}
