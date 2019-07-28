package strato

import (
	"context"
	"fmt"
	"net"

	"github.com/lucperkins/strato/proto"

	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type Server struct {
	address string
	srv     *grpc.Server
	mem     *Memory
	log     *logrus.Entry
}

var (
	_ proto.CacheServer   = (*Server)(nil)
	_ proto.CounterServer = (*Server)(nil)
	_ proto.KVServer      = (*Server)(nil)
	_ proto.SearchServer  = (*Server)(nil)
	_ proto.SetServer     = (*Server)(nil)
)

func NewServer(cfg *ServerConfig) (*Server, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	addr := fmt.Sprintf(":%d", cfg.Port)

	srv := grpc.NewServer()

	mem := NewMemoryBackend()

	log := logrus.New().WithField("process", "server")

	return &Server{
		address: addr,
		srv:     srv,
		mem:     mem,
		log:     log,
	}, nil
}

func (s *Server) CacheGet(_ context.Context, req *proto.CacheGetRequest) (*proto.CacheGetResponse, error) {
	val, err := s.mem.CacheGet(req.Key)
	if err != nil {
		return nil, err
	}

	res := &proto.CacheGetResponse{
		Value: val,
	}

	return res, nil
}

func (s *Server) CacheSet(_ context.Context, req *proto.CacheSetRequest) (*proto.Empty, error) {
	key, val, ttl := req.Key, req.Item.Value, req.Item.Ttl

	if err := s.mem.CacheSet(key, val, ttl); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (s *Server) IncrementCounter(_ context.Context, req *proto.IncrementCounterRequest) (*proto.Empty, error) {
	s.mem.CounterIncrement(req.Key, req.Amount)

	return &proto.Empty{}, nil
}

func (s *Server) GetCounter(_ context.Context, req *proto.GetCounterRequest) (*proto.GetCounterResponse, error) {
	val := s.mem.CounterGet(req.Key)

	return &proto.GetCounterResponse{
		Value: val,
	}, nil
}

func (s *Server) KVGet(_ context.Context, location *proto.Location) (*proto.GetResponse, error) {
	loc := &Location{
		Bucket: location.Bucket,
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

func (s *Server) KVPut(_ context.Context, req *proto.PutRequest) (*proto.Empty, error) {
	loc := &Location{
		Bucket: req.Location.Bucket,
		Key: req.Location.Key,
	}

	val := &Value{
		Content: req.Value.Content,
	}

	if err := s.mem.KVPut(loc, val); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (s *Server) KVDelete(_ context.Context, location *proto.Location) (*proto.Empty, error) {
	loc := &Location{
		Bucket: location.Bucket,
		Key: location.Key,
	}

	if err := s.mem.KVDelete(loc); err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (s *Server) Index(_ context.Context, req *proto.IndexRequest) (*proto.Empty, error) {
	doc := docFromProto(req.Document)

	s.mem.Index(doc)

	return &proto.Empty{}, nil
}

func (s *Server) Query(_ context.Context, query *proto.SearchQuery) (*proto.SearchResults, error) {
	q := query.Query

	docs := s.mem.Query(q)

	return docsToResults(docs), nil
}

func (s *Server) GetSet(_ context.Context, req *proto.GetSetRequest) (*proto.SetResponse, error) {
	items := s.mem.GetSet(req.Set)

	return &proto.SetResponse{
		Items: items,
	}, nil
}

func (s *Server) AddToSet(_ context.Context, req *proto.ModifySetRequest) (*proto.Empty, error) {
	s.mem.AddToSet(req.Set, req.Item)

	return &proto.Empty{}, nil
}

func (s *Server) RemoveFromSet(_ context.Context, req *proto.ModifySetRequest) (*proto.Empty, error) {
	s.mem.RemoveFromSet(req.Set, req.Item)

	return &proto.Empty{}, nil
}

func (s *Server) Start() error {
	proto.RegisterCacheServer(s.srv, s)

	s.log.Debug("registered gRPC cache service")

	proto.RegisterCounterServer(s.srv, s)

	s.log.Debug("registered gRPC counter service")

	proto.RegisterKVServer(s.srv, s)

	s.log.Debug("registered gRPC KV service")

	proto.RegisterSearchServer(s.srv, s)

	s.log.Debug("registered gRPC search service")

	proto.RegisterSetServer(s.srv, s)

	s.log.Debug("registered gRPC set service")

	lis, _ := net.Listen("tcp", s.address)

	s.log.Debugf("starting TCP listener on %s", s.address)

	return s.srv.Serve(lis)
}

func (s *Server) ShutDown() {
	s.log.Debug("shutting down")

	s.srv.GracefulStop()
}
