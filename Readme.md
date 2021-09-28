# FinBuild

Build your Finance Life

![Go Report Card](https://goreportcard.com/badge/github.com/tonnytg/finbuild)


### How to use

To request a user balance <br/>
Method: `POST` <br/>
`http://localhost:8888/user`
```
{
    "id": "00000000-0000-0000-0000-000000000000",
    "first_name": "Tonnytg",
    "last_name": "TG",
    "social_id": "001.001.001.01",
    "phone": "11946302400",
    "email": "tonnytg@gmail.com",
    "valid_email": true,
    "address": "Brazil",
    "age": 30,
    "sex": "male",
    "sign": "Aquarius"
}
```

<br />

To request a user balance <br/>
Method: `GET` <br/>
`http://localhost:8888?id=00000000-0000-0000-0000-000000000000`

Return

```
{
    "status": "success",
    "data": [
        {
            "user": {
                "id": "00000000-0000-0000-0000-000000000000",
                "first_name": "Tonnytg"
            }
        },
        {
            "finance": {
                "Balance": 2634.01
            }
        }
    ],
    "message": "test"
}
```

<hr/>


To request a user informations <br/>
Method: `GET` <br/>
`http://localhost:8888/user?id=00000000-0000-0000-0000-000000000000`

Return
```
{
    "status": "success",
    "data": [
        {
            "exchange": {
                "id": "00000000-0000-0000-0000-000000000000",
                "first_name": "Tonnytg",
                "last_name": "TG",
                "social_id": "001.001.001.01",
                "phone": "11946302400",
                "email": "tonnytg@gmail.com",
                "valid_email": true,
                "address": "Brazil",
                "age": 30,
                "sex": "male",
                "sign": "Aquarius"
            }
        }
    ],
    "message": "test"
}
```


<hr/>

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

Return
```
{
    "status": "success",
    "data": [
        {
            "exchange": {
                "user_id": "00000000-0000-0000-0000-000000000000",
                "id": "PRIO3",
                "price": 100,
                "quantity": 10,
                "action": "BUY",
                "date": "2021/01/01 10:04:05"
            }
        }
    ],
    "message": "test"
}
```