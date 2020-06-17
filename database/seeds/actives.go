package seeds

import (
	"bytes"
	"encoding/json"
	"github.com/rs/xid"
	"hero/controllers/api/active/hero"
	"hero/database/ent"
	"hero/pkg/logger"
	"hero/utils"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type PlayRequest struct {
	FbUserID    string `json:"fb_user_id" form:"fb_user_id" query:"fb_user_id"`
	FBAvatarUrl string `json:"fb_avatar_url" form:"fb_avatar_url" query:"fb_avatar_url"`
	FbEmail     string `json:"fb_email" form:"fb_email" query:"fb_email"`
	FbName      string `json:"fb_name" form:"fb_name" query:"fb_name"`
}

var (
	success int32
	fail    int32
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

const HOST = "https://hero-lxaqvhvivq-an.a.run.app"

//const HOST = "http://127.0.0.1:9000"

func ActiveSeed() {
	start := time.Now()
	logger.Print("seed to ", HOST)
	var wg sync.WaitGroup
	for i := 0; i <= 1600; i++ {
		wg.Add(1)
		xid := xid.New()
		go activeSeed(xid.String(), &wg)
		time.Sleep(time.Second / 30)
	}
	wg.Wait()
	log.Print("fail: ", fail)
	log.Print("success: ", success)
	log.Print("start: ", start)
	log.Print("end: ", time.Now())
}

func activeSeed(fbID string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}
	req, err := http.NewRequest("POST", HOST+"/api/active/hero/play", nil)
	//req, err := http.NewRequest("POST", "http://127.0.0.1:9000/api/active/hero/play", nil)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		atomic.AddInt32(&fail, 1)
		log.Print("new request err: ", err)
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
		atomic.AddInt32(&fail, 1)
		return
	}
	b := bytes.NewBuffer(jsonStr)
	req.Body = ioutil.NopCloser(b)
	md5Key := utils.GenMd5Key(fbID)
	req.Header.Add("Authorization", "Bearer "+md5Key)
	apiResp, err := client.Do(req)
	if err != nil {
		logger.Print("test err", err.Error())
		atomic.AddInt32(&fail, 1)
		return
	}
	log.Print("seed play status:", apiResp.Status)
	if apiResp.StatusCode == 200 {
		req, err := http.NewRequest("POST", HOST+"/api/active/hero/record", nil)
		req.Header.Add("Content-Type", "application/json; charset=UTF-8")
		playResp := &struct {
			Data *ent.UserActiveRecord `json:"data"`
		}{}
		resp, err := ioutil.ReadAll(apiResp.Body)
		if err != nil {
			logger.Print("read apiResp err ", err.Error())
			atomic.AddInt32(&fail, 1)
			return
		}
		err = json.Unmarshal(resp, playResp)
		if err != nil {
			logger.Print("unmarshal playResp err ", err.Error())
			atomic.AddInt32(&fail, 1)
			return
		}

		recordReqeust := &hero.RecordRequest{
			FbUserID: fbID,
			RecordID: playResp.Data.ID,
			Score:    rand.Intn(300),
		}

		jsonStr, err := json.Marshal(recordReqeust)
		if err != nil {
			log.Print("marshal err: ", err.Error())
			atomic.AddInt32(&fail, 1)
			return
		}
		b := bytes.NewBuffer(jsonStr)
		req.Body = ioutil.NopCloser(b)
		md5Key := utils.GenMd5Key(fbID)
		req.Header.Add("Authorization", "Bearer "+md5Key)
		apiResp, err := client.Do(req)
		if err != nil {
			atomic.AddInt32(&fail, 1)
			logger.Print("test record err", err.Error())
			return
		}
		atomic.AddInt32(&success, 1)
		apiResp.Body.Close()
		log.Print("seed record record:", apiResp.Status)
		return
	} else {
		atomic.AddInt32(&fail, 1)
		b, _ := ioutil.ReadAll(apiResp.Body)
		log.Print("resp body: ", string(b))
	}
}
