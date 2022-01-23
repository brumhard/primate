# Authentication

- multiple scenarios: web, native apps

## Web

- using grpc-web instead of native grpc over http/2

### oauth2_proxy

- can easily use oauth2_proxy as reverse proxy and it will instantly work for the web with social login for example github
- filter allowed users by email
- oauth2_proxy can also be deployed as a separate service in k8s and used for remote_auth for nginx/traefik ingress
  - https://blog.codecentric.de/en/2021/06/how-to-use-oauth2-proxy-for-central-authentication/
  - or also just as a sidecar for a single service

## native apps

- using "real" grpc over http/2
- need to implement signin with redirects to social logins myself (sth like https://github.com/Clancey/simple_auth/tree/master/simple_auth could help but would also make sense to implement myself)
- redirect to browser for login
- callback url would be something like myapp:// that is redirected to the app locally to access and safe the token (how would that work on macos/windows?)
- how to check on app startup, whether i'm logged in or not?
  - -> could be using the oauth2_proxy /oauth2/userinfo or /oauth2/auth endpoint or provider specific endpoint like "https://api.github.com/user"

### oauth2_proxy

- uses cookies to check authentication state
- can technically also validate jwt keys but not for all providers (only for oidc providers/ for providers that implement "https://github.com/oauth2-proxy/oauth2-proxy/blob/master/providers/providers.go" CreateSessionFromToken since oauth2_proxy needs the infos to validate the session email etc.)
  - progress on this is tracked for github here: https://github.com/oauth2-proxy/oauth2-proxy/issues/1499
  - -> currently only oidc would be possible
- not sure if that jwt key validation would also work for oauth2_proxy
  - different format on the wire
  - example: wireshark with settings shown here (https://grpc.io/blog/wireshark/) and filter like (tcp.port == 8124 and http2) or (tcp.port == 8123 and http)
  - oauth2_proxy uses "auth := req.Header.Get("Authorization")" to extract the token, does this also work in http2 and grpc metadata
    - loooks like grpc metadata is put into headers (https://chromium.googlesource.com/external/github.com/grpc/grpc/+/HEAD/doc/PROTOCOL-HTTP2.md)
  - could be tested easily with a simple http server in go that automatically uses http/2 when possible as well as `curl --http2`
    - headers are also available in http.Request.Header when usign http2 with curl and grpcurl (look at ~/repos/personal/http2_test)
  - should be tested with grpc client and set metadata to assure that the values are also available there
      - from what is implemented here: https://github.com/fullstorydev/grpcurl/blob/ae7473c7a7e1172a5418377b64000d55697c2282/grpcurl.go#L172 it looks like the -H option is directly translated to metadata in client request
  - should now be tested with a mock oidc provider and oauth2_proxy if authorization in metadata is picked of nicely

### custom implementation

- could take https://github.com/markbates/goth as foundation and add multiple flow implementations (e.g. cookies callback for web, jwt token returned for native apps...)
- also return 403/gprc permission denied code when native app flow, for web redirect to sign_in page
- could be tricky for such things as token refreshes etc.
