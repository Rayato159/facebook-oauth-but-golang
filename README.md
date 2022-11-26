<h1>Facebook API</h1>

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
    <li><a href="#revoke_token">Revoke Token (Logout)</a></li>
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