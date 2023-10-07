// 与中央后台交互

package models

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type CenterRsp struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data,omitempty"`
}

var urlPre = []string{
	"http://center-test.jmxy.kgogame.com/api/install/",
	"https://center.jmxy.kgogame.com/api/install/",
}

func doGet(urlIdx int, api string, params *url.Values) (*CenterRsp, error) {
	params.Set("t", fmt.Sprintf("%d", time.Now().Unix())) //当前时间戳
	paramstr := params.Encode()

	// sign
	key := md5.Sum([]byte(paramstr + "ee445b8afee814706ebefcbe1cd26ae8"))
	params.Set("sign", fmt.Sprintf("%x", key))
	paramstr = params.Encode()

	apiUrl := urlPre[urlIdx] + api
	urlObj, err := url.Parse(apiUrl)
	if err != nil {
		return nil, err
	}
	urlObj.RawQuery = paramstr

	res, err := http.Get(urlObj.String())
	if err != nil {
		return nil, err
	}

	rsp, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	var rspJson CenterRsp
	err = json.Unmarshal(rsp, &rspJson)
	if err != nil {
		return nil, err
	}
	return &rspJson, nil
}

// // 获取中央后台平台信息
// func CenterGetAgent(game *Game) error {
// 	agent := AgentGetByAid(game.Aid)

// 	params := url.Values{}
// 	params.Set("aid", fmt.Sprintf("%d", game.Aid))
// 	params.Set("name", agent.Flag)
// 	params.Set("timezone", "Asia/Shanghai")
// 	params.Set("domain", game.Domain)

// 	urlIdx := getCenterUrlIndx(game.Sid)
// 	rsp, err := doGet(urlIdx, "agent", &params)
// 	if err != nil {
// 		return err
// 	} else if rsp.Code != 1 {
// 		return fmt.Errorf("CenterGetAgent failed, err:%s", rsp.Msg)
// 	}

// 	return nil
// }

// // 注册新服，并获取新服 serial
// func CenterRegisterServer(game *Game) error {
// 	agent := AgentGetByAid(game.Aid)

// 	params := url.Values{}
// 	params.Set("aid", fmt.Sprintf("%d", game.Aid))          //平台id
// 	params.Set("agent", agent.Flag)                         //平台名称
// 	params.Set("server_id", fmt.Sprintf("%d", game.Sid))    //服id
// 	params.Set("server_flag", fmt.Sprintf("S%d", game.Sid)) //新服标示（eg.: S123）
// 	params.Set("domain", game.Domain)                       //域名
// 	params.Set("open_time", libs.FormatTime(game.OpenTime)) //开服时间
// 	params.Set("cdn", fmt.Sprintf("%s/%s", agent.Domain, agent.GetFlag()))
// 	params.Set("server_version", fmt.Sprintf("%d", game.Version))

// 	urlIdx := getCenterUrlIndx(game.Sid)
// 	rsp, err := doGet(urlIdx, "server", &params)
// 	if err != nil {
// 		return err
// 	} else if rsp.Code != 1 {
// 		return fmt.Errorf("CenterRegisterServer failed, err:%s", rsp.Msg)
// 	}
// 	return nil
// }

// // 更新游服
// func CenterUpdateServer(game *Game) error {
// 	agent := AgentGetByAid(game.Aid)

// 	params := url.Values{}
// 	params.Set("agent", agent.Flag)                      //平台名称
// 	params.Set("server_id", fmt.Sprintf("%d", game.Sid)) //服id
// 	params.Set("server_version", fmt.Sprintf("%d", game.Version))

// 	urlIdx := getCenterUrlIndx(game.Sid)
// 	rsp, err := doGet(urlIdx, "server", &params)
// 	if err != nil {
// 		return err
// 	} else if rsp.Code != 1 {
// 		return fmt.Errorf("CenterUpdateServer failed, err:%s", rsp.Msg)
// 	}
// 	return nil
// }
