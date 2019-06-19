## bson-objectId

BSON ObjectIds are small, likely unique, fast to generate, and ordered.

Uses:

```go
package main

import (
	"fmt"
	bsonid "github.com/ankur-anand/bson-objectId"
)

func main() {
	uuid := bsonid.New()
	fmt.Println(uuid) // 5d0a83d691a2fa500f000001
}
```