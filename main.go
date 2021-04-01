package main

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	openapi_v2 "github.com/googleapis/gnostic/openapiv2"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	jb, err := ioutil.ReadFile("./openapi.json")
	if err != nil {
		return err
	}
	fmt.Printf("length of bytes for input json openapi doc: %v\n", len(jb))

	// https://pkg.go.dev/github.com/golang/protobuf/jsonpb#Unmarshaler.Unmarshal may be any alternative here for decoding json.
	doc, err := openapi_v2.ParseDocument(jb)
	if err != nil {
		return err
	}

	pb, err := proto.Marshal(doc)
	if err != nil {
		return err
	}
	fmt.Printf("length of bytes for the proto serialized openapi doc: %v\n", len(pb))

	var newDoc openapi_v2.Document
	if err = proto.Unmarshal(pb, &newDoc); err != nil {
		return err
	}

	swg := newDoc.String()
	fmt.Printf("length of string for the json serialized openapi doc: %v\n", len(swg))
	fmt.Printf("Printing part of the serialized json openapi doc:\n %s\n", swg[:200])

	return nil
}
