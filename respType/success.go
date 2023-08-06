package respType

type SuccessResp struct {
	Code int         `json:"Code"` // 用户名
	Data interface{} `json:"Data"` // 密码
}
