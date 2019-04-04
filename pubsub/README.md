# pubsub
--
    import "github.com/autom8ter/gcloud/pubsub"


## Usage

#### func  NewMessage

```go
func NewMessage(opts ...MessageOption) *pubsub.Message
```

#### func  NewSubscription

```go
func NewSubscription(name string, client *pubsub.Client, opts ...SubscriptionOption) *pubsub.Subscription
```

#### func  NewTopic

```go
func NewTopic(name string, client *pubsub.Client, opts ...TopicOption) *pubsub.Topic
```

#### func  WrapContext

```go
func WrapContext(ctx context.Context, opts ...ContextOption) context.Context
```

#### func  WrapMessage

```go
func WrapMessage(t *pubsub.Message, opts ...MessageOption) *pubsub.Message
```

#### func  WrapSubscription

```go
func WrapSubscription(t *pubsub.Subscription, opts ...SubscriptionOption) *pubsub.Subscription
```

#### func  WrapTopic

```go
func WrapTopic(t *pubsub.Topic, opts ...TopicOption) *pubsub.Topic
```

#### type ContextOption

```go
type ContextOption func(ctx context.Context)
```


#### type MessageOption

```go
type MessageOption func(s *pubsub.Message)
```


#### type PubSub

```go
type PubSub struct {
}
```


#### func  New

```go
func New(ctx context.Context, options ...option.ClientOption) (*PubSub, error)
```

#### func (*PubSub) AddMember

```go
func (g *PubSub) AddMember(ctx context.Context, topicName string, member string, name iam.RoleName) error
```

#### func (*PubSub) Client

```go
func (g *PubSub) Client() *pubsub.Client
```

#### func (*PubSub) Close

```go
func (g *PubSub) Close()
```

#### func (*PubSub) CreateSubscriptionIfNoneExists

```go
func (g *PubSub) CreateSubscriptionIfNoneExists(ctx context.Context, sub string, config pubsub.SubscriptionConfig) (*pubsub.Subscription, error)
```

#### func (*PubSub) DeleteTopic

```go
func (g *PubSub) DeleteTopic(ctx context.Context, topic string) error
```

#### func (*PubSub) GetTopicPolicy

```go
func (g *PubSub) GetTopicPolicy(ctx context.Context, topicName string) (*iam.Policy, error)
```

#### func (*PubSub) ListSubscriptions

```go
func (g *PubSub) ListSubscriptions(ctx context.Context) ([]*pubsub.Subscription, error)
```

#### func (*PubSub) NewMessage

```go
func (g *PubSub) NewMessage(opts ...MessageOption) *pubsub.Message
```

#### func (*PubSub) NewSubscription

```go
func (g *PubSub) NewSubscription(name string, opts ...SubscriptionOption) *pubsub.Subscription
```

#### func (*PubSub) NewTopic

```go
func (g *PubSub) NewTopic(name string, opts ...TopicOption) *pubsub.Topic
```

#### func (*PubSub) PublishMessage

```go
func (g *PubSub) PublishMessage(ctx context.Context, topic string, message *pubsub.Message) error
```

#### func (*PubSub) TestPermissions

```go
func (g *PubSub) TestPermissions(ctx context.Context, topicName string, list []string) ([]string, error)
```

#### func (*PubSub) WrapContext

```go
func (g *PubSub) WrapContext(ctx context.Context, opts ...ContextOption) context.Context
```

#### func (*PubSub) WrapMessage

```go
func (g *PubSub) WrapMessage(msg *pubsub.Message, opts ...MessageOption) *pubsub.Message
```

#### func (*PubSub) WrapSubscription

```go
func (g *PubSub) WrapSubscription(msg *pubsub.Subscription, opts ...SubscriptionOption) *pubsub.Subscription
```

#### func (*PubSub) WrapTopic

```go
func (g *PubSub) WrapTopic(t *pubsub.Topic, opts ...TopicOption) *pubsub.Topic
```

#### type SubscriptionOption

```go
type SubscriptionOption func(s *pubsub.Subscription)
```


#### type TopicOption

```go
type TopicOption func(t *pubsub.Topic)
```
