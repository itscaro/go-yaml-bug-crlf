package main

import (
	"fmt"
	// "gopkg.in/yaml.v2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func main()  {
	var out1, out2 interface{}
	var in1, in2 []byte
	var err error
	in1, _ = ioutil.ReadFile("lr.yaml")
	fmt.Printf("%s\n", string(in1))
	err = yaml.Unmarshal(in1, &out1)
	fmt.Printf("%s\n", out1)
	fmt.Printf("%s\n", err)

	fmt.Println("-----")

	in2, _ = ioutil.ReadFile("crlf.yaml")
	fmt.Printf("%s\n", string(in2))
	err = yaml.Unmarshal(in2, &out2)
	fmt.Printf("%s\n", out2)
	fmt.Printf("%s\n", err)
}
