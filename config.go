package ollama_api

type Config struct {
	Url   string `json:"url"`   //ollama地址
	Model string `json:"model"` //模型
}
