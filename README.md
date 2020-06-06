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