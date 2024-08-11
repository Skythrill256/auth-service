module github.com/Skythrill256/auth-service

go 1.22.5

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.26.0
	golang.org/x/oauth2 v0.22.0
)

require cloud.google.com/go/compute/metadata v0.3.0 // indirect
