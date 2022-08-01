# Make Hello World

### 1. Initializing project

```bash
go mod init _project_name_
```

### 2. Install gorilla/mux

```bash
go get -u github.com/gorilla/mux
```

Package `gorilla/mux` implements a request router and dispatcher for matching incoming requests to their respective handler.

### 3. Create `main.go` file and write this below code to print 'Hello wolrd'

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	// On Terminal/Command Propt
  fmt.Println("Hello World!")

	// On http (API)
	r := mux.NewRouter()

	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
```
