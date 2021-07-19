package Dto

import "github.com/WenkanHuang/gin_gorm/Model"

type GroupDto struct {
	GroupId   uint   `gorm:"primaryKey;" json:"groupId" uri:"groupId"`
	GroupName string `gorm:"varchar(255);unique" json:"groupName" uri:"groupName"`
	ItemCOUNT int    `json:"item" uri:"item"`
}

func TodoGroupId(group Model.Group) GroupDto {
	return GroupDto{
		GroupId:   group.GroupId,
		GroupName: group.GroupName,
		ItemCOUNT: group.ItemCOUNT,
	}
}
