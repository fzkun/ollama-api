# ollama-api

## 介绍

接入ollama的api工具

## 例子

```go
o := NewOllamaApi(Config{
		Url:   "http://localhost:11434",
		Model: "qwen:32b",
	})

prompt := "你好你是谁"

data, err := o.Generate(prompt)

if err != nil {
    t.Fatal(err)
}

fmt.Println(jsonutil.StructToJsonString(data))

//{"model":"qwen:32b","created_at":"2024-04-23T13:10:36.132216626Z","response":"我是通义千问，由阿里云开发的AI助手。我可以回答各种问题、提供信息和与用户进行对话。有什么我可以帮助你的吗？","done":true,"total_duration":3141457095,"load_duration":716703,"prompt_eval_duration":99576000,"eval_count":34,"eval_duration":3039901000}
```

## SSE方式请求
```go
o := NewOllamaApi(Config{
Url:   "http://localhost:11434",
Model: "qwen:32b",
})

prompt := "你好你是谁"
err := o.GenerateSSE(prompt, func (data types.OllamaGenerateResp) {
fmt.Println(data.Response)
})
if err != nil {
t.Fatal(err)
}

```

