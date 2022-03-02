package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

// 选择选项输入
func selectInput(label string, items []string) string {
	prompt := promptui.Select{
		Label:     label,
		HideHelp:  true,
		Size:      10,
		CursorPos: 0,
		Items:     items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

// 添加文本输入
func addInput(label string) string {
	prompt := promptui.Prompt{
		HideEntered: true,
		Label:       label,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

// 获得 commit message 信息
func getMessage() string {
	type1 := selectInput("Type. 选择一个 commit 类型", []string{
		"🎃 feat:	新功能",
		"🎃 fix:	修补错误或异常",
		"🎃 docs:	修改文档",
		"🎃 style:	格式",
		"🎃 refactor:	重构",
		"🎃 perf:	提高性能的代码修改",
		"🎃 test:	增加或修改单元测试",
		"🎃 chore:	构建过程或辅助工具的变动",
		"🎃 ci: CI	配置修改",
	})

	scope := selectInput("确认你的选择", []string{
		"🎃 service:	服务层",
		"🎃 model:	模型层",
		"🎃 repository:	存储层",
		"🎃 library:	库包",
		"🎃 other:	其他",
	})

	msg := fmt.Sprintf("\n%s\n\n%s\n\n", type1, scope)
	return msg
}

func main() {
	msg := getMessage()
	fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n%s", 0x1B, "\n\ncommit message:", 0x1B, msg)
}
