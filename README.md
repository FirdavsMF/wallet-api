# Test wallet API for alif

## Before instalation
 You need install [docker](https://docs.docker.com/engine/install/ubuntu/) and [go](https://go.dev/doc/install)
## Installation 
- Run `docker-compose up -d`
- Go to folder `commands` and run `backup.sh` to import DB data
- Now you call APIs
## API
All requests sent by POST - method. All requests except `/login` should have header `X-UserId` and `X-Digest`
`X-UserId` - Id of user 
`X-Digest` - hash of body request with sha1 encondig
- `/login` - login and auth user. Format:

**Request:**
 ```sh
 {
  	"login": "username",      
	 "password":"password",      
}
```
**Response:**
 ```sh
 {
  	"messsage": "success" ,// string
 } 
```

- `v1/checkwallet` - check wallet for exist.
 
**Request:**

```sh
{
  "login": "wallet login", 
}
```
**Response:**
```sh
{
  "message": " wallet  SomeWallet exists" // string 
}
```

- `v1/charge` - Charging wallet from one to another

**Request:**
```sh
{
    "src": "Source wallet login",
    "dest": "Destination wallet login",
    "sum": "some in float format"
}
```
**Response:**
```sh
{
    "message": {
    "Payment proccesed","transactionID": tranID // string
    } 
}
```
- `v1/gethistory` - Total count and sum operations of wallet in current month

**Request:**

```sh
{
    "login": "wallet"
}
```
**Response:**

```sh
{
  "data": {
      "Count": "some count", // int
      "Sum": "some sum" // float
  }
}
```

- `v1/getbalance` - Total balance of considering wallet

**Request:**
```sh
{
    "login": "wallet"
}
```
**Response:**
```sh
{
"message":{
    "WalletLogin":"SomeWallet", // string
    "balance":"wallet sum" // float
 }
}

```

