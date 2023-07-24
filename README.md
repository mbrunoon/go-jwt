# Go API with JWT Auth

 Golang Authentication with JWT, Gin-Gonic and MongoDB.

## Tutorial

Videos by [Akhil Sharma](https://www.youtube.com/@AkhilSharmaTech)

[Part 1: Building App structure, User Model and Routes](https://www.youtube.com/watch?v=X5BmktFrAlo)
[Part 2: Creating Database and Controller and Routes especifications](https://www.youtube.com/watch?v=LCWrfrsZARo)
[Part 3: Creating GetUser and SignUp methods](https://www.youtube.com/watch?v=8nUKNJqgYLo)
[Part 4: Creating tokenHelper, add login and verify password methods to userController](https://www.youtube.com/watch?v=uaydySiRU9M)
[Part 5: Creating GetUsers method](https://www.youtube.com/watch?v=kCCdf0Ytcts)
[Part 6: Bugfixes](https://www.youtube.com/watch?v=dHSs4XSkCdk)

# Tools Used

- [Gin-Gonic](https://github.com/gin-gonic/gin): web framework;
- [MongoDB](https://github.com/mongodb/mongo-go-driver): NoSQL database;
- [DotEnv](https://github.com/joho/godotenv): loads env vars from a .env file;
- [JWT](https://github.com/golang-jwt/jwt): A Go implementation of JSON Web Tokens;

# Environment variables
```
PORT=9000
MONGODB_URL=yourmongodb-url
SECRET_KEY="your-secure-hash"
```

# Routes

#### POST /users/signup

**Request sample:**
```
curl --location 'localhost:9000/users/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "First Name",
    "last_name": "Last Name",
    "password": "yourpassword",
    "email": "your@email.com",
    "phone": "0000000",
    "user_type": "ADMIN"
}'
```

**Expected Response:**

*HTTP Status 200*
```
{
    "InsertedID": "64b8e7451c68e298739528ad"
}
```
----
#### POST /users/login

**Request sample:**

```
curl --location 'localhost:9000/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "your@email.com",
    "password": "yourpassword"
}'
```

**Expected Response:**

*HTTP Status 200*

```
{
    "ID": "64ba3cdf324d062cc21df5b8",
    "first_name": "First Name",
    "last_name": "Last Name",
    "password": "$2a$10$8Hy0ohEtKvw3kycL5eLmae1VYcc21IjzBq8..XtEYE2WTtc/S2v32",
    "email": "your@email.com",
    "phone": "0000000",
    "json_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImNvbnRhdG9AbWFyY29zYnJ1bm8uY29tIiwiRmlyc3ROYW1lIjoiTWFyY29zIiwiTGFzdE5hbWUiOiJCcnVubyIsIlVpZCI6IjY0YmEzY2RmMzI0ZDA2MmNjMjFkZjViOCIsIlVzZXJUeXBlIjoiQURNSU4iLCJleHAiOjE2OTAwMTQwMTd9.8CGXCfcwRX_Y_71mlowD3nbr9TOXEjft6bgEqMtzx2w",
    "user_type": "ADMIN",
    "RefreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0TmFtZSI6IiIsIkxhc3ROYW1lIjoiIiwiVWlkIjoiIiwiVXNlclR5cGUiOiIiLCJleHAiOjE2OTA1MzE2Nzl9._nnANVM-Q__kX3mWgSIwIlOcjexLTtpyxz1EVcrzAmA",
    "created_at": "2023-07-21T08:07:59Z",
    "updated_at": "2023-07-21T08:07:59Z",
    "user_id": "64ba3cdf324d062cc21df5b8"
}
```
---
#### GET /users


**Request sample:**
```
curl --location 'localhost:9000/users' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImNvbnRhdG9AbWFyY29zYnJ1bm8uY29tIiwiRmlyc3ROYW1lIjoiTWFyY29zIiwiTGFzdE5hbWUiOiJCcnVubyIsIlVpZCI6IjY0YmEzY2RmMzI0ZDA2MmNjMjFkZjViOCIsIlVzZXJUeXBlIjoiQURNSU4iLCJleHAiOjE2OTAwMTQwMTd9.8CGXCfcwRX_Y_71mlowD3nbr9TOXEjft6bgEqMtzx2w'
```

**Expected Response:**

*HTTP Status 200*

```
{
    "total_count": 2,
    "user_items": [
        {
            "_id": "64ba3cdf324d062cc21df5b9",
            "createdat": "2023-07-21T08:07:59Z",
            "email": "your@email.com",
            "firstname": "First Name",
            "id": "64ba3cdf324d062cc21df5b8",
            "lastname": "Last Name",
            "password": "$2a$10$8Hy0ohEtKvw3kycL5eLmae1VYcc21IjzBq8..XtEYE2WTtc/S2v32",
            "phone": "0000000",
            "refreshtoken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0TmFtZSI6IiIsIkxhc3ROYW1lIjoiIiwiVWlkIjoiIiwiVXNlclR5cGUiOiIiLCJleHAiOjE2OTA1MzE2Nzl9._nnANVM-Q__kX3mWgSIwIlOcjexLTtpyxz1EVcrzAmA",
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImNvbnRhdG9AbWFyY29zYnJ1bm8uY29tIiwiRmlyc3ROYW1lIjoiTWFyY29zIiwiTGFzdE5hbWUiOiJCcnVubyIsIlVpZCI6IjY0YmEzY2RmMzI0ZDA2MmNjMjFkZjViOCIsIlVzZXJUeXBlIjoiQURNSU4iLCJleHAiOjE2OTAwMTMyNzl9.TiufFgOPRQM8iOR_e_lSeEDcU3JQjZ2ii_DT30EGRFc",
            "updatedat": "2023-07-21T08:07:59Z",
            "userid": "64ba3cdf324d062cc21df5b8",
            "usertype": "ADMIN"
        },
        {
            "_id": "64ba3fc15ae4496667345c7b",
            "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0TmFtZSI6IiIsIkxhc3ROYW1lIjoiIiwiVWlkIjoiIiwiVXNlclR5cGUiOiIiLCJleHAiOjE2OTA1MzI0MTd9.oaPA0KAg5Byvo-f7IiyswavDHIO7o0GyzdFS5t7WoiI",
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImNvbnRhdG9AbWFyY29zYnJ1bm8uY29tIiwiRmlyc3ROYW1lIjoiTWFyY29zIiwiTGFzdE5hbWUiOiJCcnVubyIsIlVpZCI6IjY0YmEzY2RmMzI0ZDA2MmNjMjFkZjViOCIsIlVzZXJUeXBlIjoiQURNSU4iLCJleHAiOjE2OTAwMTQwMTd9.8CGXCfcwRX_Y_71mlowD3nbr9TOXEjft6bgEqMtzx2w",
            "updated_at": "2023-07-21T08:20:17Z",
            "userId": "64ba3cdf324d062cc21df5b8"
        }
    ]
}
```
---
#### GET /users/:user_id

**Request sample:**

```
curl --location 'localhost:9000/users/64ba3cdf324d062cc21df5b8' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImNvbnRhdG9AbWFyY29zYnJ1bm8uY29tIiwiRmlyc3ROYW1lIjoiTWFyY29zIiwiTGFzdE5hbWUiOiJCcnVubyIsIlVpZCI6IjY0YmEzY2RmMzI0ZDA2MmNjMjFkZjViOCIsIlVzZXJUeXBlIjoiQURNSU4iLCJleHAiOjE2OTAyNjkwMjJ9.-DSkCN2Y8NQYEU_f_l1v-YcTXxtkeQIcNarixLNtygw'
```

**Expected Response:**

*HTTP Status 200*

```
{
    "ID": "64ba3cdf324d062cc21df5b8",
    "first_name": "First Name",
    "last_name": "Last Name",
    "password": "$2a$10$8Hy0ohEtKvw3kycL5eLmae1VYcc21IjzBq8..XtEYE2WTtc/S2v32",
    "email": "your@email.com",
    "phone": "0000000",
    "json_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImNvbnRhdG9AbWFyY29zYnJ1bm8uY29tIiwiRmlyc3ROYW1lIjoiTWFyY29zIiwiTGFzdE5hbWUiOiJCcnVubyIsIlVpZCI6IjY0YmEzY2RmMzI0ZDA2MmNjMjFkZjViOCIsIlVzZXJUeXBlIjoiQURNSU4iLCJleHAiOjE2OTAwMTMyNzl9.TiufFgOPRQM8iOR_e_lSeEDcU3JQjZ2ii_DT30EGRFc",
    "user_type": "ADMIN",
    "RefreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0TmFtZSI6IiIsIkxhc3ROYW1lIjoiIiwiVWlkIjoiIiwiVXNlclR5cGUiOiIiLCJleHAiOjE2OTA1MzE2Nzl9._nnANVM-Q__kX3mWgSIwIlOcjexLTtpyxz1EVcrzAmA",
    "created_at": "2023-07-21T08:07:59Z",
    "updated_at": "2023-07-21T08:07:59Z",
    "user_id": "64ba3cdf324d062cc21df5b8"
}
```