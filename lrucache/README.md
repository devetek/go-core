## Description

Cache data to memory, it will helpful when communicate with network such as REST API, graphql, etc.

### Usage

```sh
package main

import (
    "fmt"
	"github.com/devetek/go-core/lrucache"
)

type DataStruct struct {
    ...............
    ...............
    ...............
}

func getData() DataStruct {
    var data DataStruct
    cache := lrucache.New(100, time.Second*60)

    // get data from cache before fetch
    dc := cache.Get("data")

    if dc != nil {
		data = dc.(DataStruct)
	}

    if dc == nil {
        // fetch data to origin url and validate data error
        ............
        // store origin response to data
        ............

        // then set data to cache if origin return success
        cache.Set("data", data)
    }

    return data
}

func main() {
    data := getData()

    fmt.Println(data)
}

```
