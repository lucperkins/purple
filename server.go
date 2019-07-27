package strato

import (
	"context"
	"fmt"
	"net"

	"github.com/lucperkins/strato/proto"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	address string
	srv     *grpc.Server
	mem     *Memory
	log     *logrus.Entry
}

var (
	_ proto.CacheServer   = (*GrpcServer)(nil)
	_ proto.CounterServer = (*GrpcServer)(nil)
	_ proto.KVServer      = (*GrpcServer)(nil)
	_ proto.SearchServer  = (*GrpcServer)(nil)
)

func NewServer(cfg *ServerConfig) (*GrpcServer, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	addr := fmt.Sprintf(":%d", cfg.Port)

	srv := grpc.NewServer()

	mem := NewMemory()

	log := logrus.New().WithField("process", "server")

	return &GrpcServer{
		address: addr,
		srv:     srv,
		mem:     mem,
		log:     log,
	}, nil
}

func (s *GrpcServer) CacheGet(_ context.Context, req *proto.CacheGetRequest) (*proto.CacheGetResponse, error) {
	val, err := s.mem.CacheGet(req.Key)
	if err != nil {
		return nil, err
	}

	res := &proto.CacheGetResponse{
		Value: val,
	}

	return res, nil
}

func (s *GrpcServer) CacheSet(_ context.Context, req *proto.CacheSetRequest) (*proto.Empty, error) {
	key, val, ttl := req.Key, req.Item.Value, req.Item.Ttl

	if err := s.mem.CacheSet(key, val, ttl); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (s *GrpcServer) IncrementCounter(_ context.Context, req *proto.IncrementCounterRequest) (*proto.Empty, error) {
	s.mem.IncrementCounter(req.Key, req.Amount)

	return &proto.Empty{}, nil
}

func (s *GrpcServer) GetCounter(_ context.Context, req *proto.GetCounterRequest) (*proto.GetCounterResponse, error) {
	val := s.mem.GetCounter(req.Key)

	return &proto.GetCounterResponse{
		Value: val,
	}, nil
}

func (s *GrpcServer) Get(_ context.Context, location *proto.Location) (*proto.GetResponse, error) {
	loc := &Location{
		Key: location.Key,
	}

	val, err := s.mem.KVGet(loc)
	if err != nil {
		return nil, NotFound(loc).AsProtoStatus()
	}

	res := &proto.GetResponse{
		Value: &proto.Value{
			Content: val.Content,
		},
	}

	return res, nil
}

func (s *GrpcServer) Put(_ context.Context, req *proto.PutRequest) (*proto.Empty, error) {
	loc := &Location{
		Key: req.Location.Key,
	}

	val := &Value{
		Content: req.Value.Content,
	}

	s.mem.KVPut(loc, val)

	return &proto.Empty{}, nil
}

func (s *GrpcServer) Delete(_ context.Context, location *proto.Location) (*proto.Empty, error) {
	loc := &Location{
		Key: location.Key,
	}

	s.mem.KVDelete(loc)

	return &proto.Empty{}, nil
}

func (s *GrpcServer) Index(_ context.Context, req *proto.IndexRequest) (*proto.Empty, error) {
	doc := docFromProto(req.Document)

	s.mem.Index(doc)

	return &proto.Empty{}, nil
}

func (s *GrpcServer) Query(_ context.Context, query *proto.SearchQuery) (*proto.SearchResults, error) {
	q := query.Query

	docs := s.mem.Query(q)

	return docsToResults(docs), nil
}

func (s *GrpcServer) Start() error {
	proto.RegisterCacheServer(s.srv, s)

	s.log.Debug("registered gRPC cache service")

	proto.RegisterCounterServer(s.srv, s)

	s.log.Debug("registered gRPC counter service")

	proto.RegisterKVServer(s.srv, s)

	s.log.Debug("registered gRPC KV service")

	proto.RegisterSearchServer(s.srv, s)

	s.log.Debug("registered gRPC search service")

	lis, _ := net.Listen("tcp", s.address)

	s.log.Debugf("starting TCP listener on %s", s.address)

	return s.srv.Serve(lis)
}

func (s *GrpcServer) ShutDown() {
	s.log.Debug("shutting down")

	s.srv.GracefulStop()
}
