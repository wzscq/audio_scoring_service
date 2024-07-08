package score

import (
	"audioscoring/crv"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"audioscoring/common"
)

type ScoreController struct {
	CRVClient *crv.CRVClient
	AudioPath string
	ScoreConf *common.ScoreConf
}

func (sc *ScoreController) Bind(router *gin.Engine) {
	router.POST("/predict", sc.scoring)
}

func (sc *ScoreController) scoring(c *gin.Context) {
	log.Println("start ScoreController scoring")

	var header crv.CommonHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		log.Println(err)
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		log.Println("end ScoreController scoring with error")
		return
	}	

	var rep crv.CommonReq
	if err := c.BindJSON(&rep); err != nil {
		log.Println(err)
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		log.Println("end ScoreController scoring with error")
		return
  }	

	if rep.SelectedRowKeys==nil || len(*rep.SelectedRowKeys)==0 {
		log.Println("end ScoreController scoring with error:SelectedRowKeys is empty")
		rsp:=common.CreateResponse(common.CreateError(common.ResultWrongRequest,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		return
	}

	testRecID:=(*rep.SelectedRowKeys)[0]
	log.Println("testRecID:",testRecID)
	recItem,errorCode:=GetRecItem(testRecID,sc.CRVClient,header.Token)
	if errorCode!=common.ResultSuccess {
		log.Println("end ScoreController scoring with error")
		rsp:=common.CreateResponse(common.CreateError(errorCode,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		return
	}

	originalAudio:=GetAudioFile(sc.AudioPath,"original_audio",testRecID)
	if originalAudio==nil {
		log.Println("end ScoreController scoring with error:original_audio is empty")
		rsp:=common.CreateResponse(common.CreateError(common.ResultOriginalAudioNoExist,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		return
	}

	callerAudio:=GetAudioFile(sc.AudioPath,"caller_audio",testRecID)
	calledAudio:=GetAudioFile(sc.AudioPath,"called_audio",testRecID)
	if callerAudio==nil&&calledAudio==nil {
		log.Println("end ScoreController scoring with error:test_audio is empty")
		rsp:=common.CreateResponse(common.CreateError(common.ResultTestNoExist,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		return	
	}
	
	if callerAudio != nil {
		recItem.CallerScore=GetScore(*originalAudio,*callerAudio)
		recItem.CallerScore=AdjuestScore(recItem.CallerScore,sc.ScoreConf)
	}

	if calledAudio != nil {
		recItem.CalledScore=GetScore(*originalAudio,*calledAudio)
		recItem.CalledScore=AdjuestScore(recItem.CalledScore,sc.ScoreConf)
	}

	errorCode=UpdateRecItem(recItem,sc.CRVClient,header.Token)
	if errorCode!=common.ResultSuccess {
		log.Println("end ScoreController scoring with error")
		rsp:=common.CreateResponse(common.CreateError(errorCode,nil),nil)
		c.IndentedJSON(http.StatusOK, rsp)
		return
	}
	
	rsp:=common.CreateResponse(nil,nil)
	c.IndentedJSON(http.StatusOK, rsp)
}