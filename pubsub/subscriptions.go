package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
)

func CreateSubscriptionIfNotExists(
	ctx context.Context,
	client *pubsub.Client,
	subscription string,
	topic *pubsub.Topic,
	config *pubsub.SubscriptionConfig,
) (*pubsub.Subscription, error) {

	sub := client.Subscription(subscription)

	ok, err := sub.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if ok {
		return sub, nil
	}

	if config == nil {
		config = &pubsub.SubscriptionConfig{}
	}

	config.Topic = topic

	sub, err = client.CreateSubscription(ctx, subscription, *config)
	if err != nil {
		return nil, err
	}

	return sub, nil
}
