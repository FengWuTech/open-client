package openclient

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/parnurzeal/gorequest"
	"github.com/xxjwxc/public/errors"
	"log"
)


type WarningOptions struct {
	CommunityID int `json:"community_id"  valid:"Required"`  //小区ID
	ReportType int `json:"report_type"  valid:"Required"`  //上报类型 1自动 2手动
	DeviceID string `json:"device_id"  valid:"Required"`  //设备编号
	AlarmTime int `json:"alarm_time"  valid:"Required"`  //报警时间,unix时间戳 1590646303
	AlarmPicture string `json:"alarm_picture"  valid:"Required"`  //报警图片的url
	AlarmType int `json:"alarm_type"  valid:"Range(1, 4)"`  //报警类型 1. fire_escape_grade 2. danger_invasion_grade 3. camera_failure 4.garbage_lasting_grade
	AlarmGrade int `json:"alarm_grade"  valid:"Range(1, 6)"`  //报警等级等级 1. 1级 2. 2级 3. 3级 设备故障等级：1 .视频信号丢失 ；2. 阵列异常； 3.IP地址冲突；  4.网络断开；5. IPC 连接异常； 6. 视频制式不匹配
	AlarmID string `json:"alarm_id"  valid:"Required"`  //报警id(按照该id去sync处理结果)
	DeviceIP string `json:"device_ip"  valid:"IP"`  //设备ip地址
	DeviceSerialID string `json:"device_serial_id"  valid:"Required"`  //设备序列号信息
	PhysicalLocation string `json:"physical_location"  valid:"Required"`  //设备物理地址
}

//视觉上报报警
func (cli *Client) ReportWarning(param WarningOptions) (error, Response) {

	path := "/fpms/open/vision/report"
	query := cli.genQuery(nil)
	rawURL := cli.BaseURL + path + "?" + query

	// param check
	valid := validation.Validation{}
	b, err := valid.Valid(&param)
	if err != nil {
		return err,Response{}
	}
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		return errors.New("参数错误"),Response{}
	}

	paramJSON, _ := json.Marshal(param)
	paramStr := string(paramJSON)

	var urlResp Response
	request := gorequest.New()
	_, _, errs := request.Post(rawURL).Send(paramStr).EndStruct(&urlResp)
	if errs != nil || len(errs) > 0 {
		return errs[0], urlResp
	}
	return nil, urlResp
}