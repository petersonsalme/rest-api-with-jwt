# REST API with JWT

A Go-based REST API showcasing authentication using JSON Web Tokens (JWT). 

## Endpoints

- **`POST /login`**: Authenticate using username and password (in JSON body). Returns `access_token` and `refresh_token`.
- **`POST /logout`**: Invalidate the user's tokens (requires `Authorization: Bearer <access_token>`).
- **`POST /token/refresh`**: Get a new `access_token` by sending a JSON body with `{"refresh_token": "..."}`.
- **`POST /todo`**: Create a new Todo item (requires `Authorization: Bearer <access_token>`).

## How to Authenticate

1. Call `/login` with your credentials.
2. The response will contain an `access_token`.
3. For protected endpoints (like `/todo` and `/logout`), include the token in your headers:
   `Authorization: Bearer <access_token>`

## API Specification

This repository contains an OpenAPI 3.0 specification file: `openapi.yaml`. You can import this file into Swagger UI, Postman, or Insomnia to explore and test the endpoints.

## Security Best Practices Followed

- **Algorithm Verification**: The application verifies that the JWT signing method is specifically HMAC (`jwt.SigningMethodHMAC`). This prevents the infamous algorithm confusion vulnerability where attackers could supply a token signed with the `none` algorithm or an asymmetric public key.
- **Dependency Management**: Uses the updated and actively maintained `github.com/golang-jwt/jwt/v5` package instead of the deprecated `dgrijalva/jwt-go` which had known vulnerabilities.
- **Secure Secrets**: The `ACCESS_SECRET` and `REFRESH_SECRET` must be set via environment variables and should be long, random strings in a production environment.
- **Refresh Tokens**: Uses short-lived access tokens and longer-lived refresh tokens that can be revoked or rotated. Refresh tokens are tracked in Redis.

## Important Note for Deployment
Ensure the application is deployed behind HTTPS (TLS) to prevent man-in-the-middle attacks from intercepting the JWTs.