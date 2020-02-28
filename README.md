# grpc-sample-voice-over

## Start servers

```fish
set -x SLACK_VERIFICATION_TOKEN [TOKEN] & go run server/*.go
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

1. install ngrok application
2. run ngrok with port specified `ngrok http 3000`
3. register Slack Event Subscriptions at `https://api.slack.com/apps/***/event-subscriptions?`
