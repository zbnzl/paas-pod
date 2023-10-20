package data

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/jinzhu/gorm"
	"github.com/zbnzl/paas-pod/internal/biz"
	"github.com/zbnzl/paas-pod/internal/model"
)

type PodRepository struct {
	mysqlDb *gorm.DB
}

func NewPodRepository(mysqlDb *gorm.DB) biz.IPodRepository {
	return &PodRepository{mysqlDb: mysqlDb}
}

// 初始化表
func (u *PodRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Pod{}, &model.PodEnv{}, &model.PodPort{}).Error
}

// 根据ID查找Pod信息
func (u *PodRepository) FindPodByID(podID int64) (pod *model.Pod, err error) {
	pod = &model.Pod{}
	if err = u.mysqlDb.Preload("PodEnv").Preload("PodPort").First(pod, podID).Error; gorm.IsRecordNotFoundError(err) {
		return pod, errors.NotFound("", err.Error())
	}
	return pod, err
}

// 创建 Pod
func (u *PodRepository) CreatePod(pod *model.Pod) (int64, error) {
	return pod.ID, u.mysqlDb.Create(pod).Error
}

// 根据 ID 删除 Pod 信息
func (u *PodRepository) DeletePodByID(podID int64) error {
	tx := u.mysqlDb.Begin()
	//遇到问题回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}

	//彻底删除 POD 信息
	if err := u.mysqlDb.Where("id = ?", podID).Delete(&model.Pod{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//彻底删除 podenv 信息
	if err := u.mysqlDb.Where("pod_id = ?", podID).Delete(&model.PodEnv{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//彻底删除 podport 信息
	if err := u.mysqlDb.Where("pod_id = ?", podID).Delete(&model.PodPort{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// 更新Pod信息
func (u *PodRepository) UpdatePod(pod *model.Pod) error {
	return u.mysqlDb.Model(pod).Update(pod).Error
}

// 获取结果集合
func (u *PodRepository) FindAll() (podAll []model.Pod, err error) {
	return podAll, u.mysqlDb.Find(&podAll).Error
}
