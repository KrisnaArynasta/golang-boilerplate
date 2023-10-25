# Welcome
### Simple project structure for Go applications.
##
##
## .env Dummy Example:

```
APP_ENV=development
SERVER_ADDRESS=:xxx
PORT=0000
CONTEXT_TIMEOUT=0
DB_HOST=xxx
DB_PORT=0000
DB_USER=xxx
DB_PASS=xxx
DB_NAME=xxx
ACCESS_TOKEN_EXPIRY_HOUR = 0
REFRESH_TOKEN_EXPIRY_HOUR = 0
ACCESS_TOKEN_SECRET=xxx
REFRESH_TOKEN_SECRET=xxx
EXTERNAL_API_BASE_URL=xxx
```

## API Call
Get Data from Database
```
http://localhost:8080/get_db
http://localhost:8080/get_db?id=139
```

Get Data from External API
```
http://localhost:8080/get_data
http://localhost:8080/get_data?method=BANK_TRANSFER
```

Post Data to External API and Save to Database
```
http://localhost:8080/post_data
```