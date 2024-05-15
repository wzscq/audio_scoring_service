package score

import (
	"audioscoring/crv"
	"audioscoring/common"
	"log"
	"path/filepath"
	"fmt"
)

type TestRecItem struct {
	RecID string
	CallerScore string
	CalledScore string 
	Version string 	
}

func GetRecItem(recID string,crvClient *crv.CRVClient,token string)(*TestRecItem,int) {
	req:=crv.CommonReq{
		ModelID:"audio_test_file",
		Fields:&[]map[string]interface{}{
			map[string]interface{}{
				"field":"id",
			},
			map[string]interface{}{
				"field":"version",
			},
		},
		Filter:&map[string]interface{}{
			"id":map[string]interface{}{
				"Op.eq":recID,
			},
		},
	}

	rsp,errorCode:=crvClient.Query(&req,token)
	if errorCode!=common.ResultSuccess {
		return nil,errorCode
	}

	if rsp.Result==nil{
		return nil,common.ResultNoRecData
	}

	//获取result中的list
	resultMap,ok:=rsp.Result.(map[string]interface{})
	if !ok {
		log.Println("GetRecItem can not be converted to map")
		return nil,common.ResultNoRecData
	}

	list,ok:=resultMap["list"]
	if !ok {
		log.Println("GetRecItem queryResult no list")
		return nil,common.ResultNoRecData
	}

	recList,ok:=list.([]interface{})
	if !ok || len(recList)<=0 {
		log.Println("GetRecItem queryResult no list")
		return nil,common.ResultNoRecData
	}

	//获取第一条记录
	rec,ok:=recList[0].(map[string]interface{})
	if !ok {
		log.Println("GetRecItem queryResult row 0 can not convert to map")
		return nil,common.ResultNoRecData
	}

	recItem:=&TestRecItem{
		RecID:recID,
		Version:rec["version"].(string),
	}

	return recItem,common.ResultSuccess
}

func UpdateRecItem(recItem *TestRecItem,crvClient *crv.CRVClient,token string)(int){
	req:=crv.CommonReq{
		ModelID:"audio_test_file",
		List:&[]map[string]interface{}{
			map[string]interface{}{
				"id":recItem.RecID,
				"version":recItem.Version,
				"caller_score":recItem.CallerScore,
				"called_score":recItem.CalledScore,
				"_save_type":"update",
			},
		},
	}
	_,errorCode:=crvClient.Save(&req,token)
	return errorCode
}

func GetAudioFile(path,audioType,recID string)(*string){
	fileMatch:=fmt.Sprintf("%s/%s_row%s_*",path,audioType,recID)
	files,_:=filepath.Glob(fileMatch)
	if len(files)<=0 {
		return nil
	}

	return &files[0]
}