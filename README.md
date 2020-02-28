# grpc-sample-voice-over

## Start servers

```bash
# bash
export SLACK_VERIFICATION_TOKEN=[TOKEN]; export GRPC_PASSWORD=[PASSWORD]; go run server/*.go
# fisn
set -x SLACK_VERIFICATION_TOKEN [TOKEN] & set -x GRPC_PASSWORD=[PASSWORD] & go run server/*.go
```

## Run test client

```bash
# bash
export GRPC_PASSWORD=[PASSWORD]; go run client/*.go
# fisn
set -x GRPC_PASSWORD [PASSWORD] & go run client/*.go

```


## For Swift protoc

at root directory of [grpc-swift](https://github.com/grpc/grpc-swift)

```
sudo xcode-select -s /Applications/Xcode.app/Contents/Developer
make plugins
sudo cp protoc-gen-swift protoc-gen-grpc-swift /usr/local/bin
```

reference: https://github.com/grpc/grpc-swift/tree/nio#getting-the-protoc-plugins

## For Slack Event Subscriptions

1. Install ngrok application
2. Run ngrok with port specified `ngrok http 3000`
3. Register [Slack Event Subscriptions](https://api.slack.com/apps/***/event-subscriptions) and retrieve `Verification Token`
