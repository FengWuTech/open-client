package openclient

import (
	"testing"
)

const APP_ID = "test"
const APP_SECRET_KEY = "6tNgqxXRbo"

func New() *Client {
	return NewClient(APP_ID, APP_SECRET_KEY, BASE_URL_DEV)
}

//测试上报警报
func TestReportWarning(t *testing.T) {

	param := WarningOptions {
		CommunityID :2,
		ReportType:1,
		DeviceID:"5309ee0e3fb14296b17d962055c07dcf",
		AlarmTime:1590646303,
		AlarmPicture:"/images/20200528026243215422184310035.jpeg",
		AlarmType:1,
		AlarmGrade:1,
		AlarmID:"h3ai34ye",
		DeviceIP:"192.168.15.134",
		DeviceSerialID:"nvr3-6",
		PhysicalLocation:"西大门向北",
	}

	err, result := New().ReportWarning(param)
	if err!=nil || result.Code != 200 {
		t.Errorf("request failed, code[%v]msg[%v]", result.Code, result.Msg)
	}

}


//测试警报结果通知
func TestSyncWarningResult(t *testing.T) {

	param := WarningProcessResultOptions {
		CommunityID :2,
		DeviceID:"5309ee0e3fb14296b17d962055c07dcf",
		AlarmID:"h3ai34ye",
		ProcessName:"张晓华",
		ProcessStatus:1,
	}

	err, result := New().SyncWarningProcessResult(param)
	if err!=nil || result.Code != 200 {
		t.Errorf("request failed, code[%v]msg[%v]", result.Code, result.Msg)
	}

}