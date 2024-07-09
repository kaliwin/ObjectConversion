package http

import (
	"github.com/kaliwin/Needle/ObjectConversion/ByteClassification/Classification"
	"github.com/spf13/cobra"
)

var byteSeparationCmd = &cobra.Command{
	Use:   "byteSeparation",
	Short: "字节分离",
	Long:  "将文件数据按照类型分类",
	Run: func(cmd *cobra.Command, args []string) {

		if rawPath == "" || output == "" {

			_ = cmd.Help()
			return
		}
		classification, err := Classification.NewByteClassification(output, rawPath)
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		err = classification.Do()
		if err != nil {
			cmd.PrintErr(err)
			return
		}
	},
}

func init() {

	byteSeparationCmd.Flags().StringVarP(&rawPath, "rawPath", "r", "", "原始数据目录")
	byteSeparationCmd.Flags().StringVarP(&output, "output", "o", "", "输出文件目录")
}
