package pubsub

import (
	"cloud.google.com/go/iam"
	"cloud.google.com/go/pubsub"
	"context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"os"
)

type PubSub struct {
	cli *pubsub.Client
}

func New(ctx context.Context, options ...option.ClientOption) (*PubSub, error) {
	s, err := pubsub.NewClient(ctx, os.Getenv("GCLOUD_PROJECTID"), options...)
	if err != nil {
		return nil, err
	}
	return &PubSub{
		cli: s,
	}, nil

}

func (g *PubSub) Client() *pubsub.Client {
	return g.cli
}

func (g *PubSub) Close() {
	_ = g.cli.Close()
}

func (g *PubSub) NewMessage(opts ...MessageOption) *pubsub.Message {
	return NewMessage(opts...)
}

func (g *PubSub) WrapMessage(msg *pubsub.Message, opts ...MessageOption) *pubsub.Message {
	return WrapMessage(msg, opts...)
}

func (g *PubSub) NewSubscription(name string, opts ...SubscriptionOption) *pubsub.Subscription {
	return NewSubscription(name, g.cli, opts...)
}

func (g *PubSub) WrapSubscription(msg *pubsub.Subscription, opts ...SubscriptionOption) *pubsub.Subscription {
	return WrapSubscription(msg, opts...)
}

func (g *PubSub) NewTopic(name string, opts ...TopicOption) *pubsub.Topic {
	return NewTopic(name, g.cli, opts...)
}

func (g *PubSub) WrapTopic(t *pubsub.Topic, opts ...TopicOption) *pubsub.Topic {
	return WrapTopic(t, opts...)
}

func (g *PubSub) WrapContext(ctx context.Context, opts ...ContextOption) context.Context {
	return WrapContext(ctx, opts...)
}

func (g *PubSub) CreateSubscriptionIfNoneExists(ctx context.Context, sub string, config pubsub.SubscriptionConfig) (*pubsub.Subscription, error) {
	// Create a topic to subscribe to.
	s := g.cli.Subscription(sub)
	ok, err := s.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if ok {
		return s, nil
	}

	s, err = g.cli.CreateSubscription(ctx, sub, config)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (g *PubSub) DeleteTopic(ctx context.Context, topic string) error {
	t := g.cli.Topic(topic)
	if err := t.Delete(ctx); err != nil {
		return err
	}
	return nil
}

func (g *PubSub) PublishMessage(ctx context.Context, topic string, message *pubsub.Message) error {
	t := g.cli.Topic(topic)

	result := t.Publish(ctx, message)
	id, err := result.Get(ctx)
	if err != nil {
		return err
	}
	log.Printf("Published message; ID: %v\n", id)
	return nil
}

func (g *PubSub) GetTopicPolicy(ctx context.Context, topicName string) (*iam.Policy, error) {
	policy, err := g.cli.Topic(topicName).IAM().Policy(ctx)
	if err != nil {
		return nil, err
	}
	for _, role := range policy.Roles() {
		log.Print(policy.Members(role))
	}
	return policy, nil
}

func (g *PubSub) TestPermissions(ctx context.Context, topicName string, list []string) ([]string, error) {
	topic := g.cli.Topic(topicName)
	perms, err := topic.IAM().TestPermissions(ctx, list)
	if err != nil {
		return nil, err
	}
	for _, perm := range perms {
		log.Printf("Allowed: %v", perm)
	}
	return perms, nil
}

func (g *PubSub) ListSubscriptions(ctx context.Context) ([]*pubsub.Subscription, error) {
	var subs []*pubsub.Subscription
	it := g.cli.Subscriptions(ctx)
	for {
		s, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}

func (g *PubSub) AddMember(ctx context.Context, topicName string, member string, name iam.RoleName) error {
	topic := g.cli.Topic(topicName)
	policy, err := topic.IAM().Policy(ctx)
	if err != nil {
		return err
	}
	policy.Add(member, name)
	if err := topic.IAM().SetPolicy(ctx, policy); err != nil {
		log.Fatalf("SetPolicy: %v", err)
	}
	return nil
}
