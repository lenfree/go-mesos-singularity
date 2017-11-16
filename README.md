go-singularity
--------------

A Mesos hubspot/Singularity package written in Go. Since I couldn't
manage to find one, hence, write a new one. One of the intention of
having this package is so I could write a Terraform provider to
interface with this.

## Usage

For package dependency management, we use dep:
```bash
$ go get -u github.com/golang/dep/cmd/dep
```

If new package is required, pls run below command
after go get. For more information about dep, please
visit this URL https://github.com/golang/dep.
```bash
$ dep ensure
```

```go
package main

import (
	"fmt"

	singularity "github.com/lenfree/go-singularity"
)

func main() {
	c := singularity.Config{
		Host: "singularity.net/singularity",
	}
	client := singularity.New(c)
	r, _ := client.GetRequests()
	for _, i := range r {
		body, _ := client.GetRequestByID(i.Request.ID)
		fmt.Println(body)
	}
}
```


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request