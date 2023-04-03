## shop challenge
### endpoints:
* {host}/card/transfer [POST]
  * body:
```json
  {
   "sourceCardNo": "sourceCardNo",
   "amount": "amount",
   "targetCardNo": "targetNo"
  }
```

* {host}/localhost:9001/customer/active/transactions [GET]

### sample curl:
1. curl --location --request GET 'http://localhost:9001/customer/active/transactions'
2. curl --location --request POST 'http://localhost:9001/card/transfer' \
   --header 'Accept-Charset;' \
   --header 'Content-Type: application/json' \
   --data-raw '{
   "sourceCardNo": "6037703932049631",
   "amount": "50000",
   "targetCardNo": "6104337743428383"
   }
   '

--- 
##### note:
for each transfer transaction, two transaction will be inserted to table as follow,so all 
the withdrawal and deposit txns will be saved independently of the other account:

| from-account | to-account | amount |
|:-------------|:----------:|-------:|
| 6037703932049631   | 6104337743428383 |  50000 |
| 6104337743428383     | 6037703932049631  | -50000 |

