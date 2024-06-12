package http

import (
	Interface "github.com/kaliwin/Needle/MagicRing/Integrate"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"github.com/kaliwin/Needle/PublicStandard/ObjectHandling"
	"github.com/kaliwin/ObjectConversion/process/HttpRawByteStreamList"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var MergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "合并http数据",
	Long:  "将多个http组 合并成一个列表 以MB为单位",
	Run: func(cmd *cobra.Command, args []string) {
		if output == "" || RawPath == "" || size == 0 {
			cmd.Help()
			return
		}

		Merge(RawPath, output, size)
	},
}

// Merge 合并http数据
func Merge(RawPath string, output string, size int) {

	if _, err := os.ReadDir(output); err != nil {
		log.Fatalln(err)
	}

	list := HttpRawByteStreamList.BurpFlowToHttpRawByteStreamList{
		OutPath: output, MaxSize: size * 1024 * 1024, TmpList: &HttpStructureStandard.HttpRawByteStreamList{},
	}

	stream, err := ObjectHandling.BuildFIleObjectStream(RawPath, true, Interface.ObjectTypeHttpGroupProto)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		next, err := stream.Next()
		if err != nil {
			break
		}

		if http, ok := next.(*HttpStructureStandard.HttpReqAndRes); ok {

			err := list.WriteFile(http)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	// 写入剩余数据
	list.TmpSize = 2
	list.MaxSize = 0
	_ = list.WriteFile(nil)
}

var size int // 大小 单位MB

//var MFilter string  // 过滤条件

//var MOutput string  // 目标目录
//var MRawPath string // 原始数据目录

func init() {

	MergeCmd.Flags().IntVarP(&size, "size", "s", 0, "合并大小 单位MB")
	MergeCmd.Flags().StringVarP(&output, "output", "o", "", "输出文件目录")
	MergeCmd.Flags().StringVarP(&RawPath, "rawPath", "r", "", "原始数据目录")

}
