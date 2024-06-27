package http

import (
	"github.com/kaliwin/ObjectConversion/process/HttpRawByteStreamList"
	"github.com/spf13/cobra"
	"log"
)

var ProxyImportCmd = &cobra.Command{
	Use:   "proxy",
	Short: "proxy 导入http数据",
	Long:  "从指定目录中将HttpList转为通信实例通过http代理导入到目标程序",
	Run: func(cmd *cobra.Command, args []string) {
		if bo.address == "" || bo.fileOutPath == "" {
			cmd.Help()
			return
		}
		server, err := HttpRawByteStreamList.BuildBurpFlowToHttpRawByteStreamListServer(bo.address, bo.fileOutPath)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("服务启动成功 : " + bo.address)
		server.Start()

	},
}

var target string

var loadPort string

var rawPath string

func init() {
	ProxyImportCmd.Flags().StringVarP(&target, "target", "t", "", "目标地址 只支持http 例:-t 127.0.0.1:8080")
	ProxyImportCmd.Flags().StringVarP(&rawPath, "rawPath", "r", "", "原始数据目录 目录文件必须为httpList")
	ProxyImportCmd.Flags().StringVarP(&loadPort, "loadPort", "l", "", "http代理监听端口 例: -l :8081")
}
