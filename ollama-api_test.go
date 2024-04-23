package ollama_api

import (
	"fmt"
	"github.com/fzkun/goutil/jsonutil"
	"github.com/fzkun/ollama-api/types"
	"testing"
)

func Init() *OllamaApi {
	o := NewOllamaApi(Config{
		Url:   "http://10.8.0.22:11434",
		Model: "qwen:32b",
	})
	return o
}

func TestOllamaApi_Generate(t *testing.T) {
	o := Init()
	prompt := "你好你是谁"
	data, err := o.Generate(prompt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(jsonutil.StructToJsonString(data))
}

func TestOllamaApi_GenerateSSE(t *testing.T) {
	o := Init()
	prompt := "你好你是谁"
	err := o.GenerateSSE(prompt, func(data types.OllamaGenerateResp) {
		fmt.Println(data.Response)
	})
	if err != nil {
		t.Fatal(err)
	}
}
