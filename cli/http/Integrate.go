package http

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/kaliwin/Needle/IO"
	"github.com/kaliwin/Needle/IO/Interface"
	"github.com/kaliwin/Needle/PublicStandard/HttpStructureStandard/grpc/HttpStructureStandard"
	"github.com/kaliwin/Needle/PublicStandard/sign"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"net/url"
	"os"
	"strings"
)

var IntegrateCmd = &cobra.Command{
	Use:   "integrate",
	Short: "数据整合",
	Long:  "将http数据通过url进行分类整合 并且会写入唯一标识符,可参照burp中的map",
	Run: func(cmd *cobra.Command, args []string) {
		if output == "" || RawPath == "" {
			cmd.Help()
			return
		}
		err := Diversion(RawPath, filter, output)
		if err != nil {
			cmd.Println(err)
		}
	},
}

var filter string  // 过滤条件
var output string  // 目标目录
var RawPath string // 原始数据目录

func init() {

	IntegrateCmd.Flags().StringVarP(&filter, "filter", "f", "", "url过滤条件 例如: -f www.baidu.com")
	IntegrateCmd.Flags().StringVarP(&output, "output", "o", "", "输出文件目录")
	IntegrateCmd.Flags().StringVarP(&RawPath, "rawPath", "r", "", "原始数据目录")

}

// Diversion 数据分流
// urlFilter 过滤条件 要包含的url
func Diversion(rawPath string, urlFilter string, outPath string) error {

	stream, err := IO.BuildResourceDescriptionRead(Interface.ResourceDescription{
		Protocol:   Interface.IOFile,
		Path:       rawPath,
		ObjectType: Interface.ObjectTypeHttpGroupListProto,
		Config:     nil,
	})
	if err != nil {
		return err
	}

	for {
		next, err := stream.Next()
		if err != nil {
			break
		}
		if list, ok := next.(*HttpStructureStandard.HttpRawByteStreamList); ok {
			for _, res := range list.GetHttpRawByteStreamList() {

				if res.GetReq().GetData() == nil && res.GetReq().GetData() == nil {
					continue
				}

				u := res.GetReq().GetUrl()
				//fmt.Println(u)
				if !strings.Contains(u, urlFilter) { // 过滤条件
					continue
				}

				parse, err := url.Parse(u)
				if err != nil {
					return err
				}

				path := parse.Path

				if i := strings.LastIndex(path, "/"); i != -1 { // 路径中有/
					if i != len(path)-1 { // 不是最后一个字符
						//fileName = path[i+1:]
						path = path[:i]
					} else { // 是最后一个字符
						if j := strings.LastIndex(path[:i], "/"); j != -1 {
							//fileName = path[j+1 : len(path)-1]
							path = path[:j]
						}
					}
				}

				fileName := res.GetInfo().GetId() // info 是否事先有签名
				if fileName == "" {
					fileName = sign.HttpBleveIdSign(res)

					res.Info = &HttpStructureStandard.HttpInfo{
						Id: fileName,
					}
				}

				marshal, err := proto.Marshal(res)
				if err != nil {
					return err
				}

				dir := outPath + "/" + parse.Host + "/" + path

				err = os.MkdirAll(dir, os.ModePerm) // 创建目录
				if err != nil {
					return err
				}

				filePath := dir + "/" + fileName + ".httpGroup"

				if _, err := os.Stat(filePath); err == nil { // 文件存在

					continue

				} else { // 文件不存在
					err = os.WriteFile(filePath, marshal, os.ModePerm) // 写入文件
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// BodySign 体签名
func BodySign(b []byte) string {
	var d []byte
	bytes := md5.Sum(b)
	d = append(d, bytes[12:]...)
	return hex.EncodeToString(d)
}
