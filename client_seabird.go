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

// Client is a convenience wrapper around an inner gRPC SeabirdClient.
//
// It handles authentication, HTTP/HTTPS connection boilerplate, and provides a wrapper around the StreamEvents method to make it easier to work with.
type Client struct {
	grpcChannel *grpc.ClientConn
	Inner       pb.SeabirdClient
}

func NewClient(seabirdCoreUrl, seabirdCoreToken string) (*Client, error) {
	grpcChannel, err := newGRPCClient(seabirdCoreUrl, seabirdCoreToken)
	if err != nil {
		return nil, err
	}

	return &Client{
		grpcChannel: grpcChannel,
		Inner:       pb.NewSeabirdClient(grpcChannel),
	}, nil
}

func (c *Client) Close() error {
	return c.grpcChannel.Close()
}

func (c *Client) Reply(source *pb.ChannelSource, msg string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.Inner.SendMessage(ctx, &pb.SendMessageRequest{
		ChannelId: source.GetChannelId(),
		Text:      msg,
	})

	return err
}

func (c *Client) Replyf(source *pb.ChannelSource, format string, args ...interface{}) error {
	return c.Reply(source, fmt.Sprintf(format, args...))
}

func (c *Client) MentionReply(source *pb.ChannelSource, msg string) error {
	return c.Reply(source, fmt.Sprintf("%s: %s", source.GetUser().GetDisplayName(), msg))
}

func (c *Client) MentionReplyf(source *pb.ChannelSource, format string, args ...interface{}) error {
	return c.MentionReply(source, fmt.Sprintf(format, args...))
}

func (c *Client) StreamEvents(cmds map[string]*pb.CommandMetadata) (*EventStream, error) {
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

	es := &EventStream{
		inner:      inner,
		cancelFunc: cancel,
		errChan:    errChan,
		C:          in,
	}

	go func() {
		done := ctx.Done()

		for {
			event, err := inner.Recv()
			if err != nil {
				select {
				case errChan <- err:
				default:
				}

				close(in)

				return
			}

			select {
			case in <- event:
			case <-done:
			}
		}
	}()

	return es, nil
}

type EventStream struct {
	inner      pb.Seabird_StreamEventsClient
	cancelFunc func()
	cancelOnce sync.Once
	errChan    chan error
	C          <-chan *pb.Event
}

func (s *EventStream) Close() error {
	s.cancelOnce.Do(s.cancelFunc)

	select {
	case err := <-s.errChan:
		return err
	case <-time.After(1 * time.Second):
		return errors.New("timeout when closing")
	}
}
