package constant

type ApiResponseEnum string

const (
	SUCCESS ApiResponseEnum = "SUCCESS"
	FAIL    ApiResponseEnum = "FAIL"
)

func (resp ApiResponseEnum) GetCode() string {
	respCodeMap := map[ApiResponseEnum]string{
		SUCCESS: "0",
		FAIL:    "-1",
	}
	return respCodeMap[resp]
}

func (resp ApiResponseEnum) GetMsg() string {
	respMsgMap := map[ApiResponseEnum]string{
		SUCCESS: "SUCCESS",
		FAIL:    "FAIL",
	}
	return respMsgMap[resp]
}

type ApiResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func BuildSuccess(data interface{}) ApiResponse {
	resp := ApiResponse{}
	resp.SetMsg(SUCCESS.GetMsg())
	resp.SetData(data)
	resp.SetCode(SUCCESS.GetCode())
	return resp

}

func BuildFail(msg string) ApiResponse {
	resp := ApiResponse{}
	resp.SetMsg(msg)
	resp.SetData(nil)
	resp.SetCode(FAIL.GetCode())
	return resp
}

func (resp *ApiResponse) GetData() interface{} {
	return resp.Data
}

func (resp *ApiResponse) GetCode() string {
	return resp.Code
}

func (resp *ApiResponse) GetMsg() string {
	return resp.Msg
}

// SetCode 设置 Code 字段的值
func (resp *ApiResponse) SetCode(code string) {
	resp.Code = code
}

// SetMsg 设置 Msg 字段的值
func (resp *ApiResponse) SetMsg(msg string) {
	resp.Msg = msg
}

// SetData 设置 Data 字段的值
func (resp *ApiResponse) SetData(data interface{}) {
	resp.Data = data
}
