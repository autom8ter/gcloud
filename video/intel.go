package video

import (
	video "cloud.google.com/go/videointelligence/apiv1"
	"context"
	"google.golang.org/api/option"
	videopb "google.golang.org/genproto/googleapis/cloud/videointelligence/v1"
	"io"
	"io/ioutil"
)

type Intelligence struct {
	vid *video.Client
}

func NewIntelligence(ctx context.Context, opts ...option.ClientOption) (*Intelligence, error) {
	client, err := video.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &Intelligence{
		vid: client,
	}, nil
}

func (v *Intelligence) Close() {
	_ = v.vid.Close()
}

func (v *Intelligence) Client() *video.Client {
	return v.vid
}

// objectTracking analyzes a video and extracts entities with their bounding boxes.
func (i *Intelligence) TrackObjects(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error) {
	bits, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	op, err := i.vid.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputContent: bits,
		Features: []videopb.Feature{
			videopb.Feature_OBJECT_TRACKING,
		},
	})
	if err != nil {
		return nil, err
	}

	return op.Wait(ctx)
}

func (i *Intelligence) TrackObjectsFromStorage(ctx context.Context, gcsURI string, w io.Writer) (*videopb.AnnotateVideoResponse, error) {
	op, err := i.vid.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputUri: gcsURI,
		Features: []videopb.Feature{
			videopb.Feature_OBJECT_TRACKING,
		},
	})
	if err != nil {
		return nil, err
	}

	return op.Wait(ctx)
}

func (i *Intelligence) DetectText(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error) {
	fileBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	op, err := i.vid.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputContent: fileBytes,
		Features: []videopb.Feature{
			videopb.Feature_TEXT_DETECTION,
		},
	})
	if err != nil {
		return nil, err
	}

	return op.Wait(ctx)
}

func (i *Intelligence) DetectFaces(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error) {
	fileBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	op, err := i.vid.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputContent: fileBytes,
		Features: []videopb.Feature{
			videopb.Feature_FACE_DETECTION,
		},
	})
	if err != nil {
		return nil, err
	}

	return op.Wait(ctx)
}

func (i *Intelligence) DetectExplicitContent(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error) {
	fileBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	op, err := i.vid.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputContent: fileBytes,
		Features: []videopb.Feature{
			videopb.Feature_EXPLICIT_CONTENT_DETECTION,
		},
	})
	if err != nil {
		return nil, err
	}

	return op.Wait(ctx)
}

func (i *Intelligence) TranscribeSpeech(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error) {
	fileBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	op, err := i.vid.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputContent: fileBytes,
		Features: []videopb.Feature{
			videopb.Feature_SPEECH_TRANSCRIPTION,
		},
	})
	if err != nil {
		return nil, err
	}

	return op.Wait(ctx)
}

func (i *Intelligence) DetectLabel(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error) {
	fileBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	op, err := i.vid.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputContent: fileBytes,
		Features: []videopb.Feature{
			videopb.Feature_LABEL_DETECTION,
		},
	})
	if err != nil {
		return nil, err
	}

	return op.Wait(ctx)
}

func (i *Intelligence) DetectAll(ctx context.Context, r io.Reader, w io.Writer) (*videopb.AnnotateVideoResponse, error) {
	fileBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	op, err := i.vid.AnnotateVideo(ctx, &videopb.AnnotateVideoRequest{
		InputContent: fileBytes,
		Features: []videopb.Feature{
			videopb.Feature_LABEL_DETECTION,
			videopb.Feature_SPEECH_TRANSCRIPTION,
			videopb.Feature_TEXT_DETECTION,
			videopb.Feature_SHOT_CHANGE_DETECTION,
			videopb.Feature_EXPLICIT_CONTENT_DETECTION,
			videopb.Feature_FACE_DETECTION,
			videopb.Feature_OBJECT_TRACKING,
		},
	})
	if err != nil {
		return nil, err
	}

	return op.Wait(ctx)
}
