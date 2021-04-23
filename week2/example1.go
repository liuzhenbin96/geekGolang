package main

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

/**
What does handling an error mean? It means three things:
	1.The error has been logged.
	2.The application is back to 100% integrity.
	3.The current error is not reported any longer.
*/

func Download(url string) ([]byte, error) {
	// Download the tar file.
	r, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "get url: %s", url)
	}
	defer r.Body.Close()

	// Read in the entire contents of the file.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "writing file")
	}

	// Return the file.
	return body, nil
}
