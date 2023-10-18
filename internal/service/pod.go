package service

import (
	"context"

	pb "github.com/zbnzl/paas-pod/api/pod/v1"
)

type PodService struct {
	pb.UnimplementedPodServer
}

func NewPodService() *PodService {
	return &PodService{}
}

func (s *PodService) AddPod(ctx context.Context, req *pb.PodInfo) (*pb.Response, error) {
	return &pb.Response{}, nil
}
func (s *PodService) DeletePod(ctx context.Context, req *pb.PodId) (*pb.Response, error) {
	return &pb.Response{}, nil
}
func (s *PodService) FindPodByID(ctx context.Context, req *pb.PodId) (*pb.PodInfo, error) {
	return &pb.PodInfo{}, nil
}
func (s *PodService) UpdatePod(ctx context.Context, req *pb.PodInfo) (*pb.Response, error) {
	return &pb.Response{}, nil
}
func (s *PodService) FindAllPod(ctx context.Context, req *pb.FindAll) (*pb.AllPod, error) {
	return &pb.AllPod{}, nil
}
