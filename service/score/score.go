package score

import (
	"os/exec"
	"log"
	"audioscoring/common"
	"strconv"
)

func GetScore(ref,test string)(string) {
	cmd := exec.Command("python3", "audioscore.py",ref,test)
  	// 设置执行命令时的工作目录
  	out, err := cmd.Output()
	if err != nil {
		log.Println("GetScore exec audioscore.py error:",err)
		return ""
	}

	result:=string(out)
	log.Println("GetScore py result",result)
  	return result
}

func AdjuestScore(score string,scoreConf *common.ScoreConf)(string){
	log.Println("AdjuestScore start ",score)
	//按照匹配对分值做一个调整
	if scoreConf!=nil {
		//string to float
		fres, _ := strconv.ParseFloat(score, 64) 
		for _,item:= range scoreConf.Adjustments {
			log.Println("AdjuestScore item ",item,fres)
			if item.MIN<fres && item.MAX>=fres {
				log.Println("AdjuestScore by ",item)
				fres=fres*item.Factor	
				score = strconv.FormatFloat(fres, 'f', -1, 64)
				break
			}
		}
	}

	log.Println("AdjuestScore result ",score)
	return score
}