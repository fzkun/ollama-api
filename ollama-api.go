package ollama_api

import (
	"bufio"
	"fmt"
	"github.com/fzkun/goutil/jsonutil"
	"github.com/fzkun/ollama-api/types"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type OllamaApi struct {
	ctx Context
}

func NewOllamaApi(cfg Config) *OllamaApi {
	o := &OllamaApi{
		ctx: Context{
			cfg: cfg,
		},
	}
	return o
}

// Generate /api/generate
func (o *OllamaApi) Generate(prompt string) (data types.OllamaGenerateResp, err error) {
	var (
		httpResp *resty.Response
	)
	if httpResp, err = resty.New().R().SetBody(types.OllamaGenerateReq{
		Model:  o.ctx.cfg.Model,
		Stream: false,
		Prompt: prompt,
	}).Post(fmt.Sprintf(o.ctx.cfg.Url + "/api/generate")); err != nil {
		return
	}
	respJson := httpResp.String()
	if err = jsonutil.JsonStrToStruct(respJson, &data); err != nil {
		err = errors.New(fmt.Sprintf("解析json失败,err=%s,json=%s", err.Error(), respJson))
		return
	}
	return
}

// GenerateSSE /api/generate sse方式对接
func (o *OllamaApi) GenerateSSE(prompt string, callback func(data types.OllamaGenerateResp)) (err error) {
	var (
		httpResp *resty.Response
	)
	httpResp, err = resty.New().
		R().
		SetBody(types.OllamaGenerateReq{
			Model:  o.ctx.cfg.Model,
			Stream: true,
			Prompt: prompt,
		}).
		SetDoNotParseResponse(true).
		Post(fmt.Sprintf(o.ctx.cfg.Url + "/api/generate"))
	if err != nil {
		return err
	}
	defer httpResp.RawResponse.Body.Close()

	scanner := bufio.NewScanner(httpResp.RawResponse.Body)
	//reply := ""
	for scanner.Scan() {
		_res := scanner.Text()
		if _res == "" {
			continue
		}
		var data types.OllamaGenerateResp
		err = jsonutil.JsonStrToStruct(_res, &data)
		if err != nil {
			err = errors.Wrap(err, "解析json失败")
			return
		}
		callback(data)
		//reply += _res
		//fmt.Println(_res)
	}
	return err
}
