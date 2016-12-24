package main

import (
    "fmt"
    "log"
    "net/http"
    "bytes"
    "io/ioutil"
    "github.com/sidbusy/weixinmp"
    "encoding/json"
)

const(
    API_ADDRESS="http://www.tuling123.com/openapi/api"
    API_KEY=""
)

func main() {
    // 注册处理函数
    fmt.Println("server start...")
    http.HandleFunc("/receiver", receiver)
    log.Fatal(http.ListenAndServe(":80", nil))
}

func receiver(w http.ResponseWriter, r *http.Request) {
    fmt.Println("get request from weixin")
    token := "" // 微信公众平台的Token
    appid := "" // 微信公众平台的AppID
    secret := "" // 微信公众平台的AppSecret
    // 仅被动响应消息时可不填写appid、secret
    // 仅主动发送消息时可不填写token
    mp := weixinmp.New(token, appid, secret)
    // 检查请求是否有效
    // 仅主动发送消息时不用检查
    if !mp.Request.IsValid(w, r) {
        return
    }
    // 判断消息类型
    if mp.Request.MsgType == weixinmp.MsgTypeText {
        // 回复消息
        fmt.Println(mp.Request.Content)
        respStr := turing(mp.Request.Content)
        mp.ReplyTextMsg(w, respStr)
    }
}

func turing(content string) string {
    // fmt.Println("URL:>", API_ADDRESS)
    var jsonStr = []byte("{\"key\":\""+API_KEY+"\",\"info\":\""+content+"\"}")
    req, err := http.NewRequest("POST", API_ADDRESS, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    var jrsp map[string]interface{}
    if err := json.Unmarshal(body,&jrsp); err!=nil{
        return "404 Not Found"
    }
    fmt.Println("蛋蛋：",jrsp["text"].(string))
    return jrsp["text"].(string)
}