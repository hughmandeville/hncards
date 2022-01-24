# Get Hacker News Top Stories

```sh
go run get_hn.go -out ../../client/public/hn_topstories.json

jq '.[] | {publisher,url}' ../../client/public/hn_topstories.json
```
