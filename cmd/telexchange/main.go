package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Human interface {
	increseAge() int
}
type Person struct {
	name string
	age  int
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	val, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(val)
	return nil
}

func (p Person) increseAge() int {
	return p.age + 10
}

func normalfunc(ival int) {
	ival = 8
}

func ptrfunc(ival *int) {
	fmt.Println("I value is: ", ival)
	*ival = 8
}

func main() {
	fmt.Println("Hi")

	whatAmi := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Println("I am unknown")
		}
	}
	whatAmi(0)
	r := make([]bool, 3)
	r = append(r, true)
	fmt.Println(r)

	// Initiate a new map
	m := make(map[string]int)
	m["Moha"] = 2
	_, val := m["Moha"]
	fmt.Println(val)

	nums := []int{1, 2, 3}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("total is: ", sum)

	i := 1
	normalfunc(i)
	fmt.Println("I is: ", i)
	ptrfunc(&i)
	fmt.Println("I is: ", i)

	p := []Human{&Person{name: "Moha", age: 44}}
	for _, val := range p {
		fmt.Println(val.increseAge())
	}

	var vale map[string]Timestamp

	input := `
    {
        "created_at": "Thu May 31 00:00:01 +0000 2012"
    }`

	// Json encoding example

	if err := json.Unmarshal([]byte(input), &vale); err != nil {
		panic(err)
	}
	fmt.Println(vale)

	for key, value := range vale {
		fmt.Println(key, reflect.TypeOf(value))
	}
}
