## Curl command

to connect with the opensearch using the curl command, use th following command:

```
curl -XGET --insecure -u 'admin:admin' 'https://localhost:9200'
```

```
curl -X POST "http://localhost:9200/users/_doc" -H "Content-Type: application/json" -d '{
  "doc": {
    "@timestamp": "2023-07-26T14:49:12.642538Z",
    "name":"tushar"
}
}'
```
