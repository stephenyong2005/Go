package urlDemo

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func FetchURL(fn ...string){

	for _, u := range fn{
		resp, err := http.Get(u)
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v \n",u, err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v \n",u, err)
			os.Exit(1)
		}

		fmt.Printf("%s",b)
	}

}
