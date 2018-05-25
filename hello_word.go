package main
import (
		"fmt" 
		"encoding/json"
		)


type people struct{
	Name string `json: "name"`
	Craft string `json: "craft"`
}

type peoples struct {
	People []people `json: "people"`
}

func main() {
	text := `{"people": [{"name": "Anton Shkaplerov", "craft": "ISS"}, {"name": "Scott Tingle", "craft": "ISS"}, {"name": "Norishige Kanai", "craft": "ISS"}, {"name": "Oleg Artemyev", "craft": "ISS"}, {"name": "Andrew Feustel", "craft": "ISS"}, {"name": "Richard Arnold", "craft": "ISS"}], "number": 6, "message": "success"}`
	textBytes := []byte(text)

	p1 := peoples{}
	err := json.Unmarshal(textBytes, &p1)

	if err != nil {
		fmt.Println(err)
		return
	}
	for _,e := range p1.People{
		fmt.Println(e.Name)
	}
	
}