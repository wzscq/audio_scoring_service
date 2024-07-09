package score

import (
	"fmt"
	"testing"
	"audioscoring/crv"
	"audioscoring/common"
)

func _TestGetRecItem(t *testing.T) {
	crvClient:=&crv.CRVClient{
		Server:"http://localhost:8200",
		Token:"audio_score_service",
	}

	tempRecItem,errorCode:=GetRecItem("4",crvClient,"audio_score_service")
	if errorCode!=common.ResultSuccess {
		t.Error("GetRecItem error")
	}

	fmt.Println(tempRecItem)
}

func _TestUpdateRecItem(t *testing.T) {
	crvClient:=&crv.CRVClient{
		Server:"http://localhost:8200",
		Token:"audio_score_service",
	}

	tempRecItem,errorCode:=GetRecItem("4",crvClient,"audio_score_service")
	if errorCode!=common.ResultSuccess {
		t.Error("GetRecItem error")
	}

	fmt.Println(tempRecItem)

	tempRecItem.CallerScore="0.5"
	tempRecItem.CalledScore="0.5"
	errorCode=UpdateRecItem(tempRecItem,crvClient,"audio_score_service")
	if errorCode!=common.ResultSuccess {
		t.Error("UpdateRecItem error")
	}
}

func _TestGetAudioFile(t *testing.T) {
	
	fileName:=GetAudioFile("D:/github/crvframe/service/appfile/test1/test2/audio_test/audio_test_file/","caller_audio","4")
	if fileName==nil {
		t.Error("GetAudioFile error")
	}
	fmt.Println(*fileName)

	fileName=GetAudioFile("D:/github/crvframe/service/appfile/test1/test2/audio_test/audio_test_file/","called_audio","4")
	if fileName==nil {
		t.Error("GetAudioFile error")
	}
	fmt.Println(*fileName)

	fileName=GetAudioFile("D:/github/crvframe/service/appfile/test1/test2/audio_test/audio_test_file/","original_audio","4")
	if fileName==nil {
		t.Error("GetAudioFile error")
	}
	fmt.Println(*fileName)
}

func TestAdjustScore(t *testing.T){


	scoreConf:=&common.ScoreConf{
		Adjustments:[]common.ScoreAdjustmentItem{
			{
				MIN:0,
				MAX:4.0,
				Factor:1.3,
			},
		},
	}

	recItem:=&TestRecItem{}
	recItem.CallerScore="	1.3948158025741577"
	recItem.CallerScore=AdjuestScore(recItem.CallerScore,scoreConf)	
	fmt.Println(recItem.CallerScore)
}
