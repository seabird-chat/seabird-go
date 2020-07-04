package seabird

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/seabird-chat/seabird-go/pb"
)

type SeabirdClient struct {
	grpcChannel *grpc.ClientConn
	Inner       pb.SeabirdClient
}

func NewSeabirdClient(seabirdCoreUrl, seabirdCoreToken string) (*SeabirdClient, error) {
	grpcChannel, err := newGRPCClient(seabirdCoreUrl, seabirdCoreToken)
	if err != nil {
		return nil, err
	}

	return &SeabirdClient{
		grpcChannel: grpcChannel,
		Inner:       pb.NewSeabirdClient(grpcChannel),
	}, nil
}

func (c *SeabirdClient) Close() error {
	return c.grpcChannel.Close()
}

func (c *SeabirdClient) Reply(source *pb.ChannelSource, msg string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.Inner.SendMessage(ctx, &pb.SendMessageRequest{
		ChannelId: source.GetChannelId(),
		Text:      msg,
	})

	return err
}

func (c *SeabirdClient) Replyf(source *pb.ChannelSource, format string, args ...interface{}) error {
	return c.Reply(source, fmt.Sprintf(format, args...))
}

func (c *SeabirdClient) MentionReply(source *pb.ChannelSource, msg string) error {
	return c.Reply(source, fmt.Sprintf("%s: %s", source.GetUser().GetDisplayName(), msg))
}

func (c *SeabirdClient) MentionReplyf(source *pb.ChannelSource, format string, args ...interface{}) error {
	return c.MentionReply(source, fmt.Sprintf(format, args...))
}

func (c *SeabirdClient) StreamEvents(cmds map[string]*pb.CommandMetadata) (*SeabirdEventStream, error) {
	ctx, cancel := context.WithCancel(context.Background())

	inner, err := c.Inner.StreamEvents(ctx, &pb.StreamEventsRequest{Commands: cmds})
	if err != nil {
		cancel()
		return nil, err
	}

	// Because our input channel will drop events, we make sure it's buffered in
	// case some come in while we're handling messages.
	in := make(chan *pb.Event, 10)

	// Make sure the channel is buffered so we can store a value.
	errChan := make(chan error, 1)

	go func() {
		done := ctx.Done()

		for {
			event, err := inner.Recv()
			if err != nil {
				select {
				case errChan <- err:
				default:
				}

				return
			}

			select {
			case in <- event:
			case <-done:
			}
		}
	}()

	return &SeabirdEventStream{
		inner:   inner,
		cancel:  cancel,
		errChan: errChan,
		C:       in,
	}, nil
}

type SeabirdEventStream struct {
	inner      pb.Seabird_StreamEventsClient
	cancelLock sync.Mutex
	cancel     func()
	errChan    chan error
	C          <-chan *pb.Event
}

func (s *SeabirdEventStream) Close() error {
	s.cancelLock.Lock()
	defer s.cancelLock.Unlock()

	if s.cancel != nil {
		s.cancel()
		s.cancel = nil
	}

	select {
	case err := <-s.errChan:
		return err
	case <-time.After(1 * time.Second):
		return errors.New("timeout when closing")
	}
}
