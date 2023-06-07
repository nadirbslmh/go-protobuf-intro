package main

import (
	"fmt"
	"go-protobuf-intro/bookpb/bookpb"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	// create book entity
	var book bookpb.Book = bookpb.Book{
		Id:          "123123",
		Title:       "learn protocol buffers",
		Description: "really good book",
		Author:      "ray krieger",
	}

	// convert to JSON
	result, err := convertToJSON(&book)
	if err != nil {
		panic(err)
	}

	// print the result
	fmt.Println("result: ", result)

	convertedBook := bookpb.Book{}

	// convert to Protocol Buffers
	err = converToProto(result, &convertedBook)
	if err != nil {
		panic(err)
	}

	// print the result
	fmt.Println("convert to proto result")
	fmt.Println(convertedBook.GetId(), convertedBook.GetTitle(),
		convertedBook.GetDescription(), convertedBook.GetAuthor())
}

func convertToJSON(pb proto.Message) (string, error) {
	result, err := protojson.Marshal(pb)

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func converToProto(data string, pb proto.Message) error {
	err := protojson.Unmarshal([]byte(data), pb)

	if err != nil {
		return err
	}

	return nil
}
