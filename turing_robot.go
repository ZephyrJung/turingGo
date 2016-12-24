package main
 
import (
    "net/http"
    "io/ioutil"
    "fmt"
    "bytes"
)

const(
    API_ADDRESS="http://www.tuling123.com/openapi/api"
    API_KEY=""
)

func main() {
    fmt.Println("URL:>", API_ADDRESS)
    var jsonStr = []byte("{\"key\":\""+API_KEY+"\",\"info\":\"你好\"}")
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
    fmt.Println("response Body:", string(body))
}