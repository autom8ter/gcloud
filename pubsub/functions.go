package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
)

type TopicOption func(t *pubsub.Topic)

type SubscriptionOption func(s *pubsub.Subscription)

type MessageOption func(s *pubsub.Message)

type ContextOption func(ctx context.Context)

func NewTopic(name string, client *pubsub.Client, opts ...TopicOption) *pubsub.Topic {
	t := client.Topic(name)
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func WrapTopic(t *pubsub.Topic, opts ...TopicOption) *pubsub.Topic {
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func NewSubscription(name string, client *pubsub.Client, opts ...SubscriptionOption) *pubsub.Subscription {
	s := client.Subscription(name)
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WrapSubscription(t *pubsub.Subscription, opts ...SubscriptionOption) *pubsub.Subscription {
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func NewMessage(opts ...MessageOption) *pubsub.Message {
	m := &pubsub.Message{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func WrapMessage(t *pubsub.Message, opts ...MessageOption) *pubsub.Message {
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func WrapContext(ctx context.Context, opts ...ContextOption) context.Context {
	for _, opt := range opts {
		opt(ctx)
	}
	return ctx
}
