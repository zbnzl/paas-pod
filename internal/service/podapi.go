package service

import (
	"context"

	pb "github.com/zbnzl/paas-pod/api/podapi/v1"
)

type PodApiService struct {
	pb.UnimplementedPodApiServer
}

func NewPodApiService() *PodApiService {
	return &PodApiService{}
}

func (s *PodApiService) FindPodById(ctx context.Context, req *pb.api_pod_v1_PodId) (*pb.api_pod_v1_PodInfo, error) {
	return &pb.api_pod_v1_PodInfo{}, nil
}
func (s *PodApiService) AddPod(ctx context.Context, req *pb.api_pod_v1_PodInfo) (*pb.api_pod_v1_Response, error) {
	return &pb.api_pod_v1_Response{}, nil
}
func (s *PodApiService) DeletePodById(ctx context.Context, req *pb.api_pod_v1_PodId) (*pb.api_pod_v1_Response, error) {
	return &pb.api_pod_v1_Response{}, nil
}
func (s *PodApiService) UpdatePod(ctx context.Context, req *pb.api_pod_v1_PodInfo) (*pb.api_pod_v1_Response, error) {
	return &pb.api_pod_v1_Response{}, nil
}
func (s *PodApiService) Call(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}
