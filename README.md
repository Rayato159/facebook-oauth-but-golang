<h1>🐳 MongoDb Docker Run Container</h1>

pull mongo image
```bash
docker pull mongo
```
start container
```bash
docker run 
    -e MONGO_INITDB_ROOT_USERNAME=root 
    -e MONGO_INITDB_ROOT_PASSWORD=123456 
    -p 27017:27017 
    -v data:/data
    --name oauth-db-dev 
    -d mongo:latest
```
configuration (before into the container)
```bash
docker exec -it oauth-db-dev bash
```
configuration (on the container)
```bash
mongosh mongodb://root:123456@127.0.0.1:27017
```
```bash
use oauth
```
```bash
show dbs
```

output
```bash
admin   100.00 KiB
config  108.00 KiB
local    72.00 KiB
oauth    80.00 KiB
```

<h1>🍃 MongoDb Migration</h1>

```bash
cd ./package/database/mongodb/albums-migrations
```

install dependencies
```bash
npm install
```

up or down
```bash
migrate-mongo [up|down]
```

<h1>📱 Facebook API</h1>

<h2>References</h2>
<ul>
  <li>Login Flow -> <a href="https://developers.facebook.com/docs/facebook-login/guides/advanced/manual-flow" target="_blank">https://developers.facebook.com/docs/facebook-login/guides/advanced/manual-flow</a></li>
  <li>Access Token -> <a href="https://developers.facebook.com/docs/facebook-login/guides/access-tokens" target="_blank">https://developers.facebook.com/docs/facebook-login/guides/access-tokens</a></li>
  <li>Revoke Token -> <a href="https://developers.facebook.com/docs/facebook-login/guides/permissions/request-revoke" target="_blank">https://developers.facebook.com/docs/facebook-login/guides/permissions/request-revoke</a></li>
</ul>

<h2>Base Url</h2>

```bash
https://graph.facebook.com/v15.0
```

<h2>API List</h2>
<ul>
    <li><a href="#redirect_url">Redirect Url</a></li>
    <li><a href="#get_access_token">Get Access Token</a></li>
    <li><a href="#get_profile">Get Profile</a></li>
    <li><a href="#revoke_token">Revoke Token (Logout)</a></li>
    <li><a href="#error_response">Error Response</a></li>
</ul>

<h2 id="redirect_url">Redirect Url</h2>

permission request (scopes)

```bash
public_profile
email
```

for more permission lists -> <a href="https://developers.facebook.com/docs/permissions/reference" target="_blank">https://developers.facebook.com/docs/permissions/reference</a>

request url

```bash
# Note that the redirect url in this part no need to parse any query params "JUST PURE REDIRECT URL"
GET https://www.facebook.com/v15.0/dialog/oauth?
  client_id={app_id}
  &redirect_uri={redirect_url}
  &state={hashu_access_token}
  &scope=public_profile,email
```

example of request url to get a redirect

```bash
https://localhost:3000/v1/facebook/login
```

<h2 id="get_access_token">Get Access Token</h2>

request url

```bash
GET https://graph.facebook.com/v15.0/oauth/access_token?
   client_id={app-id}
   &redirect_uri={redirect-uri}
   &client_secret={app-secret}
   &code={code-parameter}
```

response

```json
{
  "access_token": {access-token},
  "token_type": {type},
  "expires_in":  {seconds-til-expiration}
}
```

<h2 id="get_profile">Get Profile</h2>

request url

```bash
GET https://graph.facebook.com/v15.0/me?
   fields=id,email,name
   &access_token={access_token}
```

response

```json
{
    "id": "1234567890123456",
    "email": "nouzen@example.com",
    "name": "Shinei Nouzen"
}
```

<h2 id="revoke_token">Revoke Token</h2>

request url

```bash
DELETE https://graph.facebook.com/v15.0/{user-id}/permissions
```

response

```json
{
    "status": true
}
```

<h2 id="error_response">Error Response</h2>

**code 190** -> invalid token

```json
{
    "error": {
        "message": "The access token could not be decrypted",
        "type": "OAuthException",
        "code": 190,
        "fbtrace_id": "AZRUxySOh6GeksCUrq-uRZ2"
    }
}
```

**code 190** -> empty token

```json
{
    "error": {
        "message": "An active access token must be used to query information about the current user.",
        "type": "OAuthException",
        "code": 2500,
        "fbtrace_id": "AlBO3fDa-A_v3FoK8Sby9Ed"
    }
}
```
