package seeds

import (
	"bytes"
	"encoding/json"
	"github.com/rs/xid"
	"hero/pkg/logger"
	"hero/utils"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type PlayRequest struct {
	FbUserID    string `json:"fb_user_id" form:"fb_user_id" query:"fb_user_id"`
	FBAvatarUrl string `json:"fb_avatar_url" form:"fb_avatar_url" query:"fb_avatar_url"`
	FbEmail     string `json:"fb_email" form:"fb_email" query:"fb_email"`
	FbName      string `json:"fb_name" form:"fb_name" query:"fb_name"`
}

const TIME_LAYOUT = "2006-01-02 15:04:05"

func ActiveSeed() {
	var wg sync.WaitGroup
	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		xid := xid.New()
		go activeSeed(xid.String(), wg)
	}
	wg.Wait()
}

func activeSeed(fbID string, wg sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://hero-lxaqvhvivq-an.a.run.app/api/active/hero/play", nil)
	//req, err := http.NewRequest("POST", "http://127.0.0.1:9000/api/active/hero/play", nil)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		return
	}
	playRequest := &PlayRequest{
		FbUserID:    fbID,
		FBAvatarUrl: "http://localhost.hero.com/123",
		FbEmail:     "frankie.lee.job@gmail.com",
		FbName:      "frankie",
	}
	jsonStr, err := json.Marshal(playRequest)
	if err != nil {
		log.Print("marshal err: ", err.Error())
		return
	}
	b := bytes.NewBuffer(jsonStr)
	req.Body = ioutil.NopCloser(b)
	md5Key := utils.GenMd5Key(fbID)
	req.Header.Add("Authorization", "Bearer "+md5Key)
	apiResp, err := client.Do(req)
	if err != nil {
		logger.Print("test err", err.Error())
		return
	}
	log.Print("seed play record:", apiResp.Status)
	if apiResp.Body != nil {
		apiResp.Body.Close()
		return
	}
}
