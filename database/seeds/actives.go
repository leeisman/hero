package seeds

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
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
	fbUserIDPrefix := 9999999999
	playRequest := &PlayRequest{
		FbUserID:    "",
		FBAvatarUrl: "http://localhost.hero.com/123",
		FbEmail:     "frankie.lee.job@gmail.com",
		FbName:      "frankie",
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:9000/api/active/hero/play", nil)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		return
	}

	for i := 0; i <= 3000; i++ {
		jsonStr, err := json.Marshal(playRequest)
		if err != nil {
			log.Print("marshal err: ", err.Error())
		}
		b := bytes.NewBuffer(jsonStr)
		req.Body = ioutil.NopCloser(b)
		fbID := fbUserIDPrefix + i
		fbIDStr := strconv.Itoa(fbID)
		playRequest.FbUserID = fbIDStr
		h := md5.New()
		h.Write([]byte(playRequest.FbUserID))
		md5Key := hex.EncodeToString(h.Sum(nil))
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", md5Key))
		apiResp, _ := client.Do(req)
		log.Print("seed play record:", apiResp)
		if apiResp.Body != nil {
			apiResp.Body.Close()
		}
	}
}
