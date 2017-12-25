package controller

import (
	"encoding/json"
	"vip_services/models"
	"strings"
	pb "vip_services/vip_services"
)

type ApplyController struct{}

func (s *ApplyController) GetOneApplyById(in *pb.ApplyIdRequest) string {

	ApplyModel := new(models.ApplyModel)

	userInfo, _ := ApplyModel.GetOneByApply(in.ApplyId)

	userInfoJson, _ := json.Marshal(userInfo)

	return string(userInfoJson)

}

func (s *ApplyController) GetAllApplyByUserId(in *pb.ApplyUidRequest) string {

	ApplyModel := new(models.ApplyModel)

	userArr := strings.Split(in.UserId, ",")

	userInfo, _ := ApplyModel.GetByUserIdByApply(userArr)

	userInfoJson, _ := json.Marshal(userInfo)

	return string(userInfoJson)

}
