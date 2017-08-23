package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Jsondata struct {
	Text       string `json:"text"`
	Username   string `json:"username"`
	Icon_emoji string `json:"icon_emoji"`
	Link_names int    `json:"link_names"`
	Channel    string `json:"channel"`
}

// JsonDataのダンプ
func Dump(j Jsondata) {
	//str := `{"text":"testmessage","username":"usermessage","icon_emoji":":icon:","link_names":1,"channel":"@hilolih"}`
	//json.Unmarshal([]byte(str), &j)
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	log.Println(string(b))
}

// ファイルからJsonを取り出す
func Load(file string) Jsondata {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c Jsondata
	json.Unmarshal(raw, &c)
	return c
}

func Send(path string, j Jsondata) {
	b, _ := json.Marshal(j)

	req, _ := http.NewRequest(
		"POST",
		path,
		bytes.NewBuffer(b),
	)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
}
