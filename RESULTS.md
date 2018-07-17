## Service test with endpoint test

Using curl:

```
$ curl localhost:8080 -d'{"url":"http://www.example.org", "body": "sample payload"}'
$ curl localhost:8080 -d'{"url":"http://www.example.org", "body": "sample payload"}'
```

Results

```
$ mesg-core service test
Image built with success
Image hash: sha256:6202fba0119b9f7a14af92efa0bdb07d0b19d20453214661b1d09b70ff9c191d
Service deployed with success
Service ID: v1_669f52e72fe9d7af659d2a1dfe2d049b
Service started
Listening for results from the service...
Listening for events from the service...
2018/07/17 10:14:20 Serving endpoint at [:8080]
2018/07/17 22:15:22 Receive event onRequest : {"body":"{\"url\":\"http://www.example.org\", \"body\": \"sample payload\"}","date":"2018-07-17 10:15:25","id":"525fe2a3-89aa-11e8-b8cf-02420aff000c"}
2018/07/17 10:15:25 endpoint: reply from mesg:
2018/07/17 22:16:25 Receive event onRequest : {"body":"{\"url\":\"http://www.example.org\", \"body\": \"sample payload\"}","date":"2018-07-17 10:16:28","id":"77ea63d0-89aa-11e8-b8cf-02420aff000c"}
2018/07/17 10:16:28 endpoint: reply from mesg:
```

## Test of execute task

Using input: [execute-task-data.json](test-inputs/execute-task-data.json)

```
$ mesg-core service test --logs-all --task execute --data ./test-inputs/execute-task-data.json
Image built with success
Image hash: sha256:8813a5ea366b858e6e67b2356fa97caf817a59a4e2c890d304cc864c2d97b395
Service deployed with success
Service ID: v1_e7d4eeca8dbd632d7498924e5610cc1c
Service started
Listening for events from the service...
Listening for results from the service...
2018/07/17 11:04:48 Serving endpoint at [:8080]
2018/07/17 23:04:48 Execute task execute with data {
  "url":"http://www.example.org",
  "body": "sample payload"
}

2018/07/17 23:04:49 Receive result execute success with data {"body":"\u003c!doctype html\u003e\n\u003chtml\u003e\n\u003chead\u003e\n    \u003ctitle\u003eExample Domain\u003c/title\u003e\n\n    \u003cmeta charset=\"utf-8\" /\u003e\n    \u003cmeta http-equiv=\"Content-type\" content=\"text/html; charset=utf-8\" /\u003e\n    \u003cmeta name=\"viewport\" content=\"width=device-width, initial-scale=1\" /\u003e\n    \u003cstyle type=\"text/css\"\u003e\n    body {\n        background-color: #f0f0f2;\n        margin: 0;\n        padding: 0;\n        font-family: \"Open Sans\", \"Helvetica Neue\", Helvetica, Arial, sans-serif;\n        \n    }\n    div {\n        width: 600px;\n        margin: 5em auto;\n        padding: 50px;\n        background-color: #fff;\n        border-radius: 1em;\n    }\n    a:link, a:visited {\n        color: #38488f;\n        text-decoration: none;\n    }\n    @media (max-width: 700px) {\n        body {\n            background-color: #fff;\n        }\n        div {\n            width: auto;\n            margin: 0 auto;\n            border-radius: 0;\n            padding: 1em;\n        }\n    }\n    \u003c/style\u003e    \n\u003c/head\u003e\n\n\u003cbody\u003e\n\u003cdiv\u003e\n    \u003ch1\u003eExample Domain\u003c/h1\u003e\n    \u003cp\u003eThis domain is established to be used for illustrative examples in documents. You may use this\n    domain in examples without prior coordination or asking for permission.\u003c/p\u003e\n    \u003cp\u003e\u003ca href=\"http://www.iana.org/domains/example\"\u003eMore information...\u003c/a\u003e\u003c/p\u003e\n\u003c/div\u003e\n\u003c/body\u003e\n\u003c/html\u003e\n","statusCode":200}
```

## Test of async executeMany task

Using input: [executeMany-task-data.json](test-inputs/executeMany-task-data.json)

```
$ mesg-core service test --logs-all --task executeMany --data ./test-inputs/executeMany-task-data.json
Image built with success
Image hash: sha256:bd9d0183bffa5e1dd5c4da056318321ed16c2d859f4233aa964b273af63c435d
Service deployed with success
Service ID: v1_c750784ed5fbc8a112234f7690bdd6d4
Service started
Listening for events from the service...
Listening for results from the service...
2018/07/17 11:01:49 Serving endpoint at [:8080]
2018/07/17 23:01:50 Execute task executeMany with data {
  "requests": [
    {
      "url":"http://www.example.org",
      "body": "sample payload"
    },
    {
      "url":"http://www.example.com",
      "body": "different payload"
    }
  ],
  "async": true
}

2018/07/17 23:01:50 Receive result executeMany success with data [{"error":{"message":""},"success":{"body":"\u003c!doctype html\u003e\n\u003chtml\u003e\n\u003chead\u003e\n    \u003ctitle\u003eExample Domain\u003c/title\u003e\n\n    \u003cmeta charset=\"utf-8\" /\u003e\n    \u003cmeta http-equiv=\"Content-type\" content=\"text/html; charset=utf-8\" /\u003e\n    \u003cmeta name=\"viewport\" content=\"width=device-width, initial-scale=1\" /\u003e\n    \u003cstyle type=\"text/css\"\u003e\n    body {\n        background-color: #f0f0f2;\n        margin: 0;\n        padding: 0;\n        font-family: \"Open Sans\", \"Helvetica Neue\", Helvetica, Arial, sans-serif;\n        \n    }\n    div {\n        width: 600px;\n        margin: 5em auto;\n        padding: 50px;\n        background-color: #fff;\n        border-radius: 1em;\n    }\n    a:link, a:visited {\n        color: #38488f;\n        text-decoration: none;\n    }\n    @media (max-width: 700px) {\n        body {\n            background-color: #fff;\n        }\n        div {\n            width: auto;\n            margin: 0 auto;\n            border-radius: 0;\n            padding: 1em;\n        }\n    }\n    \u003c/style\u003e    \n\u003c/head\u003e\n\n\u003cbody\u003e\n\u003cdiv\u003e\n    \u003ch1\u003eExample Domain\u003c/h1\u003e\n    \u003cp\u003eThis domain is established to be used for illustrative examples in documents. You may use this\n    domain in examples without prior coordination or asking for permission.\u003c/p\u003e\n    \u003cp\u003e\u003ca href=\"http://www.iana.org/domains/example\"\u003eMore information...\u003c/a\u003e\u003c/p\u003e\n\u003c/div\u003e\n\u003c/body\u003e\n\u003c/html\u003e\n","statusCode":200}},{"error":{"message":""},"success":{"body":"\u003c!doctype html\u003e\n\u003chtml\u003e\n\u003chead\u003e\n    \u003ctitle\u003eExample Domain\u003c/title\u003e\n\n    \u003cmeta charset=\"utf-8\" /\u003e\n    \u003cmeta http-equiv=\"Content-type\" content=\"text/html; charset=utf-8\" /\u003e\n    \u003cmeta name=\"viewport\" content=\"width=device-width, initial-scale=1\" /\u003e\n    \u003cstyle type=\"text/css\"\u003e\n    body {\n        background-color: #f0f0f2;\n        margin: 0;\n        padding: 0;\n        font-family: \"Open Sans\", \"Helvetica Neue\", Helvetica, Arial, sans-serif;\n        \n    }\n    div {\n        width: 600px;\n        margin: 5em auto;\n        padding: 50px;\n        background-color: #fff;\n        border-radius: 1em;\n    }\n    a:link, a:visited {\n        color: #38488f;\n        text-decoration: none;\n    }\n    @media (max-width: 700px) {\n        body {\n            background-color: #fff;\n        }\n        div {\n            width: auto;\n            margin: 0 auto;\n            border-radius: 0;\n            padding: 1em;\n        }\n    }\n    \u003c/style\u003e    \n\u003c/head\u003e\n\n\u003cbody\u003e\n\u003cdiv\u003e\n    \u003ch1\u003eExample Domain\u003c/h1\u003e\n    \u003cp\u003eThis domain is established to be used for illustrative examples in documents. You may use this\n    domain in examples without prior coordination or asking for permission.\u003c/p\u003e\n    \u003cp\u003e\u003ca href=\"http://www.iana.org/domains/example\"\u003eMore information...\u003c/a\u003e\u003c/p\u003e\n\u003c/div\u003e\n\u003c/body\u003e\n\u003c/html\u003e\n","statusCode":200}}]
2018/07/17 11:01:51
```
