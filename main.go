package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync/atomic"
	"time"

	pb "chat-examination/grpc"
)

const (
	streamName = "chat"
)

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
	redisClient *redis.Client

	nowStreaming atomic.Int32
}

func (s *ChatServiceServer) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {
	_, err := s.redisClient.XAdd(&redis.XAddArgs{
		Stream: streamName,
		Values: map[string]interface{}{
			"msg":  req.GetMessage(),
			"name": req.GetName(),
		},
	}).Result()
	if err != nil {
		return nil, err
	}

	log.Printf("successfully sent message: %s:%s", req.GetName(), req.GetMessage())

	return &pb.SendResponse{}, nil
}

func (s *ChatServiceServer) Receive(ctx context.Context, req *pb.ReceiveRequest) (*pb.ReceiveResponse, error) {
	streams, err := s.redisClient.XRead(&redis.XReadArgs{
		Streams: []string{streamName, "0"},
	}).Result()
	if err != nil {
		return nil, err
	}

	messages := make([]*pb.Message, 0, len(streams[0].Messages))
	for _, msg := range streams[0].Messages {
		message := &pb.Message{
			Name:    msg.Values["name"].(string),
			Message: msg.Values["msg"].(string),
		}
		messages = append(messages, message)
	}

	return &pb.ReceiveResponse{Messages: messages}, nil
}

func (s *ChatServiceServer) ReceiveStream(req *pb.ReceiveRequest, stream pb.ChatService_ReceiveStreamServer) error {
	uid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	latestMsgID := "0"
	s.nowStreaming.Add(1)
	log.Printf("%s: start chat streaming...", uid.String())
	defer func() {
		s.nowStreaming.Add(-1)
	}()

	// 1秒おきにメッセージを取得してメッセージがあれば送信する
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-stream.Context().Done():
			log.Printf("%s: streaming is closed", uid.String())
			return nil
		case <-t.C:
			streams, err := s.redisClient.XRead(&redis.XReadArgs{
				Streams: []string{streamName, latestMsgID},
				Block:   time.Millisecond,
			}).Result()
			if err != nil {
				if errors.Is(err, redis.Nil) {
					continue
				}
				log.Printf("%s: failed to read stream: %v", uid.String(), err)
				return err
			}
			if len(streams[0].Messages) == 0 {
				log.Printf("%s: no new messages", uid.String())
				continue
			}

			messages := make([]*pb.Message, 0, len(streams[0].Messages))
			for _, msg := range streams[0].Messages {
				message := &pb.Message{
					Name:    msg.Values["name"].(string),
					Message: msg.Values["msg"].(string),
				}
				messages = append(messages, message)
			}
			latestMsgID = streams[0].Messages[len(streams[0].Messages)-1].ID

			if err := stream.Send(&pb.ReceiveResponse{Messages: messages}); err != nil {
				log.Printf("%s: failed to send message: %v", uid.String(), err)
				return err
			}
			log.Printf("%s: sent %d messages", uid.String(), len(messages))
		}
	}
}

func (s *ChatServiceServer) StreamConnectionCount(ctx context.Context, req *pb.StreamConnectionCountRequest) (*pb.StreamConnectionCountResponse, error) {
	return &pb.StreamConnectionCountResponse{Count: s.nowStreaming.Load()}, nil
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 1000,
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &ChatServiceServer{redisClient: rdb})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
