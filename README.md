# jwt2header

JSON Web Tokens aka JWT are an open, industry standard method for representing claims securely between two parties. This plugin can be used to simply convert JWT claims to headers that can be consumed by the upstream service or leveraged by other components of Gateway to manage traffic.

## Problem Statement

APISIX plays a critical role in managing the traffic to backend services and ensuring that requests are handled securely, efficiently, and reliably. APISIX the ability to manipulate and inspect HTTP headers, which provide important metadata about the incoming requests. Headers can be leveraged for: 
- Traffic routing: Additional headers can be used to route traffic to different backend services based on the value of the header. For example, a header might be used to indicate the version of an API that the client is using, which could be used to route the request to the appropriate backend service.
- Load balancing: Additional headers can be used to provide information about the current load on backend servers, which can be used by the gateway to route requests to the least busy server.
- Request and response modification: Additional headers can be used to modify the request or response before it is forwarded to the backend service or returned to the client. For example, headers might be added or removed, or the value of a header might be modified.


With JWT, claims are the payload which may contain the data to be used to fulfill the above  mentioned usecases. But the JWT claims are generally verified by the upstream service only usually which is a missed opportunity.

## Proposed Solution

Consider a scenario where we want to create a special release for beta users. This special release is a different upstream service.



In the context of API Gateway, the traffic split plugin is often used to distribute incoming traffic across different services or versions of a service based on specific criteria, such as header values. In this scenario, the goal is to create a special release for beta users by directing their traffic to a different upstream service

![](https://i.imgur.com/EYnUt44.png)


To achieve this, the beta user clients attach a special header, "Special_User": "True," to their requests. However, there is a risk that normal users may mimic this behavior by adding the same header to their requests, which could result in incorrect routing of traffic.

![](https://i.imgur.com/roq25Yv.png)


To address this issue, the solution proposes using JWT (JSON Web Token) claims to verify and authorize beta users. JWT is a compact, URL-safe means of representing claims to be transferred between two parties and is commonly used for authentication and authorization purposes.

However, since JWT claims are encoded, a method is needed to decode them while verifying their authenticity. To accomplish this, the solution proposes using a plugin called jwt2header, which can convert JWT claims to HTTP headers that can be consumed by the upstream service or other components of the Gateway.

Once the JWT claims are converted to HTTP headers, the traffic split plugin can leverage these headers to navigate the request and route the traffic to the appropriate upstream service based on the user's beta status. This solution provides a more secure and reliable method for managing beta users' traffic without the risk of unauthorized access or incorrect routing.

![](https://i.imgur.com/ehZ6mdH.png)

## Implementation Details

This repo contains a sample implementation of "jwt2header". 
It leverages Apache APISIX External Plugin Runners.
We hook onto `ext-plugin-pre-req` where we have setup `RequestFilter` to decode JWT Claims and convert them to HTTP Headers.

Please note that this plugin does NOT validate JWT tokens. This plugin is meant to be used in conjuction with `jwt-auth` which will actually validate the JWT Signature.

## TODO

- Implement `ext-plugin-post-req` which validates whether `jwt-auth` was able to verify the signature and strip or reject headers if needed.
- Make JWT location configurable, currently it expects the JWT to be part of Authorisation Header but JWT can be part of cookies and query parameters too. This implementation will be similar to how things are as part of `jwt-auth`
