package seabird

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/seabird-chat/seabird-go/pb"
	"google.golang.org/grpc"
)

type ChatIngestClient struct {
	grpcChannel *grpc.ClientConn
	Inner       pb.ChatIngestClient
}

func NewChatIngestClient(seabirdCoreUrl, seabirdCoreToken string) (*ChatIngestClient, error) {
	grpcChannel, err := newGRPCClient(seabirdCoreUrl, seabirdCoreToken)
	if err != nil {
		return nil, err
	}

	return &ChatIngestClient{
		grpcChannel: grpcChannel,
		Inner:       pb.NewChatIngestClient(grpcChannel),
	}, nil
}

func (c *ChatIngestClient) Close() error {
	return c.grpcChannel.Close()
}

func (c *ChatIngestClient) IngestEvents(backendType, backendID string) (*SeabirdChatIngestStream, error) {
	ctx, cancel := context.WithCancel(context.Background())

	inner, err := c.Inner.IngestEvents(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	// Send the proper hello event to get things started
	err = inner.Send(&pb.ChatEvent{
		Inner: &pb.ChatEvent_Hello{
			Hello: &pb.HelloChatEvent{
				BackendInfo: &pb.Backend{
					Type: backendType,
					Id:   backendID,
				},
			},
		},
	})
	if err != nil {
		cancel()
		return nil, err
	}

	// Because our input channel will drop events, we make sure it's buffered in
	// case some come in while we're handling messages.
	in := make(chan *pb.ChatRequest)

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

	return &SeabirdChatIngestStream{
		inner:   inner,
		cancel:  cancel,
		errChan: errChan,
		C:       in,
	}, nil
}

type SeabirdChatIngestStream struct {
	sendLock   sync.Mutex
	inner      pb.ChatIngest_IngestEventsClient
	cancelLock sync.Mutex
	cancel     func()
	errChan    chan error
	C          <-chan *pb.ChatRequest
}

func (s *SeabirdChatIngestStream) Send(event *pb.ChatEvent) error {
	s.sendLock.Lock()
	defer s.sendLock.Unlock()

	return s.inner.Send(event)
}

func (s *SeabirdChatIngestStream) Close() error {
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
