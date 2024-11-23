package main

import (
	"encoding/json"
	"fmt"
)

// structs allows 3rd parameter
type course struct {
	Name     string 	`json:"coursename"`
	Price    int		`json:"courseprice"`
	Platform string		`json:"platform"`
	Password string		`json:"-"`   // The '-' here means the field will not be consumed means not sent in the final response 
	Tags     []string	`json:"tags,omitempty"`
}

func main() {
	EncodeJson()
	DecodeJson()

}

func EncodeJson(){
	courses := []course{
		{"Python masterclass", 100, "IKAD institutions", "abc123", []string{"Python", "Software"}},
		{"Golang masterclass", 500, "IKAD institutions", "bcd123", []string{"Golang", "Software"}},
		{"JS masterclass", 399, "IKAD institutions", "hij123", nil},
	}

	// result, err := json.Marshal(courses)
	result, err := json.MarshalIndent(courses,"","\t")
	if err!=nil{
		panic(err)
	}

	fmt.Printf("%s\n",result)
}

func DecodeJson(){
	// The recieved data from an API is in the form of bytes
	jsonData := []byte(`
	{
                "coursename": "Python masterclass",
                "courseprice": 100,
                "platform": "IKAD institutions",
                "tags": [
                        "Python",
                        "Software"
                ]
        }
	`)

	//checking if JSOn data is valid or not
	checkValid := json.Valid(jsonData)

	//Case 1: Unmarshling JSON into struct
	var Course course
	if (checkValid){
		fmt.Println("JSON is valid")
		err := json.Unmarshal(jsonData,&Course) //unmarshaling json data into a struct type
		if err!=nil{
			panic(err)
		}
		fmt.Printf("%#v\n",Course)
	}else{
		fmt.Println("JSON was not valid")
	}


	//Case 2: Coverting JSON data to map

	var data map[string]interface{}
	if (checkValid){
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData,&data)
		fmt.Printf("%#v\n",data)
	}else{
		fmt.Println("JSON was not valid")
	}
	

}