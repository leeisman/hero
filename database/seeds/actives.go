package seeds

import (
	"bytes"
	"encoding/json"
	"hero/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type PlayRequest struct {
	FbUserID    string `json:"fb_user_id" form:"fb_user_id" query:"fb_user_id"`
	FBAvatarUrl string `json:"fb_avatar_url" form:"fb_avatar_url" query:"fb_avatar_url"`
	FbEmail     string `json:"fb_email" form:"fb_email" query:"fb_email"`
	FbName      string `json:"fb_name" form:"fb_name" query:"fb_name"`
}

func ActiveSeed() {
	fbUserIDPrefix := 999999999

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://hero-lxaqvhvivq-an.a.run.app/api/active/hero/play", nil)
	//req, err := http.NewRequest("POST", "http://127.0.0.1:9000/api/active/hero/play", nil)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		return
	}

	for i := 0; i <= 1; i++ {
		fbID := fbUserIDPrefix + i
		fbUserID := strconv.Itoa(fbID)
		playRequest := &PlayRequest{
			FbUserID:    fbUserID,
			FBAvatarUrl: "http://localhost.hero.com/123",
			FbEmail:     "frankie.lee.job@gmail.com",
			FbName:      "frankie",
		}
		jsonStr, err := json.Marshal(playRequest)
		if err != nil {
			log.Print("marshal err: ", err.Error())
		}
		b := bytes.NewBuffer(jsonStr)
		req.Body = ioutil.NopCloser(b)
		fbIDStr := strconv.Itoa(fbID)
		playRequest.FbUserID = fbIDStr
		md5Key := utils.GenMd5Key(fbIDStr)
		req.Header.Add("Authorization", "Bearer "+md5Key)
		apiResp, _ := client.Do(req)
		log.Print("seed play record:", apiResp.Status)
		if apiResp.Body != nil {
			apiResp.Body.Close()
		}
	}
}
