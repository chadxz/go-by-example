package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page,omitempty"`
	Fruits []string `json:"fruits,omitempty"`
}

func main() {
	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(2.34)
	fmt.Println(string(floatB))

	stringB, _ := json.Marshal("gopher")
	fmt.Println(string(stringB))

	sliceB, _ := json.Marshal([]string{"apple", "peach"})
	fmt.Println(string(sliceB))

	mapB, _ := json.Marshal(map[string]int{"apple": 1, "peach": 2})
	fmt.Println(string(mapB))

	response1B, _ := json.Marshal(&response1{
		Page:   2,
		Fruits: []string{"apple", "banana"},
	})
	fmt.Println(string(response1B))

	response2B, _ := json.Marshal(&response2{
		Page:   3,
		Fruits: []string{"pineapple", "kiwi"},
	})
	fmt.Println(string(response2B))

	jsonBytes := []byte(`{"num": 6.13, "strs": ["a", "b"]}`)
	var jsonData map[string]interface{}

	if err := json.Unmarshal(jsonBytes, &jsonData); err != nil {
		panic(err)
	}
	fmt.Println(jsonData)

	num := jsonData["num"].(float64)
	fmt.Printf("%T %f\n", num, num)

	strs := jsonData["strs"].([]interface{})
	fmt.Printf("%T %s\n", strs, strs)

	str := strs[0].(string)
	fmt.Println(str)

	jsonData2 := []byte(`{"page": 1, "fruits": ["apple", "peach"]}`)
	response := response2{}
	_ = json.Unmarshal(jsonData2, &response)
	fmt.Printf("%+v\n", response)
	fmt.Println(response.Fruits[0])

	encoder := json.NewEncoder(os.Stdout)
	myData := map[string]int{"apple": 7, "lettuce": 8}
	_ = encoder.Encode(myData)
}
