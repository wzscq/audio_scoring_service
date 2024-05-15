package score

import (
	"os/exec"
	"log"
)

func GetScore(ref,test string)(string) {
	cmd := exec.Command("python3", "audioscore.py",ref,test)
  	// 设置执行命令时的工作目录
  	out, err := cmd.Output()
	if err != nil {
		log.Println("GetScore exec audioscore.py error:",err)
		return ""
	}
	log.Println(string(out))
  	return string(out)
}