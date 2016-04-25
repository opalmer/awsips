# AWSIPs
Simple package for golang, and possibly other languages in the future, which
downloads and returns information from
https://ip-ranges.amazonaws.com/ip-ranges.json.


# Usage
## go

```go
package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/opalmer/awsips"
)


func main() {
	ranges, err := awsips.Get()
	if err != nil {
		log.Fatal(err)
	}

	for _, prefix := range ranges.Prefixes {
		// logic
	}
}
```
