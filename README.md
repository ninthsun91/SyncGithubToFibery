# User Guide

This is an AWS lambda function, which receives Github Webhook and sends Fibery API to sync the specified repository.

## AWS Lambda

### Runtime
Since this is built in Go and I use Apple chip MacOS, I created lambda function with `Amazon Linux 2023` runtime and architecture to `arm64`. Also, change the handler name to `main`.

![alt text](/public/aws-lambda-runtime.png)

However, if you use different environment such as Windows or Ubuntu, then feel free to change the architecture. In this case, you'll also have to update `Makerfile`'s `GOARCH` value to your corresponding architecture type.

### Environment Variables
Set following enviroment variables:
- `FIBERY_API_TOKEN`
- `FIBERY_SPACE_ID`
- `FIBERY_SPACE_NAME`
- `FIBERY_SYNC_SOURCE_ID`
- `FIBERY_WORKSPACE_NAME`
- `GITHUB_WEBHOOK_SECRET`

For those who don't know where to get these values, follow the guide below.

## Github Webhook

### Create Webhook
When creating Github Webhook, make sure to check the **Content Type** as `application/json`.
![alt text](/public/github-webhook.png)

The **Secret** used here is the value for `GITHUB_WEBHOOK_SECRET` in environment variables.

Feel free to customize webhook trigger events. Normally, push or pull request trigger is enough.

### Validation
Read [Validating webhook deliveries](https://docs.github.com/en/webhooks/using-webhooks/validating-webhook-deliveries) for detail.

## Fibery

### API Token
Go to **Settings** and create your own API key. That's your `FIBERY_API_TOKEN`.

### Github Integration
Follow Fibery's [Github Integration guide](https://the.fibery.io/@public/search/93kmG#User_Guide/Guide/GitHub-Integration-71) for basic integration.

After, you finished integration, open Github Integration setting in web browser. Open developer tool and go to network tab. 

Press **Sync now** button from Fibery, and find **sync** request . From the **Request URL** you'll get your sync source id. (`https://<FIBERY_WORKSPACE_NAME>.fibery.io/api/data-sync/sync-sources/<FIBERY_SYNC_SOURCE_ID>/sync`)


Then, open payload tab. You'll get your space name and space id. (`{app: {name: <FIBERY_SPACE_NAME>, id: <FIBERY_SPACE_ID>}, options: {}}`).