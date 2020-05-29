## Installation

```bash
$ go get github.com/FengWuTech/open-client
```

## Usage
```bash
import "github.com/FengWuTech/open-client"
```
Create a new client, then use the exposed services to access different parts of the OPenPlatform API.

You can use your AppID/AppSecret to create a new client:
```bash
package main

import (
    "github.com/FengWuTech/open-client"
)

func main() {
    client := openclient.NewClient(APP_ID, APP_SECRET_KEY, BASE_URL_DEV)
}
```