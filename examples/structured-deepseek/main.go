package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/revrost/go-openrouter"
)

func main() {
	ctx := context.Background()
	client := openrouter.NewClient[any](os.Getenv("OPENROUTER_API_KEY"))

	type Result struct {
		Location    string  `json:"location"`
		Temperature float64 `json:"temperature"`
		Condition   string  `json:"condition"`
	}
	result := Result{
		Location:    "London",
		Temperature: 20.0,
		Condition:   "Sunny",
	}
	jsonString, err := json.Marshal(result)
	if err != nil {
		log.Fatalf("GenerateSchemaForType error: %v", err)
	}

	request := openrouter.ChatCompletionRequest[any]{
		Model: openrouter.DeepseekV3,
		Messages: []openrouter.ChatCompletionMessage{
			{
				Role:    openrouter.ChatMessageRoleSystem,
				Content: openrouter.Content{Text: "EXAMPLE JSON OUTPUT: " + string(jsonString)},
			},
			{
				Role:    openrouter.ChatMessageRoleUser,
				Content: openrouter.Content{Text: "What's the weather like in London?"},
			},
		},
		ResponseFormat: &openrouter.ChatCompletionResponseFormat[any]{
			Type: openrouter.ChatCompletionResponseFormatTypeJSONObject,
		},
	}

	pj, _ := json.MarshalIndent(request, "", "\t")
	fmt.Printf("request :\n %s\n", string(pj))

	res, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		fmt.Println("error", err)
	} else {
		b, _ := json.MarshalIndent(res, "", "\t")
		fmt.Printf("response :\n %s", string(b))
	}
}
