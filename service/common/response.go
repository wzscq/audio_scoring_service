package common

type CommonRsp struct {
	ErrorCode int `json:"errorCode"`
	Message string `json:"message"`
	Error bool `json:"error"`
	Result interface{} `json:"result"`
	Params map[string]interface{} `json:"params"`
}

type CommonError struct {
	ErrorCode int `json:"errorCode"`
	Params map[string]interface{} `json:"params"`
}

const (
	ResultSuccess = 10000000
	ResultWrongRequest = 10000001
	ResultSaveDataError = 10100010
	ResultQueryRequestError = 10100007
	ResultNoRecData = 10100008
	ResultOriginalAudioNoExist = 10100009
	ResultTestNoExist = 10100011
)

var errMsg = map[int]CommonRsp{
	ResultSuccess:CommonRsp{
		ErrorCode:ResultSuccess,
		Message:"操作成功",
		Error:false,
	},
	ResultWrongRequest:CommonRsp{
		ErrorCode:ResultWrongRequest,
		Message:"请求参数错误，请检查参数是否完整，参数格式是否正确",
		Error:true,
	},
	ResultSaveDataError:CommonRsp{
		ErrorCode:ResultSaveDataError,
		Message:"保存数据到数据时发生错误，请与管理员联系处理",
		Error:true,
	},
	ResultQueryRequestError:CommonRsp{
		ErrorCode:ResultQueryRequestError,
		Message:"查询数据失败，请与管理员联系处理",
		Error:true,
	},
	ResultNoRecData:CommonRsp{
		ErrorCode:ResultNoRecData,
		Message:"查询记录数据时未找到对应数据，请检查参数是否正确",
		Error:true,
	},
	ResultOriginalAudioNoExist:CommonRsp{
		ErrorCode:ResultOriginalAudioNoExist,
		Message:"原始音频文件不存在，请检确认文件是否已正确上传，且格式为WAV格式",
		Error:true,
	},
	ResultTestNoExist:CommonRsp{
		ErrorCode:ResultTestNoExist,
		Message:"主被叫测试音频文件都不存在，请检确认文件是否已正确上传，且格式为WAV格式",
		Error:true,
	},
}

func CreateResponse(err *CommonError,result interface{})(*CommonRsp){
	if err==nil {
		commonRsp:=errMsg[ResultSuccess]
		commonRsp.Result=result
		return &commonRsp
	}

	commonRsp:=errMsg[err.ErrorCode]
	commonRsp.Result=result
	commonRsp.Params=err.Params
	return &commonRsp
}

func CreateError(errorCode int,params map[string]interface{})(*CommonError){
	return &CommonError{
		ErrorCode:errorCode,
		Params:params,
	}
}