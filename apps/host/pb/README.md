```
protoc -I=. --go_out=. --go_opt=module="github.com/restfulapi-demo" app/*/pb/*.proto
```