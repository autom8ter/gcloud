# lang
--
    import "github.com/autom8ter/gcloud/lang"


## Usage

#### type Lang

```go
type Lang struct {
}
```


#### func  New

```go
func New(ctx context.Context, opts ...option.ClientOption) (*Lang, error)
```

#### func (*Lang) Close

```go
func (l *Lang) Close()
```

#### func (*Lang) Speech

```go
func (l *Lang) Speech() *Speech
```

#### func (*Lang) Text2Speech

```go
func (l *Lang) Text2Speech() *Text2Speech
```

#### func (*Lang) Translator

```go
func (l *Lang) Translator() *Translator
```

#### type RecognizeOption

```go
type RecognizeOption func(r *speechpb.RecognizeRequest)
```


#### type Speech

```go
type Speech struct {
}
```


#### func  NewSpeech

```go
func NewSpeech(ctx context.Context, opts ...option.ClientOption) (*Speech, error)
```

#### func (*Speech) Client

```go
func (s *Speech) Client() *speech.Client
```

#### func (*Speech) Close

```go
func (s *Speech) Close()
```

#### func (*Speech) Recognize

```go
func (s *Speech) Recognize(ctx context.Context, opts ...RecognizeOption) (*speechpb.RecognizeResponse, error)
```

#### type Text2Speech

```go
type Text2Speech struct {
}
```


#### func  NewText2Speech

```go
func NewText2Speech(ctx context.Context, opts ...option.ClientOption) (*Text2Speech, error)
```

#### func (*Text2Speech) Close

```go
func (l *Text2Speech) Close()
```

#### func (*Text2Speech) ListVoices

```go
func (l *Text2Speech) ListVoices(ctx context.Context) ([]*texttospeechpb.Voice, error)
```
ListVoices lists the available text to speech voices.

#### func (*Text2Speech) ParseLanguage

```go
func (l *Text2Speech) ParseLanguage(targetLanguage string) (language.Tag, error)
```

#### func (*Text2Speech) SynthesizeSpeech

```go
func (l *Text2Speech) SynthesizeSpeech(ctx context.Context, dest io.Writer, opts ...Text2SpeechOption) error
```

#### type Text2SpeechOption

```go
type Text2SpeechOption func(r *texttospeechpb.SynthesizeSpeechRequest)
```


#### type Translation

```go
type Translation struct {
	Text       []string
	TargetLang language.Tag
	Opts       *translate.Options
}
```


#### type TranslationOption

```go
type TranslationOption func(t *Translation)
```


#### type Translator

```go
type Translator struct {
}
```


#### func  NewTranslator

```go
func NewTranslator(ctx context.Context, opts ...option.ClientOption) (*Translator, error)
```

#### func (*Translator) Client

```go
func (t *Translator) Client() *translate.Client
```

#### func (*Translator) Close

```go
func (l *Translator) Close()
```

#### func (*Translator) DetectLanguage

```go
func (l *Translator) DetectLanguage(ctx context.Context, text string) (*translate.Detection, error)
```

#### func (*Translator) Translate

```go
func (t *Translator) Translate(ctx context.Context, opts ...TranslationOption) ([]translate.Translation, error)
```
