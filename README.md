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

200 OK
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

200 OK
{
  "success": true,
  "message": "User successfully logged in"
}
```

### GET /game/start

Sends a signal to the server to start the game.

```http-request
GET localhost:3001/game/start

200 OK
{
  "dealerCardCount": 2
  "dealerCards": ["4-S"]
  "playerCardCount": 2
  "playerCards": ["7-C", "A-S"]
}
```

### GET /forgotPassword

Gets a password for a given user (most insecure forgot password method, but lazy to do an email client)
If the user is found, the message will contain the password.

*Example:*
```http-request
GET localhost:3001/forgotPassword

{
  "username": "<username>"
}

200 OK
{
  "success": true,
  "message": "<password>"
}
```

[^1] API can be reached at `http://localhost:3001`
