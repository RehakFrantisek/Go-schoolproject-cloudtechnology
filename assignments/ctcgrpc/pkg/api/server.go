package api

import (
	"context"
	"gitlab.com/RehakFrantisek/rehak_clc/assignments/ctcgrpc/pkg"
	"gitlab.com/RehakFrantisek/rehak_clc/assignments/ctcgrpc/pkg/store"
)

type Server struct {
	st store.Store
}

var _ = ApiServer(&Server{})

func NewServer(st store.Store) *Server {
	return &Server{st: st}
}

func (s *Server) Get(ctx context.Context, request *GetRequest) (*GetResponse, error) {
	kv, err := s.st.Get(ctx, request.Key)
	if err != nil {
		return nil, pkg.ToGrpcError(err)
	}

	return &GetResponse{Value: kv}, nil
}

func (s *Server) Put(ctx context.Context, request *PutRequest) (*PutResponse, error) {
	return &PutResponse{}, s.st.Put(ctx, request.Key, request.Value)
}

func (s *Server) Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, error) {
	return &DeleteResponse{}, s.st.Delete(ctx, request.Key)
}

func (s *Server) mustEmbedUnimplementedApiServer() {
}
