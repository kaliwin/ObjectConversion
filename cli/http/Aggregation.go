package http

import (
	"github.com/kaliwin/Needle/IO"
	"github.com/kaliwin/Needle/IO/Interface"
	"github.com/kaliwin/Needle/MagicRing/Aggregation/FieldAggregation"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"github.com/spf13/cobra"
	"log"
)

// AggregationCmd 聚合响应体
var AggregationCmd = &cobra.Command{
	Use:   "aggregation",
	Short: "聚合响应体 只接受httpListProto",
	Long:  "迭代所有响应体 按照host和体sha256进行聚合 文件名是 {哈希值}-{path的最后一个/}",
	Run: func(cmd *cobra.Command, args []string) {
		if output == "" || RawPath == "" {
			_ = cmd.Help()
			return
		}
		aggregation := FieldAggregation.NewResFieldAggregation(output) // 创建聚合

		read, err := IO.BuildResourceDescriptionRead(Interface.ResourceDescription{
			Protocol:   Interface.IOFile,
			Path:       RawPath,
			ObjectType: Interface.ObjectTypeHttpGroupListProto,
			Config:     nil,
		})

		if err != nil {
			log.Println(err)
			return
		}

		read.Iteration(func(a any) bool {
			if list, ok := a.(*HttpStructureStandard.HttpRawByteStreamList); ok {
				for _, res := range list.GetHttpRawByteStreamList() {
					err2 := aggregation.Accepting(res) // 执行聚合
					if err2 != nil {
						panic(err2)
					}
				}
			} else {
				log.Println("is not :" + Interface.ObjectTypeHttpGroupListProto)
			}
			return true
		})

	},
}

func init() {
	AggregationCmd.Flags().StringVarP(&output, "output", "o", "", "输出文件目录")
	AggregationCmd.Flags().StringVarP(&RawPath, "rawPath", "r", "", "原始数据目录")
}
