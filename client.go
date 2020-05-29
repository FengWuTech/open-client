package openclient

const (
	BASE_URL_PRODUCT = "http://pms-api.gmtech.top"
	BASE_URL_DEV     = "http://pms-api-dev.gmtech.top"
)

type Client struct {
	AppID     string
	AppSecret string
	BaseURL   string
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"time"`
}

func NewClient(appID string, appSecret string, baseURL string) *Client {
	return &Client{
		AppID:     appID,
		AppSecret: appSecret,
		BaseURL:   baseURL,
	}
}



