package http

import (
	"github.com/kaliwin/ObjectConversion/process/HttpRawByteStreamList"
	"github.com/spf13/cobra"
	"log"
)

var burpCmd = &cobra.Command{
	Use:   "burp",
	Short: "burp流量转换",
	Long:  "通过MorePossibility-Burp插件导出的流量转换为proto格式的http组",
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

// burpOptions 用于定义burp流量转换的参数
type burpOptions struct {
	address     string // 服务启动地址 包含端口
	fileOutPath string // 输出文件路径
}

var bo = burpOptions{}

func init() {

	burpCmd.Flags().StringVarP(&bo.address, "address", "a", "", "Grpc 服务监听地址")
	burpCmd.Flags().StringVarP(&bo.fileOutPath, "file", "f", "", "输出文件目录")
}
