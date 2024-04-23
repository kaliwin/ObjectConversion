package http

import (
	"github.com/spf13/cobra"
)

var Init bool

var name = ""

var ConversionHttpCmd = &cobra.Command{
	Use:   "http",
	Short: "http转换",
	Long:  "可使用burp流量转换为proto格式的http组 或从wireshark中提取http ",
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("http转换")
	//},
}

func init() {

	//ConversionHttpCmd.Flags().StringVarP(&name, "as", "s", "转换的名称", "s")
	ConversionHttpCmd.AddCommand(burpCmd)
}
