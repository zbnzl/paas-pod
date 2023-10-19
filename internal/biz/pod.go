package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/zbnzl/paas-pod/api/pod/v1"
	"github.com/zbnzl/paas-pod/internal/model"
	v1 "k8s.io/api/apps/v1"
	v13 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
)

type IPodRepository interface {
	//初始化表
	InitTable() error
	//根据ID查找数据
	FindPodByID(int64) (*model.Pod, error)
	//创建一条 Pod 数据
	CreatePod(*model.Pod) (int64, error)
	//根据ID删除一条 Pod 数据
	DeletePodByID(int64) error
	//修改一条数据
	UpdatePod(*model.Pod) error
	//查找Pod所有数据
	FindAll() ([]model.Pod, error)
}

// PodUsecase is a Greeter usecase.
type PodUsecase struct {
	repo         IPodRepository
	k8sClientSet *kubernetes.Clientset
	log          *log.Helper
	deployment   *v1.Deployment
}

// NewPodUsecase 创建 PodRepository
func NewPodUsecase(repo IPodRepository, K8sClientSet *kubernetes.Clientset, logger log.Logger) *PodUsecase {
	return &PodUsecase{
		repo:         repo,
		k8sClientSet: K8sClientSet,
		log:          log.NewHelper(logger),
		deployment:   &v1.Deployment{},
	}
}

// CreatePod creates a pod, and returns the new pod.
func (pu *PodUsecase) CreatePod(ctx context.Context, pod *model.Pod) (int64, error) {
	pu.log.WithContext(ctx).Infof("CreatePod: %v", pod)
	return pu.repo.CreatePod(pod)
}

// 创建pod到k8s中
func (pu *PodUsecase) CreateToK8s(ctx context.Context, podInfo *pb.PodInfo) (err error) {
	pu.SetDeployment(podInfo)
	if _, err = pu.k8sClientSet.AppsV1().Deployments(podInfo.PodNamespace).Get(ctx, podInfo.PodName, v12.GetOptions{}); err == nil {
		//可以写自己的业务逻辑
		pu.log.WithContext(ctx).Errorf("Pod %s is exsit", err)
		return errors.New("Pod " + podInfo.PodName + " is exist")
	}
	if _, err = pu.k8sClientSet.AppsV1().Deployments(podInfo.PodNamespace).Create(ctx, pu.deployment, v12.CreateOptions{}); err != nil {
		pu.log.WithContext(ctx).Errorf("Pod create error: %s ", err)
		return err
	}
	pu.log.WithContext(ctx).Info("pod create success")
	return nil
}

// 更新deployment，pod
func (pu *PodUsecase) UpdateToK8s(ctx context.Context, podInfo *pb.PodInfo) (err error) {
	pu.SetDeployment(podInfo)
	if _, err = pu.k8sClientSet.AppsV1().Deployments(podInfo.PodNamespace).Get(ctx, podInfo.PodName, v12.GetOptions{}); err != nil {
		pu.log.WithContext(ctx).Error(err)
		return errors.New("Pod " + podInfo.PodName + " not found")
	} else {
		//如果存在
		if _, err = pu.k8sClientSet.AppsV1().Deployments(podInfo.PodNamespace).Update(ctx, pu.deployment, v12.UpdateOptions{}); err != nil {
			pu.log.WithContext(ctx).Error(err)
			return err
		}
		pu.log.WithContext(ctx).Info("pod update success")
		return nil
	}

}

// 删除pod
func (pu *PodUsecase) DeleteFromK8s(ctx context.Context, pod *model.Pod) (err error) {
	if err = pu.k8sClientSet.AppsV1().Deployments(pod.PodNamespace).Delete(ctx, pod.PodName, v12.DeleteOptions{}); err != nil {
		pu.log.WithContext(ctx).Error(err)
		//删除失败写自己的业务逻辑
		return err
	} else {
		if err := pu.DeletePod(pod.ID); err != nil {
			pu.log.WithContext(ctx).Error(err)
			return err
		}
		pu.log.WithContext(ctx).Info("pod %d delete success", strconv.FormatInt(pod.ID, 10))
	}
	return
}

func (pu *PodUsecase) SetDeployment(podInfo *pb.PodInfo) {
	deployment := &v1.Deployment{}
	deployment.TypeMeta = v12.TypeMeta{
		Kind:       "deployment",
		APIVersion: "v1",
	}
	deployment.ObjectMeta = v12.ObjectMeta{
		Name:      podInfo.PodName,
		Namespace: podInfo.PodNamespace,
		Labels: map[string]string{
			"app-name": podInfo.PodName,
			"author":   "zl",
		},
	}
	deployment.Name = podInfo.PodName
	deployment.Spec = v1.DeploymentSpec{
		//副本个数
		Replicas: &podInfo.PodReplicas,
		Selector: &v12.LabelSelector{
			MatchLabels: map[string]string{
				"app-name": podInfo.PodName,
			},
			MatchExpressions: nil,
		},
		Template: v13.PodTemplateSpec{
			ObjectMeta: v12.ObjectMeta{
				Labels: map[string]string{
					"app-name": podInfo.PodName,
				},
			},
			Spec: v13.PodSpec{
				Containers: []v13.Container{
					{
						Name:            podInfo.PodName,
						Image:           podInfo.PodImage,
						Ports:           pu.getContainerPort(podInfo),
						Env:             pu.getEnv(podInfo),
						Resources:       pu.getResources(podInfo),
						ImagePullPolicy: pu.getImagePullPolicy(podInfo),
					},
				},
			},
		},
		Strategy:                v1.DeploymentStrategy{},
		MinReadySeconds:         0,
		RevisionHistoryLimit:    nil,
		Paused:                  false,
		ProgressDeadlineSeconds: nil,
	}
	pu.deployment = deployment
}

func (pu *PodUsecase) getContainerPort(podInfo *pb.PodInfo) (containerPort []v13.ContainerPort) {
	for _, v := range podInfo.PodPort {
		containerPort = append(containerPort, v13.ContainerPort{
			Name:          "port-" + strconv.FormatInt(int64(v.ContainerPort), 10),
			ContainerPort: v.ContainerPort,
			Protocol:      pu.getProtocol(v.Protocol),
		})
	}
	return
}

func (pu *PodUsecase) getProtocol(protocol string) v13.Protocol {
	switch protocol {
	case "TCP":
		return "TCP"
	case "UDP":
		return "UDP"
	case "SCTP":
		return "SCTP"
	default:
		return "TCP"
	}
}

func (pu *PodUsecase) getEnv(podInfo *pb.PodInfo) (envVar []v13.EnvVar) {
	for _, v := range podInfo.PodEnv {
		envVar = append(envVar, v13.EnvVar{
			Name:      v.EnvKey,
			Value:     v.EnvValue,
			ValueFrom: nil,
		})
	}
	return
}

func (pu *PodUsecase) getResources(podInfo *pb.PodInfo) (source v13.ResourceRequirements) {
	//最大能够使用多少资源
	source.Limits = v13.ResourceList{
		"cpu":    resource.MustParse(strconv.FormatFloat(float64(podInfo.PodCpuMax), 'f', 6, 64)),
		"memory": resource.MustParse(strconv.FormatFloat(float64(podInfo.PodMemoryMax), 'f', 6, 64)),
	}
	//满足最少使用的资源量
	//@TODO 自己实现动态设置
	source.Requests = v13.ResourceList{
		"cpu":    resource.MustParse(strconv.FormatFloat(float64(podInfo.PodCpuMin), 'f', 6, 64)),
		"memory": resource.MustParse(strconv.FormatFloat(float64(podInfo.PodMemoryMin), 'f', 6, 64)),
	}
	return
}

func (pu *PodUsecase) getImagePullPolicy(podInfo *pb.PodInfo) v13.PullPolicy {
	switch podInfo.PodPullPolicy {
	case "Always":
		return "Always"
	case "Never":
		return "Never"
	case "IfNotPresent":
		return "IfNotPresent"
	default:
		return "Always"
	}
}

// 添加Pod
func (pu *PodUsecase) AddPod(pod2 *model.Pod) (int64, error) {
	return pu.repo.CreatePod(pod2)
}

// 删除
func (pu *PodUsecase) DeletePod(podID int64) error {
	return pu.repo.DeletePodByID(podID)
}

// 更新
func (pu *PodUsecase) UpdatePod(pod2 *model.Pod) error {
	return pu.repo.UpdatePod(pod2)
}

// 单个查找
func (pu *PodUsecase) FindPodByID(podID int64) (*model.Pod, error) {
	return pu.repo.FindPodByID(podID)
}

// 查找所有
func (pu *PodUsecase) FindAllPod() ([]model.Pod, error) {
	return pu.repo.FindAll()
}
