package redis

import (
	"context"
)

const Ws = "websocket"

// Publish 发布消息
func Publish(ctx context.Context, channel string, message interface{}) error {
	return RDB.Publish(ctx, channel, message).Err()
}

// Subscribe 订阅消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	msg, err := RDB.Subscribe(ctx, channel).ReceiveMessage(ctx)
	return msg.Payload, err
}
