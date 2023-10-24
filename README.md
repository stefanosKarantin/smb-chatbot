## SMB-CHATBOT

This coding exersice refers to connectly's assessment for the backend engineer position.

The requirements include building a simple chatbot with a predifined flow.

In this source code we have created a simple http client in order to send messages through an `X` messaging platform. On the other hand we expose some endpoints to have a way of receiving messages and notifications from the platform, like having defined webhooks that refer to these endpoints. 

The implementation was based in extensibility, thats why there are explicit layers on storage, application and http levels. 

The domain is consider to be the promotion and stats entities.

The initialization of each flow is made by an http request on the `/start-promotion` endpoint, like creating a new promotion from an internal tool, just the way a sales/operation person would do.
