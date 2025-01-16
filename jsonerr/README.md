## Description

JSONErr is a golang error handler to combine functionality between standard error output and error return from a function.

### Basic
Import jsonerr into your application, and create new instance

```sh
"github.com/devetek/go-core/jsonerr"
```

### Usage

```sh
package main

import (
    "os"
	"github.com/devetek/go-core/jsonerr"
)

func main() {
	// init logger instance
	Logger := jsonerr.New(
		jsonerr.WithCaller(true),
		jsonerr.WithAlwaysPrint(true),
	)
	
	var filename = "./my-file-does-not-exist"
	err := getFileContent(filename, Logger)
	if err != nil {
		Logger.Create(jsonerr.NewError{
			Message: "Failed to read file from main package",
			Error:   err,
		})

		Logger.Print()
		os.Exit(2)
	}
}

func getFileContent(file string, logger jsonerr.Error) error {
	// check file exist
	_, err := os.ReadFile(file)
	if err != nil {
		logger.Create(jsonerr.NewError{
			Message: "Failed to read file from child package",
			Error:   err,
		})

		return logger.Unwrap()
	}

	return nil
}
```