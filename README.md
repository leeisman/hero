
  
# 建置 migrate 使用 facebook/ent    
    
- Gen Table Schema file
 ```    
database > entc init <table>    
```    
- Gen Table migrate file
 ```    
database >  entc generate ./ent/schema     
or    
>  entc generate ./database/ent/schema     
```    
    
# API    

### 安全性驗證

Md5 雜湊加上 fb user ID：
- header帶上: ```Authorization:Bearer md5(fb_user_id)```


### 開始遊戲  
- Method: **POST**  
- URL: ```/api/active/hero/play```  
- Body:  
```  
{  
    "fb_user_id": "12345684",  
    "fb_avatar_url": "http://localhost.hero.com/123",  
    "fb_email": "frankie.lee.job@gmail.com",  
    "fb_name": "frankie"  
}  
```  
  
### 完成紀錄遊戲  
- Method: **POST**  
- URL: ```/api/active/hero/record```  
- Body:  
```  
{  
    "fb_user_id": "12345684",  
    "record_id": "UAR_breaql3ipt3dp3hsffq0", //從開始遊戲生成
    "score": 1000  
}  
```  
  
### 遊戲埋點  
- Method: **POST**  
- URL: ```/api/active/hero/tracking```   
- Body:  
```  
{  
    "fb_user_id": "12345684",  
    "fb_name": "frankie",  
    "type":"hero_share" //分享：hero_share, 下載:hero_download  
}  
```

### 遊戲紀錄報表  
- Method: **GET**  
- URL: ```/api/report/hero/count```   
- Body:  
```  
{
    "start_at": "2020-06-07 00:00:00",
    "end_at": "2020-06-08 23:59:59"
} 
```

### 遊戲答題人數統計
- Method: **GET**  
- URL: ```/api/report/hero/score_count```   
- Body:  
```  
{
    "score": 1000,
    "start_at": "2020-06-09 00:00:00",
    "end_at": "2020-06-11 23:59:59"
}
```

### 即時積分排行榜，前10名及玩家自己的(本次答對題數、最佳成績答對題數、
- Method: **GET**  
- URL: ```/api/report/hero/rank```
