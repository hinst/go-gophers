# Web retry update

Looking at file `webRetry.go`.
The current approach is working poorly, because we try to reuse `request` object.
We are getting silly errors, such as:
> panic: Post "http://prettier-server.default.svc:3000?filename=index.html": EOF

Proposed change: instead of accepting `request` object directly as a parameter,
require function returning a request as parameter.
* Before: Run(client *http.Client, request *http.Request)
* After: Run(client *http.Client, requestFactory RequestFactory)
    * restoreBody() will be no longer needed with the new approach

Evaluate the proposed approach. If it looks OK, then proceed with the change.
