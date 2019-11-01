package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
)

func CreateTopicIfNotExists(ctx context.Context, client *pubsub.Client, topic string) (*pubsub.Topic, error) {
	t := client.Topic(topic)

	ok, err := t.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if ok {
		return t, nil
	}

	t, err = client.CreateTopic(ctx, topic)
	if err != nil {
		return nil, err
	}

	return t, nil
}
