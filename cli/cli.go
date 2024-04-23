package cli

import (
	"github.com/kaliwin/ObjectConversion/cli/http"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ObjectConversion",
	Short: "将同类型不同实例转换为ManDown标准实例",
	//Long:    "可食用burp流量转换为proto格式的http组 或从wireshark中提取http",
	Version: "1.0.0",
}

func init() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
	RootCmd.AddCommand(http.ConversionHttpCmd)
}
