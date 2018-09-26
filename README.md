## Welcome

Tax Calculator:

### Quick Start

1. Clone the repo
  ```
  $ git clone https://github.com/ramb0111/tax-calc.git
  $ cd tax-calc
  ```

2. Run Docker Compose
```
$ docker-compose up --build
```

NOTE: Logs will come like this
```
2018/09/26 16:04:02 [Postgres Connect] Port: 5432
2018/09/26 16:04:02 [API Server] Port: 3000
```

## Requests
### POST:
  ```
  http://localhost:3000/tax/5
  
  BODY:
    {
        "name" : "anky",
        "taxcode" : 3,
        "amount" : 1000
    }

  Response:
    {
        "status": "success"
    }

  ```

### GET:
  ```
  http://localhost:3000/tax/5

  Response:
    {
        "Items": [
            {
                "Name": "anky",
                "Amount": 1000,
                "Tax Amount": 9,
                "Total Amount": 1009,
                "Tax Code": 3,
                "Type": "entertainment"
            }
        ],
        "Total Amount": 1000,
        "Total Tax Amount": 9,
        "Grand Total": 1009
    }
 

  http://localhost:3000/tax/4

  Response:
    {
        "Items": [],
        "Total Amount": 0,
        "Total Tax Amount": 0,
        "Grand Total": 0
    }

  ```



