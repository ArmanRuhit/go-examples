package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct{
	Page int
	Fruits []string
}

type response2 struct{
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

var Println = fmt.Println

func main() {
	bolB, _ := json.Marshal(true)
	Println("bolB:", string(bolB))

	intB, _ := json.Marshal(1)
	Println("intB:",string(intB))

	fltB, _ := json.Marshal(2.34)
	Println("fltB:",string(fltB))

	strB, _ := json.Marshal("gopher")
	Println("strB:",string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	Println("slcD:",string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	Println("mapB:", string(mapB))

	resp1D := &response1{
		Page: 1,
		Fruits: []string{
			"apple",
			"peach",
			"pear",
		},
	}
	resp1B, _ := json.Marshal(resp1D)
	Println("resp1B:", string(resp1B))

	resp2D := &response2{
		Page: 1,
		Fruits: []string{
			"apple",
			"peach",
			"pear",
		},
	}
	resp2B, _ := json.Marshal(resp2D)
	Println("resp2B:",string(resp2B))

	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	Println("dat:",dat)

	num := dat["num"].(float64)
	Println("num",num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	Println("str1",str1)

	str := `{"page" : 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	Println("res.Fruits:",res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	Println("d:")
	enc.Encode(d)

}