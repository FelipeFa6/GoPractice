# como ejecutar la API

```bash
$ go run .
```

## Ejemplo de llamadas cURL

```bash
#!/bin/sh

echo "Creating a new account..."
CREATE_ACCOUNT_RESPONSE=$(curl -X POST http://localhost:3000/account -s -H "Content-Type: application/json" -d '{"username": "FelipeFarias"}')
echo "Account created: $CREATE_ACCOUNT_RESPONSE"
ACCOUNT_ID=$(echo $CREATE_ACCOUNT_RESPONSE | jq -r '.id')
echo "Account ID: $ACCOUNT_ID"

# Test: Get account details (GET request)
echo "Getting account with ID $ACCOUNT_ID..."
GET_ACCOUNT_RESPONSE=$(curl -X GET http://localhost:3000/account/$ACCOUNT_ID -s)
echo "Account details: $GET_ACCOUNT_RESPONSE"

# Test: Create a playlist for the account (POST request to /account/{id}/playlist)
echo "Creating a playlist for account $ACCOUNT_ID..."
CREATE_PLAYLIST_RESPONSE=$(curl -X POST http://localhost:3000/account/$ACCOUNT_ID/playlist -s -H "Content-Type: application/json" -d '{"playlistName": "My Favorite Playlist"}')
echo "Playlist created: $CREATE_PLAYLIST_RESPONSE"

# Test: Get the account again to verify playlist creation (GET request)
echo "Getting account with playlists..."
GET_ACCOUNT_WITH_PLAYLISTS_RESPONSE=$(curl -X GET http://localhost:3000/account/$ACCOUNT_ID -s)
echo "Updated account details with playlists: $GET_ACCOUNT_WITH_PLAYLISTS_RESPONSE"

# Keep the script running to simulate a long-running process
while true; do
    sleep 1
done
```
