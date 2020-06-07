
# 建置 migrate  
使用 facebook/ent  
  
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
  
### 開始遊戲
- Method: **POST**
- URL:  ```/api/active/hero/play```
- Headers：```md5(fb_user_id)```
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
- URL:  ```/api/active/hero/record```
- Headers：```md5(fb_user_id)```
- Body:
```
{
    "record_id": "UAR_breaql3ipt3dp3hsffq0",
    "score": 1000
}
```

### 遊戲埋點
- Method: **POST**
- URL:  ```/api/active/hero/tracking```
- Headers：```md5(fb_user_id)```
- Body:
```
{
    "fb_user_id": "12345684",
    "fb_name": "frankie",
    "type":"hero_share" //分享：hero_share, 下載:hero_download
}
```