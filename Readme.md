# FinBuild

Build your Finance Life

![Go Report Card](https://goreportcard.com/badge/github.com/tonnytg/finbuild)


### How to use

To request a user balance <br/>
Method: `GET` <br/>
`http://localhost:8888`


To request a user informations <br/>
Method: `GET` <br/>
`http://localhost:8888/user`


To input a exchange <br/>
Method: `POST` <br/>
`http://localhost:8888/exchange`
```
{
    "id": "PRIO3",
    "price": 100,
    "quantity": 10,
    "action": "BUY",
    "date": "2021/01/01 10:04:05"
}
```