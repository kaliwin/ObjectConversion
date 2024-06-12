package HttpRawByteStreamList

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/kaliwin/Needle/MorePossibilityApi"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"github.com/kaliwin/Needle/PublicStandard/sign"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
)

// 对象转换到 HttpRawByteStreamList

// BurpFlowToHttpRawByteStreamList Burp流转换为HttpRawByteStreamList
// 可使用实时流量镜像 、 proxy / map 导出
type BurpFlowToHttpRawByteStreamList struct {
	OutPath string // 输出目录
	TmpList *HttpStructureStandard.HttpRawByteStreamList
	TmpSize int
	MaxSize int
}

// RealTimeTrafficMirroring 实时流量镜像
func (b *BurpFlowToHttpRawByteStreamList) RealTimeTrafficMirroring(server *HttpStructureStandard.HttpReqAndRes) error {
	fmt.Println(server.GetReq().GetUrl())
	return b.WriteFile(server)
}

// WriteFile 写入文件
func (b *BurpFlowToHttpRawByteStreamList) WriteFile(d *HttpStructureStandard.HttpReqAndRes) error {

	if d != nil {

		fileName := d.GetInfo().GetInfo() // info 是否事先有签名
		if fileName == "" {
			fileName, err := sign.HttpBleveIdSign(d) // 签名
			if err != nil {
				return err
			}
			d.Info = &HttpStructureStandard.HttpInfo{
				Info: fileName,
			}
		}

		b.TmpSize += len(d.GetReq().GetData()) + len(d.GetRes().GetData())
		b.TmpList.HttpRawByteStreamList = append(b.TmpList.GetHttpRawByteStreamList(), d)
	}

	if b.TmpSize > b.MaxSize { // 超过最大限制

		if len(b.TmpList.HttpRawByteStreamList) == 0 {
			return nil
		}
		marshal, err := proto.Marshal(b.TmpList)
		if err != nil {
			return err
		}

		fileName := fmt.Sprintf("%d-%d.httpList", len(b.TmpList.HttpRawByteStreamList), uuid.New().ID())

		err = os.WriteFile(b.OutPath+"/"+fileName, marshal, 0666)

		//log.Println("写入文件", b.OutPath+"/"+fileName)

		b.TmpSize = 0
		b.TmpList = &HttpStructureStandard.HttpRawByteStreamList{}

		return err
	}
	return nil
}

// HttpFlowOut 导出流量
func (b *BurpFlowToHttpRawByteStreamList) HttpFlowOut(c context.Context, reqAndRes *HttpStructureStandard.HttpReqAndRes) (*HttpStructureStandard.Str, error) {
	return &HttpStructureStandard.Str{}, b.WriteFile(reqAndRes)
}

// BuildBurpFlowToHttpRawByteStreamListServer 构建Burp流转换为HttpRawByteStreamList
// address 服务地址 outPath 输出目录 每个文件限制在20MB
// 会监听系统信号，接收到信号后会将缓存的流量写入文件并关闭服务
func BuildBurpFlowToHttpRawByteStreamListServer(address string, outPath string) (MorePossibilityApi.BurpGrpcServer, error) {

	_, err2 := os.ReadDir(outPath)
	if err2 != nil {
		return MorePossibilityApi.BurpGrpcServer{}, err2
	}

	// 默认使用200M的接收消息大小
	server, err := MorePossibilityApi.NewBurpGrpcServer(address, grpc.MaxRecvMsgSize(200*1024*1024))
	if err != nil {
		return server, err
	}

	// 创建一个通道来接收操作系统发送的信号
	sigChan := make(chan os.Signal, 1)
	// 将收到的信号发送到 sigChan
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	list := BurpFlowToHttpRawByteStreamList{OutPath: outPath, MaxSize: 20 * 1024 * 1024, TmpList: &HttpStructureStandard.HttpRawByteStreamList{}}

	go func() {
		<-sigChan
		fmt.Print("\r")      // 光标移动到行首
		fmt.Print("\033[2K") // 清除当前行
		list.TmpSize = 2
		list.MaxSize = 0
		_ = list.WriteFile(nil)

		fmt.Println("Ctrl+C pressed. Exiting...") // 在新行中显示消息
		server.Stop()                             // 接受到信号后关闭服务
	}()

	// 注册流量导出服务
	server.RegisterHttpFlowOut(&list)
	// 注册实时流量镜像服务
	server.RegisterRealTimeTrafficMirroring(&list)
	// 返回服务
	return server, nil

}
