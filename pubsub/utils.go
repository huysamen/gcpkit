package pubsub

import (
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/pubsub"
)

func SubscribeToTopic(
	ctx context.Context,
	client *pubsub.Client,
	topic string,
	subscription string,
	deadline time.Duration,
	sync bool,
	maxOutstandingMessages int,
	fx func(ctx context.Context, msg *pubsub.Message),
) error {
	t, err := CreateTopicIfNotExists(ctx, client, topic)
	if err != nil {
		return err
	}

	config := &pubsub.SubscriptionConfig{
		AckDeadline: deadline,
	}

	s, err := CreateSubscriptionIfNotExists(ctx, client, subscription, t, config)
	if err != nil {
		return err
	}

	s.ReceiveSettings.Synchronous = sync
	s.ReceiveSettings.MaxOutstandingMessages = maxOutstandingMessages

	go func() {
		_ = s.Receive(ctx, fx)
	}()

	return nil
}

func Publish(ctx context.Context, client *pubsub.Client, topic string, message interface{}) (string, error) {
	data, err := json.Marshal(message)
	if err != nil {
		return "", err
	}

	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{Data: data})

	return result.Get(ctx)
}
