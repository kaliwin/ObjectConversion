package http

import (
	"github.com/kaliwin/Needle/ObjectConversion/proxyImport"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var ProxyImportCmd = &cobra.Command{
	Use:   "proxy",
	Short: "proxy 导入http数据",
	Long:  "从指定目录中将HttpGroup转为通信实例通过http代理导入到目标程序",
	Run: func(cmd *cobra.Command, args []string) {
		if target == "" || loadPort == "" || rawPath == "" || caPath == "" || caKeyPath == "" {
			_ = cmd.Help()
			return
		}

		httpProxyImport, err := proxyImport.BuildHttpProxyImport(rawPath, caPath, caKeyPath)
		if err != nil {
			log.Println(err)
			return
		}

		go func() {
			err := httpProxyImport.StartHttpServer(loadPort)
			if err != nil {
				log.Println(err)
				os.Exit(8)
			}
		}()
		log.Println("服务启动 : " + bo.address)
		httpProxyImport.Go(target)
		log.Println("任务结束")

	},
}

var target string

var loadPort string

var rawPath string

var caPath string

var caKeyPath string

func init() {
	ProxyImportCmd.Flags().StringVarP(&target, "targetProxy", "t", "", "目标代理地址 只支持http 例:-t http://127.0.0.1:8080")
	ProxyImportCmd.Flags().StringVarP(&rawPath, "rawPath", "r", "", "原始数据目录 目录文件必须为httpGroup")
	ProxyImportCmd.Flags().StringVarP(&loadPort, "loadPort", "l", "", "http代理监听端口 例: -l :8081")
	ProxyImportCmd.Flags().StringVarP(&caPath, "caPath", "c", "", "ca证书路径 例: -c /xx/burpCa.cer")
	ProxyImportCmd.Flags().StringVarP(&caKeyPath, "caKeyPath", "k", "", "ca证书密钥路径 例: -k /xx/burpCa-key.cer")
}
