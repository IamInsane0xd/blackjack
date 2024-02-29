# BlackJack project

Web based BlackJack game made as a school project<br>
*Authors: Krisztian Bori, Ferenc Biczo*

## API Endpoints[^1]
### POST /register

Registers a new user with a username and password combo.<br>
The endpoint will respond with `200 OK`, unless it's an internal server error, and the response will contain the result and a message if needed (think error message when failed for example).<br>

*Example:*
```http-request
POST localhost:3001/register

{
  "username": "<username>",
  "password": "<password>"
}

{
  "success": false,
  "message": "User already exists"
}
```

### POST /login

Computes a login request.<br>
This endpoint will also respond with `200 OK` unless it's an internal server error, and the response will contain the result.<br>

*Example:*
```http-request
POST localhost:3001/login

{
  "username": "<username>",
  "password": "<password>"
}

{
  "success": true,
  "message": "User successfully logged in"
}
```

[^1] API can be reached at `http://localhost:3001`
