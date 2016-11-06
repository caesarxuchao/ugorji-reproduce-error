package main

import (
	"encoding/json"
	"fmt"

	"github.com/ugorji/go/codec"
)

func main() {
	exampleList := ExampleList{
		Items: []Example{
			Example{
				Spec: ExampleSpec{
					Foo: "1",
					Bar: true,
				},
			},
		},
	}

	data, err := json.Marshal(exampleList)
	if err != nil {
		panic(fmt.Sprintf("json marshal failed, err=%s", err))
	}
	fmt.Printf("marshaled data=%s\n", data)

	into := &ExampleList{}

	if err := codec.NewDecoderBytes(data, new(codec.JsonHandle)).Decode(into); err != nil {
		// I expect it to not panic, but it panic with: [pos 1]: json: expect char '"' but got char '{'
		// The unmarshal works if I define UnmarshalJSON for Example and ExampleList, as in https://github.com/caesarxuchao/k8s-tpr-playground/blob/workaround/main.go#L66-L96.
		panic(err.Error())
	}
	fmt.Printf("Sucessfully decoded, into=%#v\n", into)

}

type ExampleSpec struct {
	Foo string `json:"foo"`
	Bar bool   `json:"bar"`
}

type Example struct {
	Spec ExampleSpec `json:"spec"`
}

type ExampleList struct {
	Items []Example `json:"items"`
}

func (e *Example) UnmarshalText(data []byte) error {
	return json.Unmarshal(data, e)
}

func (el *ExampleList) UnmarshalText(data []byte) error {
	return json.Unmarshal(data, el)
}
