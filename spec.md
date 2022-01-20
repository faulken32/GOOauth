     +--------+                               +---------------+
     |        |--(A)- Authorization Request ->|   Resource    |
     |        |                               |     Owner     |
     |        |<-(B)-- Authorization Grant ---|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(C)-- Authorization Grant -->| Authorization |
     | Client |                               |     Server    |
     |        |<-(D)----- Access Token -------|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(E)----- Access Token ------>|    Resource   |
     |        |                               |     Server    |
     |        |<-(F)--- Protected Resource ---|               |
     +--------+                               +---------------+

     (A)  The client requests authorization from the resource owner.  The
        authorization request can be made directly to the resource owner
        (as shown), or preferably indirectly via the authorization
        server as an intermediary.

(B)  The client receives an authorization grant, which is a
credential representing the resource owner's authorization,
expressed using one of four grant types defined in this
specification or using an extension grant type.  The
authorization grant type depends on the method used by the
client to request authorization and the types supported by the
authorization server.

(C)  The client requests an access token by authenticating with the
authorization server and presenting the authorization grant.

(D)  The authorization server authenticates the client and validates
the authorization grant, and if valid, issues an access token.

(E)  The client requests the protected resource from the resource
server and authenticates by presenting the access token.

(F)  The resource server validates the access token, and if valid,
serves the request.

The preferred method for the client to obtain an authorization grant
from the resource owner (depicted in steps (A) and (B)) is to use the
authorization server as an intermediary, which is illustrated in
Figure 3 in Section 4.1.


Client Password
The authorization server MUST support the HTTP Basic
authentication scheme for authenticating clients that were issued a
client password.
The client identifier is encoded using the
"application/x-www-form-urlencoded" encoding algorithm per
Appendix B, and the encoded value is used as the username;

Alternatively, the authorization server MAY support including the
client credentials in the request-body using the following
parameters:

client_id
REQUIRED.  The client identifier issued to the client during
the registration process described by Section 2.2.

client_secret
REQUIRED.  The client secret.  The client MAY omit the
parameter if the client secret is an empty string.

Not requested (but nice to have) For example, a request to refresh an access token (Section 6) using
the body parameters (with extra line breaks for display purposes
only):

     POST /token HTTP/1.1
     Host: server.example.com
     Content-Type: application/x-www-form-urlencoded

     grant_type=refresh_token&refresh_token=tGzv3JOkF0XG5Qx2TlKWIA
     &client_id=s6BhdRkqt3&client_secret=7Fjfp0ZBr1KtDRbnfVdmIw

The authorization server MUST require the use of TLS as described in
Section 1.6 when sending requests using password authentication.

Since this client authentication method involves a password, the
authorization server MUST protect any endpoint utilizing it against
brute force attacks.


Protocol Endpoints

The authorization process utilizes two authorization server endpoints
(HTTP resources):

o  Authorization endpoint - used by the client to obtain
authorization from the resource owner via user-agent redirection.

o  Token endpoint - used by the client to exchange an authorization
grant for an access token, typically with client authentication.

As well as one client endpoint:

o  Redirection endpoint - used by the authorization server to return
responses containing authorization credentials to the client via
the resource owner user-agent.

Not every authorization grant type utilizes both endpoints.
Extension grant types MAY define additional endpoints as needed.


Authorization Endpoint

The authorization endpoint is used to interact with the resource
owner and obtain an authorization grant.  The authorization server
MUST first verify the identity of the resource owner.  The way in
which the authorization server authenticates the resource owner
(e.g., username and password login, session cookies) is beyond the
scope of this specification.

The means through which the client obtains the location of the
authorization endpoint are beyond the scope of this specification,
but the location is typically provided in the service documentation.

The endpoint URI MAY include an "application/x-www-form-urlencoded"
formatted (per Appendix B) query component ([RFC3986] Section 3.4),
which MUST be retained when adding additional query parameters.  The
endpoint URI MUST NOT include a fragment component.

Since requests to the authorization endpoint result in user
authentication and the transmission of clear-text credentials (in the
HTTP response), the authorization server MUST require the use of TLS
as described in Section 1.6 when sending requests to the
authorization endpoint.

The authorization server MUST support the use of the HTTP "GET"
method [RFC2616] for the authorization endpoint and MAY support the
use of the "POST" method as well.

Parameters sent without a value MUST be treated as if they were
omitted from the request.  The authorization server MUST ignore
unrecognized request parameters.  Request and response parameters
MUST NOT be included more than once.


Response Type

The authorization endpoint is used by the authorization code grant
type and implicit grant type flows.  The client informs the
authorization server of the desired grant type using the following
parameter:

response_type
REQUIRED.  The value MUST be one of "code" for requesting an
authorization code as described by Section 4.1.1, "token" for
requesting an access token (implicit grant) as described by
Section 4.2.1, or a registered extension value as described by
Section 8.4.


Redirection Endpoint

After completing its interaction with the resource owner, the
authorization server directs the resource owner's user-agent back to
the client.  The authorization server redirects the user-agent to the
client's redirection endpoint previously established with the
authorization server during the client registration process or when
making the authorization request.

The redirection endpoint URI MUST be an absolute URI as defined by
[RFC3986] Section 4.3.  The endpoint URI MAY include an
"application/x-www-form-urlencoded" formatted (per Appendix B) query
component ([RFC3986] Section 3.4), which MUST be retained when adding
additional query parameters.  The endpoint URI MUST NOT include a
fragment component.



Endpoint Request Confidentiality

The redirection endpoint SHOULD require the use of TLS

Registration Requirements

The authorization server MUST require the following clients to
register their redirection endpoint:

o  Public clients.

o  Confidential clients utilizing the implicit grant type.

The authorization server SHOULD require all clients to register their
redirection endpoint prior to utilizing the authorization endpoint.

The authorization server SHOULD require the client to provide the
complete redirection URI (the client MAY use the "state" request
parameter to achieve per-request customization).  If requiring the
registration of the complete redirection URI is not possible, the
authorization server SHOULD require the registration of the URI
scheme, authority, and path (allowing the client to dynamically vary
only the query component of the redirection URI when requesting
authorization).

The authorization server MAY allow the client to register multiple
redirection endpoints.

Lack of a redirection URI registration requirement can enable an
attacker to use the authorization endpoint as an open redirector


Dynamic Configuration

If multiple redirection URIs have been registered, if only part of
the redirection URI has been registered, or if no redirection URI has
been registered, the client MUST include a redirection URI with the
authorization request using the "redirect_uri" request parameter.

When a redirection URI is included in an authorization request, the
authorization server MUST compare and match the value received
against at least one of the registered redirection URIs (or URI
components) as defined in [RFC3986] Section 6, if any redirection
URIs were registered.  If the client registration included the full
redirection URI, the authorization server MUST compare the two URIs
using simple string comparison as defined in [RFC3986] Section 6.2.1.

If an authorization request fails validation due to a missing,
invalid, or mismatching redirection URI, the authorization server
SHOULD inform the resource owner of the error and MUST NOT
automatically redirect the user-agent to the invalid redirection URI.


Endpoint Content

The redirection request to the client's endpoint typically results in
an HTML document response, processed by the user-agent.  If the HTML
response is served directly as the result of the redirection request,
any script included in the HTML document will execute with full
access to the redirection URI and the credentials it contains.

The client SHOULD NOT include any third-party scripts (e.g., third-
party analytics, social plug-ins, ad networks) in the redirection
endpoint response.  Instead, it SHOULD extract the credentials from
the URI and redirect the user-agent again to another endpoint without
exposing the credentials (in the URI or elsewhere).  If third-party
scripts are included, the client MUST ensure that its own scripts
(used to extract and remove the credentials from the URI) will
execute first.


Token Endpoint

The endpoint URI MAY include an "application/x-www-form-urlencoded"
formatted
The authorization server MUST require the use of TLS
The client MUST use the HTTP "POST" method when making access token
requests.
Parameters sent without a value MUST be treated as if they were
omitted from the request.  The authorization server MUST ignore
unrecognized request parameters.  Request and response parameters
MUST NOT be included more than once.


Client Authentication

Confidential clients or other clients issued client credentials MUST
authenticate with the authorization server
A client MAY use the "client_id" request parameter to identify itself
when sending requests to the token endpoint.  In the
"authorization_code" "grant_type" request to the token endpoint, an
unauthenticated client MUST send its "client_id" to prevent itself
from inadvertently accepting a code intended for a client with a
different "client_id".


Access Token Scope

The authorization and token endpoints allow the client to specify the
scope of the access request using the "scope" request parameter.  In
turn, the authorization server uses the "scope" response parameter to
inform the client of the scope of the access token issued.

The value of the scope parameter is expressed as a list of space-
delimited, case-sensitive strings.  The strings are defined by the
authorization server.  If the value contains multiple space-delimited
strings, their order does not matter, and each string adds an
additional access range to the requested scope.

The authorization server MAY fully or partially ignore the scope
requested by the client, based on the authorization server policy or
the resource owner's instructions

If the client omits the scope parameter when requesting
authorization, the authorization server MUST either process the
request using a pre-defined default value or fail the request
indicating an invalid scope.  The authorization server SHOULD
document its scope requirements and default value (if defined).


Obtaining Authorization

To request an access token, the client obtains authorization from the
resource owner.  The authorization is expressed in the form of an
authorization grant, which the client uses to request the access
token.  OAuth defines four grant types: authorization code, implicit,
resource owner password credentials, and client credentials.  It also
provides an extension mechanism for defining additional grant types.



Authorization Code Grant

The authorization code grant type is used to obtain both access
tokens and refresh tokens and is optimized for confidential clients.
Since this is a redirection-based flow, the client must be capable of
interacting with the resource owner's user-agent (typically a web
browser) and capable of receiving incoming requests (via redirection)
from the authorization server.

     +----------+
     | Resource |
     |   Owner  |
     |          |
     +----------+
          ^
          |
         (B)
     +----|-----+          Client Identifier      +---------------+
     |         -+----(A)-- & Redirection URI ---->|               |
     |  User-   |                                 | Authorization |
     |  Agent  -+----(B)-- User authenticates --->|     Server    |
     |          |                                 |               |
     |         -+----(C)-- Authorization Code ---<|               |
     +-|----|---+                                 +---------------+
       |    |                                         ^      v
      (A)  (C)                                        |      |
       |    |                                         |      |
       ^    v                                         |      |
     +---------+                                      |      |
     |         |>---(D)-- Authorization Code ---------'      |
     |  Client |          & Redirection URI                  |
     |         |                                             |
     |         |<---(E)----- Access Token -------------------'
     +---------+       (w/ Optional Refresh Token)

Note: The lines illustrating steps (A), (B), and (C) are broken into
two parts as they pass through the user-agent.

The flow illustrated in Figure 3 includes the following steps:

(A)  The client initiates the flow by directing the resource owner's
user-agent to the authorization endpoint.  The client includes
its client identifier, requested scope, local state, and a
redirection URI to which the authorization server will send the
user-agent back once access is granted (or denied).

(B)  The authorization server authenticates the resource owner (via
the user-agent) and establishes whether the resource owner
grants or denies the client's access request.

(C)  Assuming the resource owner grants access, the authorization
server redirects the user-agent back to the client using the
redirection URI provided earlier (in the request or during
client registration).  The redirection URI includes an
authorization code and any local state provided by the client
earlier.

(D)  The client requests an access token from the authorization
server's token endpoint by including the authorization code
received in the previous step.  When making the request, the
client authenticates with the authorization server.  The client
includes the redirection URI used to obtain the authorization
code for verification.

(E)  The authorization server authenticates the client, validates the
authorization code, and ensures that the redirection URI
received matches the URI used to redirect the client in
step (C).  If valid, the authorization server responds back with
an access token and, optionally, a refresh token.

Authorization Request

The client constructs the request URI by adding the following
parameters to the query component of the authorization endpoint URI
using the "application/x-www-form-urlencoded" format, per Appendix B:

response_type
REQUIRED.  Value MUST be set to "code".

client_id
REQUIRED.  The client identifier The authorization server issues the registered client a client
identifier -- a unique string representing the registration
information provided by the client.  The client identifier is not a
secret;


redirect_uri
OPTIONAL.  The authorization server redirects the user-agent to the
client's redirection endpoint previously established with the
authorization server during the client registration process or when
making the authorization request.


scope
OPTIONAL.  The value of the scope parameter is expressed as a list of space-
delimited, case-sensitive strings.  The strings are defined by the
authorization server.  If the value contains multiple space-delimited
strings, their order does not matter, and each string adds an
additional access range to the requested scope.

state
RECOMMENDED.  An opaque value used by the client to maintain
state between the request and callback.  The authorization
server includes this value when redirecting the user-agent back
to the client.  The parameter SHOULD be used for preventing
cross-site request forgery as described in Section 10.12.

For example, the client directs the user-agent to make the following
HTTP request using TLS (with extra line breaks for display purposes
only):

    GET /authorize?response_type=code&client_id=s6BhdRkqt3&state=xyz
        &redirect_uri=https%3A%2F%2Fclient%2Eexample%2Ecom%2Fcb HTTP/1.1
    Host: server.example.com


Authorization Response

If the resource owner grants the access request, the authorization
server issues an authorization code and delivers it to the client by
adding the following parameters to the query component of the
redirection URI using the "application/x-www-form-urlencoded" format

code
REQUIRED.  The authorization code generated by the
authorization server.  The authorization code MUST expire
shortly after it is issued to mitigate the risk of leaks.  A
maximum authorization code lifetime of 10 minutes is
RECOMMENDED.  The client MUST NOT use the authorization code more than once.  If an authorization code is used more than
once, the authorization server MUST deny the request and SHOULD
revoke (when possible) all tokens previously issued based on
that authorization code.  The authorization code is bound to
the client identifier and redirection URI.

    state
         REQUIRED if the "state" parameter was present in the client
         authorization request.  The exact value received from the
         client.

For example, the authorization server redirects the user-agent by
sending the following HTTP response:

     HTTP/1.1 302 Found
     Location: https://client.example.com/cb?code=SplxlOBeZQQYbYS6WxSbIA
               &state=xyz

The client MUST ignore unrecognized response parameters.  The
authorization code string size is left undefined by this
specification.  The client should avoid making assumptions about code
value sizes.  The authorization server SHOULD document the size of
any value it issues.

Error Response

If the request fails due to a missing, invalid, or mismatching
redirection URI, or if the client identifier is missing or invalid,
the authorization server SHOULD inform the resource owner of the
error and MUST NOT automatically redirect the user-agent to the
invalid redirection URI.

If the resource owner denies the access request or if the request
fails for reasons other than a missing or invalid redirection URI,
the authorization server informs the client by adding the following
parameters to the query component of the redirection URI using the
"application/x-www-form-urlencoded" format, per Appendix B:

error
REQUIRED.  A single ASCII [USASCII] error code from the
following:

         invalid_request
               The request is missing a required parameter, includes an
               invalid parameter value, includes a parameter more than
               once, or is otherwise malformed.

        unauthorized_client
               The client is not authorized to request an authorization
               code using this method.

         access_denied
               The resource owner or authorization server denied the
               request.

         unsupported_response_type
               The authorization server does not support obtaining an
               authorization code using this method.

         invalid_scope
               The requested scope is invalid, unknown, or malformed.

         server_error
               The authorization server encountered an unexpected
               condition that prevented it from fulfilling the request.
               (This error code is needed because a 500 Internal Server
               Error HTTP status code cannot be returned to the client
               via an HTTP redirect.)

         temporarily_unavailable
               The authorization server is currently unable to handle
               the request due to a temporary overloading or maintenance
               of the server.  (This error code is needed because a 503
               Service Unavailable HTTP status code cannot be returned
               to the client via an HTTP redirect.)

error_description
OPTIONAL.  Human-readable ASCII [USASCII] text providing
additional information, used to assist the client developer in
understanding the error that occurred.
Values for the "error_description" parameter MUST NOT include
characters outside the set %x20-21 / %x23-5B / %x5D-7E.

error_uri
OPTIONAL.  A URI identifying a human-readable web page with
information about the error, used to provide the client
developer with additional information about the error.
Values for the "error_uri" parameter MUST conform to the
URI-reference syntax and thus MUST NOT include characters
outside the set %x21 / %x23-5B / %x5D-7E.

state
REQUIRED if a "state" parameter was present in the client
authorization request.  The exact value received from the
client.

For example, the authorization server redirects the user-agent by
sending the following HTTP response:

HTTP/1.1 302 Found
Location: https://client.example.com/cb?error=access_denied&state=xyz


Access Token Request

The client makes a request to the token endpoint by sending the
following parameters using the "application/x-www-form-urlencoded"
format per Appendix B with a character encoding of UTF-8 in the HTTP
request entity-body:

grant_type
REQUIRED.  Value MUST be set to "authorization_code".

code
REQUIRED.  The authorization code received from the
authorization server.

redirect_uri
REQUIRED, if the "redirect_uri" parameter was included in the
authorization request as described in Section 4.1.1, and their
values MUST be identical.

client_id
REQUIRED, if the client is not authenticating with the
authorization server as described in Section 3.2.1.

If the client type is confidential or the client was issued client
credentials (or assigned other authentication requirements), the
client MUST authenticate with the authorization server as described
in Section 3.2.1.

    For example, the client makes the following HTTP request using TLS
(with extra line breaks for display purposes only):

     POST /token HTTP/1.1
     Host: server.example.com
     Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
     Content-Type: application/x-www-form-urlencoded

     grant_type=authorization_code&code=SplxlOBeZQQYbYS6WxSbIA
     &redirect_uri=https%3A%2F%2Fclient%2Eexample%2Ecom%2Fcb

The authorization server MUST:

o  require client authentication for confidential clients or for any
client that was issued client credentials (or with other
authentication requirements),

o  authenticate the client if client authentication is included,

o  ensure that the authorization code was issued to the authenticated
confidential client, or if the client is public, ensure that the
code was issued to "client_id" in the request,

o  verify that the authorization code is valid, and

o  ensure that the "redirect_uri" parameter is present if the
"redirect_uri" parameter was included in the initial authorization
request as described in Section 4.1.1, and if included ensure that
their values are identical.


Access Token Response

    HTTP/1.1 200 OK
     Content-Type: application/json;charset=UTF-8
     Cache-Control: no-store
     Pragma: no-cache

     {
       "access_token":"2YotnFZFEjr1zCsicMWpAA",
       "token_type":"example",
       "expires_in":3600,
       "refresh_token":"tGzv3JOkF0XG5Qx2TlKWIA",
       "example_parameter":"example_value"
     }


Resource Owner Password Credentials Grant


     +----------+
     | Resource |
     |  Owner   |
     |          |
     +----------+
          v
          |    Resource Owner
         (A) Password Credentials
          |
          v
     +---------+                                  +---------------+
     |         |>--(B)---- Resource Owner ------->|               |
     |         |         Password Credentials     | Authorization |
     | Client  |                                  |     Server    |
     |         |<--(C)---- Access Token ---------<|               |
     |         |    (w/ Optional Refresh Token)   |               |
     +---------+                                  +---------------+

            Figure 5: Resource Owner Password Credentials Flow

The flow illustrated in Figure 5 includes the following steps:

(A)  The resource owner provides the client with its username and
password.

(B)  The client requests an access token from the authorization
server's token endpoint by including the credentials received
from the resource owner.  When making the request, the client
authenticates with the authorization server.

(C)  The authorization server authenticates the client and validates
the resource owner credentials, and if valid, issues an access
token.


Authorization Request and Response

The method through which the client obtains the resource owner
credentials is beyond the scope of this specification.  The client
MUST discard the credentials once an access token has been obtained.

The client makes a request to the token endpoint by adding the
following parameters using the "application/x-www-form-urlencoded"
format per Appendix B with a character encoding of UTF-8 in the HTTP
request entity-body:

grant_type
REQUIRED.  Value MUST be set to "password".

username
REQUIRED.  The resource owner username.

password
REQUIRED.  The resource owner password.

scope
OPTIONAL.  The scope of the access request as described by
before

For example, the client makes the following HTTP request using
transport-layer security (with extra line breaks for display purposes
only):

     POST /token HTTP/1.1
     Host: server.example.com
     Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
     Content-Type: application/x-www-form-urlencoded

     grant_type=password&username=johndoe&password=A3ddj3w


The authorization server MUST:

o  require client authentication for confidential clients or for any
client that was issued client credentials (or with other
authentication requirements),

o  authenticate the client if client authentication is included, and

o  validate the resource owner password credentials using its
existing password validation algorithm.


Since this access token request utilizes the resource owner's
password, the authorization server MUST protect the endpoint against
brute force attacks

Access Token Response


An example successful response:

     HTTP/1.1 200 OK
     Content-Type: application/json;charset=UTF-8
     Cache-Control: no-store
     Pragma: no-cache

     {
       "access_token":"2YotnFZFEjr1zCsicMWpAA",
       "token_type":"example",
       "expires_in":3600,
       "refresh_token":"tGzv3JOkF0XG5Qx2TlKWIA",
       "example_parameter":"example_value"
     }


Client Credentials Grant

The client can request an access token using only its client
credentials

The client credentials grant type MUST only be used by confidential
clients.

     +---------+                                  +---------------+
     |         |                                  |               |
     |         |>--(A)- Client Authentication --->| Authorization |
     | Client  |                                  |     Server    |
     |         |<--(B)---- Access Token ---------<|               |
     |         |                                  |               |
     +---------+                                  +---------------+

                     Figure 6: Client Credentials Flow

The flow illustrated in Figure 6 includes the following steps:

(A)  The client authenticates with the authorization server and
requests an access token from the token endpoint.

(B)  The authorization server authenticates the client, and if valid,
issues an access token.

Authorization Request and Response

Since the client authentication is used as the authorization grant,
no additional authorization request is needed.

Access Token Request

The client makes a request to the token endpoint by adding the
following parameters using the "application/x-www-form-urlencoded"
format per Appendix B with a character encoding of UTF-8 in the HTTP
request entity-body:

grant_type
REQUIRED.  Value MUST be set to "client_credentials".

scope
OPTIONAL.  The scope of the access request as described by
Section 3.3.

For example, the client makes the following HTTP request using
transport-layer security (with extra line breaks for display purposes
only):

     POST /token HTTP/1.1
     Host: server.example.com
     Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
     Content-Type: application/x-www-form-urlencoded

     grant_type=client_credentials

The authorization server MUST authenticate the client.

Access Token Response


An example successful response:

     HTTP/1.1 200 OK
     Content-Type: application/json;charset=UTF-8
     Cache-Control: no-store
     Pragma: no-cache

     {
       "access_token":"2YotnFZFEjr1zCsicMWpAA",
       "token_type":"example",
       "expires_in":3600,
       "example_parameter":"example_value"
     }


Extension Grants

The client uses an extension grant type by specifying the grant type
using an absolute URI (defined by the authorization server) as the
value of the "grant_type" parameter of the token endpoint, and by
adding any additional parameters necessary.


For example, to request an access token using a Security Assertion
Markup Language (SAML) 2.0 assertion grant type as defined by
[OAuth-SAML2], the client could make the following HTTP request using
TLS (with extra line breaks for display purposes only):

     POST /token HTTP/1.1
     Host: server.example.com
     Content-Type: application/x-www-form-urlencoded

     grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Asaml2-
     bearer&assertion=PEFzc2VydGlvbiBJc3N1ZUluc3RhbnQ9IjIwMTEtMDU
     [...omitted for brevity...]aG5TdGF0ZW1lbnQ-PC9Bc3NlcnRpb24-

If the access token request is valid and authorized, the
authorization server issues an access token and optional refresh
token as described in Section 5.1.  If the request failed client
authentication or is invalid, the authorization server returns an
error response as described in Section 5.2.

Error Response

The authorization server responds with an HTTP 400 (Bad Request)
status code (unless specified otherwise) and includes the following
parameters with the response:

error
REQUIRED.  A single ASCII [USASCII] error code from the
following:

         invalid_request
               The request is missing a required parameter, includes an
               unsupported parameter value (other than grant type),
               repeats a parameter, includes multiple credentials,
               utilizes more than one mechanism for authenticating the
               client, or is otherwise malformed.

         invalid_client
               Client authentication failed (e.g., unknown client, no
               client authentication included, or unsupported
               authentication method).  The authorization server MAY
               return an HTTP 401 (Unauthorized) status code to indicate
               which HTTP authentication schemes are supported.  If the
               client attempted to authenticate via the "Authorization"
               request header field, the authorization server MUST
               respond with an HTTP 401 (Unauthorized) status code and
               include the "WWW-Authenticate" response header field
               matching the authentication scheme used by the client.

         invalid_grant
               The provided authorization grant (e.g., authorization
               code, resource owner credentials) or refresh token is
               invalid, expired, revoked, does not match the redirection
               URI used in the authorization request, or was issued to
               another client.

         unauthorized_client
               The authenticated client is not authorized to use this
               authorization grant type.

         unsupported_grant_type
               The authorization grant type is not supported by the
               authorization server.

        invalid_scope
               The requested scope is invalid, unknown, malformed, or
               exceeds the scope granted by the resource owner.

         Values for the "error" parameter MUST NOT include characters
         outside the set %x20-21 / %x23-5B / %x5D-7E.

error_description
OPTIONAL.  Human-readable ASCII [USASCII] text providing
additional information, used to assist the client developer in
understanding the error that occurred.
Values for the "error_description" parameter MUST NOT include
characters outside the set %x20-21 / %x23-5B / %x5D-7E.

error_uri
OPTIONAL.  A URI identifying a human-readable web page with
information about the error, used to provide the client
developer with additional information about the error.
Values for the "error_uri" parameter MUST conform to the
URI-reference syntax and thus MUST NOT include characters
outside the set %x21 / %x23-5B / %x5D-7E.

For example:

     HTTP/1.1 400 Bad Request
     Content-Type: application/json;charset=UTF-8
     Cache-Control: no-store
     Pragma: no-cache

     {
       "error":"invalid_request"
     }

Refreshing an Access Token

If the authorization server issued a refresh token to the client, the
client makes a refresh request to the token endpoint by adding the
following parameters using the "application/x-www-form-urlencoded"
format per Appendix B with a character encoding of UTF-8 in the HTTP
request entity-body:

grant_type
REQUIRED.  Value MUST be set to "refresh_token".

refresh_token
REQUIRED.  The refresh token issued to the client.

scope
OPTIONAL.  The scope of the access request as described by
Section 3.3.  The requested scope MUST NOT include any scope
not originally granted by the resource owner, and if omitted is
treated as equal to the scope originally granted by the
resource owner.

Because refresh tokens are typically long-lasting credentials used to
request additional access tokens, the refresh token is bound to the
client to which it was issued.  If the client type is confidential or
the client was issued client credentials (or assigned other
authentication requirements), the client MUST authenticate with the
authorization server as described in Section 3.2.1.

For example, the client makes the following HTTP request using
transport-layer security (with extra line breaks for display purposes
only):

     POST /token HTTP/1.1
     Host: server.example.com
     Authorization: Basic czZCaGRSa3F0MzpnWDFmQmF0M2JW
     Content-Type: application/x-www-form-urlencoded

     grant_type=refresh_token&refresh_token=tGzv3JOkF0XG5Qx2TlKWIA


Access Token Types

The access token type provides the client with the information
required to successfully utilize the access token to make a protected
resource request (along with type-specific attributes).  The client
MUST NOT use an access token if it does not understand the token
type.

For example, the "bearer" token type defined in [RFC6750] is utilized
by simply including the access token string in the request:

     GET /resource/1 HTTP/1.1
     Host: example.com
     Authorization: Bearer mF_9.B5f-4.1JqM

while the "mac" token type defined in [OAuth-HTTP-MAC] is utilized by
issuing a Message Authentication Code (MAC) key together with the
access token that is used to sign certain components of the HTTP
requests:

     GET /resource/1 HTTP/1.1
     Host: example.com
     Authorization: MAC id="h480djs93hd8",
                        nonce="274312:dj83hs9s",
                        mac="kDZvddkndxvhGRXZhvuDjEWhGeE="

The above examples are provided for illustration purposes only.
Developers are advised to consult the [RFC6750] and [OAuth-HTTP-MAC]
specifications before use.

Each access token type definition specifies the additional attributes
(if any) sent to the client together with the "access_token" response
parameter.  It also defines the HTTP authentication method used to
include the access token when making a protected resource request.

Error Response

If a resource access request fails, the resource server SHOULD inform
the client of the error.  While the specifics of such error responses
are beyond the scope of this specification, this document establishes
a common registry in Section 11.4 for error values to be shared among
OAuth token authentication schemes.

New authentication schemes designed primarily for OAuth token
authentication SHOULD define a mechanism for providing an error
status code to the client, in which the error values allowed are
registered in the error registry established by this specification.


Such schemes MAY limit the set of valid error codes to a subset of
the registered values.  If the error code is returned using a named
parameter, the parameter name SHOULD be "error".

Other schemes capable of being used for OAuth token authentication,
but not primarily designed for that purpose, MAY bind their error
values to the registry in the same manner.

New authentication schemes MAY choose to also specify the use of the
"error_description" and "error_uri" parameters to return error
information in a manner parallel to their usage in this
specification.


Pas vu
Get pubkey to test externaly if jwt is correct


https://datatracker.ietf.org/doc/html/rfc6749
https://fastapi.tiangolo.com/tutorial/security/oauth2-jwt/
https://openid.net/specs/openid-connect-core-1_0.html


BlaBla

    Extensibility

8.1.  Defining Access Token Types

Access token types can be defined in one of two ways: registered in
the Access Token Types registry (following the procedures in
Section 11.1), or by using a unique absolute URI as its name.

Types utilizing a URI name SHOULD be limited to vendor-specific
implementations that are not commonly applicable, and are specific to
the implementation details of the resource server where they are
used.

All other types MUST be registered.  Type names MUST conform to the
type-name ABNF.  If the type definition includes a new HTTP
authentication scheme, the type name SHOULD be identical to the HTTP
authentication scheme name (as defined by [RFC2617]).  The token type
"example" is reserved for use in examples.

     type-name  = 1*name-char
     name-char  = "-" / "." / "_" / DIGIT / ALPHA

8.2.  Defining New Endpoint Parameters

New request or response parameters for use with the authorization
endpoint or the token endpoint are defined and registered in the
OAuth Parameters registry following the procedure in Section 11.2.

Parameter names MUST conform to the param-name ABNF, and parameter
values syntax MUST be well-defined (e.g., using ABNF, or a reference
to the syntax of an existing parameter).

     param-name  = 1*name-char
     name-char   = "-" / "." / "_" / DIGIT / ALPHA

Unregistered vendor-specific parameter extensions that are not
commonly applicable and that are specific to the implementation
details of the authorization server where they are used SHOULD
utilize a vendor-specific prefix that is not likely to conflict with
other registered values (e.g., begin with 'companyname_').

8.3.  Defining New Authorization Grant Types

New authorization grant types can be defined by assigning them a
unique absolute URI for use with the "grant_type" parameter.  If the
extension grant type requires additional token endpoint parameters,
they MUST be registered in the OAuth Parameters registry as described
by Section 11.2.

8.4.  Defining New Authorization Endpoint Response Types

New response types for use with the authorization endpoint are
defined and registered in the Authorization Endpoint Response Types
registry following the procedure in Section 11.3.  Response type
names MUST conform to the response-type ABNF.

     response-type  = response-name *( SP response-name )
     response-name  = 1*response-char
     response-char  = "_" / DIGIT / ALPHA

If a response type contains one or more space characters (%x20), it
is compared as a space-delimited list of values in which the order of
values does not matter.  Only one order of values can be registered,
which covers all other arrangements of the same set of values.

For example, the response type "token code" is left undefined by this
specification.  However, an extension can define and register the
"token code" response type.  Once registered, the same combination
cannot be registered as "code token", but both values can be used to
denote the same response type.

8.5.  Defining Additional Error Codes

In cases where protocol extensions (i.e., access token types,
extension parameters, or extension grant types) require additional
error codes to be used with the authorization code grant error
response (Section 4.1.2.1), the implicit grant error response
(Section 4.2.2.1), the token error response (Section 5.2), or the
resource access error response (Section 7.2), such error codes MAY be
defined.


Extension error codes MUST be registered (following the procedures in
Section 11.4) if the extension they are used in conjunction with is a
registered access token type, a registered endpoint parameter, or an
extension grant type.  Error codes used with unregistered extensions
MAY be registered.

Error codes MUST conform to the error ABNF and SHOULD be prefixed by
an identifying name when possible.  For example, an error identifying
an invalid value set to the extension parameter "example" SHOULD be
named "example_invalid".

     error      = 1*error-char
     error-char = %x20-21 / %x23-5B / %x5D-7E



https://datatracker.ietf.org/doc/html/rfc6749
