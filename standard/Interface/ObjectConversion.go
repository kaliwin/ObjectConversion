package Interface

import (
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
)

// ConversionToHttpRawByteStreamList 将目标转换为proto格式的http组
type ConversionToHttpRawByteStreamList interface {
	Conversion(any) (*HttpStructureStandard.HttpRawByteStreamList, error) // 转换
}

// ReportConversion 报告转换
// 将ManDown 标准报告转换为任意类型报告
// 将任意类型报告转换为ManDown 标准报告
type ReportConversion interface {
	ConversionManDown(any) (any, error) // 转换为ManDown 标准报告
	ConversionAny(any) (any, error)     // 转换为任意类型报告
}
