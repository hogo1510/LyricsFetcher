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

	//	Doordat je een spatie maakt in de naam; denkt tie dat het voor bijde scans geld

	var artiest string = "pixies"
	var nummer string = "monkey gone to heaven"
	var url, Artiest, Nummer string 

	Artiest = strings.Replace(artiest," ", "%20", -1)
	fmt.Println("replacement: " , Artiest)

	Nummer = strings.Replace(nummer," ", "%20", -1)
	fmt.Println("replacement: " , Nummer)
	url = "https://api.lyrics.ovh/v1/"+Artiest+"/"+Nummer
	fmt.Println(url)


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



