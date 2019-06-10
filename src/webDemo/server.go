package webDemo

import (
	"net/http"
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
}
