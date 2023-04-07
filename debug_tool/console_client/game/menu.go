package game

import (
	"fmt"
	"log"
)

var (
	mainIndex uint32

	menuMap = make(map[uint32]func())

	sendMsgMap = make(map[string]func())
)

func InitMenu() {
	menuMap[0] = mainMenu
	menuMap[1] = accountMenu
	menuMap[2] = itemMenu
	menuMap[3] = taskMenu
	menuMap[4] = gmMenu

}

func ProcessMenu(menuId uint32) {
	if mainIndex > 0 {
		if menuId == 0 {
			mainIndex = 0
			ShowMenu()
		} else {
			cmd := fmt.Sprintf("%d_%d", mainIndex, menuId)
			if f, ok := sendMsgMap[cmd]; ok {
				f()
			} else {
				log.Printf("cmd:%v not exist", cmd)
			}
		}
	} else {
		mainIndex = menuId
		ShowMenu()
	}

}

func ShowMenu() {
	if f, ok := menuMap[mainIndex]; ok {
		f()
	}
}

func mainMenu() {
	var line []string
	line = append(line, "----------主菜单----------")
	line = append(line, "1. 账号 ")
	line = append(line, "2. 背包")
	line = append(line, "3. 任务")
	line = append(line, "4. GM命令")

	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
}

func accountMenu() {
	var line []string
	line = append(line, "----------账号----------")
	line = append(line, "0. 返回 ")
	line = append(line, "1. 登录 ")
	line = append(line, "2. 注册 ")
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
	sendMsgMap["1_1"] = sendLoginMsg
	sendMsgMap["1_2"] = sendRegist
}

func itemMenu() {
	var line []string
	line = append(line, "----------背包----------")
	line = append(line, "0. 返回 ")
	line = append(line, "1. 加载背包 ")
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
	sendMsgMap["2_1"] = sendRequestLoadItems
}
func taskMenu() {
	var line []string
	line = append(line, "----------任务----------")
	line = append(line, "0. 返回 ")
	line = append(line, "1. 加载任务 ")
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
	sendMsgMap["3_1"] = sendLoadTasks
}

func gmMenu() {
	var line []string
	line = append(line, "----------GM命令----------")
	line = append(line, "0. 返回 ")
	line = append(line, "1. 增加道具 ")
	line = append(line, "2. 加载配置 ")
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
	sendMsgMap["4_1"] = sendGmAddItem
	sendMsgMap["4_2"] = sendGmReload
}
