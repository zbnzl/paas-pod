package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/zbnzl/paas-pod/api/pod/v1"
	"github.com/zbnzl/paas-pod/internal/biz"
	"github.com/zbnzl/paas-pod/internal/common"
	"github.com/zbnzl/paas-pod/internal/model"
	"strconv"
)

type IPodDataService interface {
	AddPod(*model.Pod) (int64, error)
	DeletePod(int64) error
	UpdatePod(*model.Pod) error
	FindPodByID(int64) (*model.Pod, error)
	FindAllPod() ([]model.Pod, error)
	CreateToK8s(*pb.PodInfo) error
	DeleteFromK8s(*model.Pod) error
	UpdateToK8s(*pb.PodInfo) error
}

type PodService struct {
	pb.UnimplementedPodServer
	puBiz *biz.PodUsecase
	log   *log.Helper
}

func NewPodService(pu *biz.PodUsecase, logger log.Logger) *PodService {
	return &PodService{puBiz: pu, log: log.NewHelper(logger)}
}

// AddPod 创建pod
func (s *PodService) AddPod(ctx context.Context, req *pb.PodInfo) (*pb.Response, error) {
	s.log.WithContext(ctx).Infof("AddPod:%v", req)
	podModel := &model.Pod{}
	if err := common.SwapTo(req, podModel); err != nil {
		return nil, err
	}
	if err := s.puBiz.CreateToK8s(ctx, req); err != nil {
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}
	//操作数据库写入数据
	podID, err := s.puBiz.AddPod(podModel)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	s.log.WithContext(ctx).Infof("Pod add id is: %d", podID)
	return &pb.Response{Msg: "Pod add id is：%d" + strconv.FormatInt(podID, 10)}, nil
}

// DeletePod 删除pod和数据库中的数据
func (s *PodService) DeletePod(ctx context.Context, req *pb.PodId) (*pb.Response, error) {
	//先查找数据
	podModel, err := s.puBiz.FindPodByID(req.Id)
	if err != nil {
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}
	if err := s.puBiz.DeleteFromK8s(ctx, podModel); err != nil {
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}
	return &pb.Response{Msg: "success"}, nil
}

// FindPodByID 查询单个信息
func (s *PodService) FindPodByID(ctx context.Context, req *pb.PodId) (*pb.PodInfo, error) {
	rsp := new(pb.PodInfo)
	//查询pod数据
	podModel, err := s.puBiz.FindPodByID(req.Id)
	if err != nil {
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}
	err = common.SwapTo(podModel, rsp)
	if err != nil {
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}

	return rsp, nil
}

// UpdatePod 更新指定pod
func (s *PodService) UpdatePod(ctx context.Context, req *pb.PodInfo) (*pb.Response, error) {
	//先更新k8s中的pod信息
	err := s.puBiz.UpdateToK8s(ctx, req)
	if err != nil {
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}
	//查询数据库中的pod
	podModel, err := s.puBiz.FindPodByID(req.Id)
	if err != nil {
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}
	err = common.SwapTo(req, podModel)
	if err != nil {
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}

	return &pb.Response{}, s.puBiz.UpdatePod(podModel)
}
func (s *PodService) FindAllPod(ctx context.Context, req *pb.FindAll) (*pb.AllPod, error) {
	rsp := new(pb.AllPod)
	//查询所有pod
	allPod, err := s.puBiz.FindAllPod()
	if err != nil {
		s.log.WithContext(ctx).Error(err)
		return nil, err
	}

	//整理格式
	for _, v := range allPod {
		podInfo := &pb.PodInfo{}
		err := common.SwapTo(v, podInfo)
		if err != nil {
			s.log.WithContext(ctx).Error(err)
			return nil, err
		}
		rsp.PodInfo = append(rsp.PodInfo, podInfo)
	}

	return rsp, nil
}
