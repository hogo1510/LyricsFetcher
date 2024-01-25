package main
import(
	"fmt"
//	"encoding/json"
	"io/ioutil"
    	"net/http"
)

type data struct {
	//lyrics: string `json:'lyrics'`
}

func main() {

	url := "https://api.lyrics.ovh/v1/Laufey/From%20The%20Start"

	response, ew := http.Get(url)
	if ew != nil {
		panic(ew)
    }    
	body, ew := ioutil.ReadAll(response.Body)
	if ew != nil {
		panic(ew)
    }
	fmt.Println(string(body))	//raw json 

	//data_obj := data{}



}



