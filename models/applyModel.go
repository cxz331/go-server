package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type ApplyModel struct {
	Id            int    `orm:"column(id);auto"json:"id"`
	BizUid        int    `orm:"column(biz_uid);null"json:"biz_uid"`
	UserId        int    `orm:"column(user_id);null"json:"user_id"`
	UserName      string `orm:"column(user_name);null"json:"user_name"`
	UserMobile    int64  `orm:"column(user_mobile);null"json:"user_mobile"`
	PostId        int    `orm:"column(post_id);null"json:"post_id"`
	CityId        int    `orm:"column(city_id);null"json:"city_id"`
	AddrId        int    `orm:"column(addr_id);null"json:"addr_id"`
	Work_timeId   int    `orm:"column(work_time_id);null"json:"work_time_id"`
	ListingStatus int    `orm:"column(listing_status);null"json:"listing_status"`
	CaPlatform    int    `orm:"column(ca_platform);null"json:"ca_platform"`
	CaSource      string `orm:"column(ca_source);null"json:"ca_source"`
	WorkTime      string `orm:"column(work_time);null"json:"work_time"`
	CreateAt      int    `orm:"column(create_at);null"json:"create_at"`
	ModifyAt      int    `orm:"column(modify_at);null"json:"modify_at"`
	WorkStatus    int    `orm:"column(work_status);null"json:"work_status"`
	MemberStatus  int    `orm:"column(member_status);null"json:"member_status"`
	Remark        string `orm:"column(remark);null"json:"remark"`
	CaCampaign    string `orm:"column(ca_campaign);null"json:"ca_campaign"`
	EntryDate     int    `orm:"column(entry_date);null"json:"entry_date"`
	ApplySource   int    `orm:"column(apply_source);null"json:"apply_source"`
	IsReadPhone   int    `orm:"column(is_read_phone);null"json:"is_read_phone"`
	WorkTimeStart string `orm:"column(work_time_start);null"json:"work_time_start"`
	WorkTimeEnd   string `orm:"column(work_time_end);null"json:"work_time_end"`
	SubmitId      int    `orm:"column(submit_id);null"json:"submit_id"`
	ListingRemark string `orm:"column(listing_remark);null"json:"listing_remark"`
}

var (
	LISTING_STATUS_EVALUATE = []int{-5, -1, 3, 5, 11, 12, 20, 21}
)

func (s *ApplyModel) TableName() string {
	return "jz_apply"
}

func init() {
	orm.RegisterModel(new(ApplyModel))
}

func (s *ApplyModel) GetOneByApply(id int32) (*ApplyModel, error) {
	r := new(ApplyModel)
	mo := orm.NewOrm()
	mo.Using("test")
	err := mo.QueryTable(r.TableName()).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *ApplyModel) GetByUserIdByApply(userId []string) ([]*ApplyModel, error) {
	r := new(ApplyModel)
	list := make([]*ApplyModel, 0)
	orm.NewOrm().QueryTable(r.TableName()).Filter("user_id__in", userId).Filter("listing_status__in", LISTING_STATUS_EVALUATE).All(&list)
	return list, nil
}
