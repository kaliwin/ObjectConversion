package proxyImport

import (
	"github.com/kaliwin/Needle/crypto/certificate"
	http3 "github.com/kaliwin/Needle/network/http"
	"net/http"
)

// 纯http代理方式导入

// HttpProxyImport http 代理导入
type HttpProxyImport struct {
	// 代理地址
	TarGetProxy http3.HttpClient // 目标代理地址
	// 原始数据目录
	HttpGroupPath string // 原始数据目录 目录文件必须为httpGroup 只支持integrate的onlyHost格式目录 一个host的所有请求放在一个文件夹下
	// 不允许使用path分成 那样会导致索引丢失

	CACert certificate.CACert // CA证书 用于签发证书 通常burp可以设置信任所有证书
	//Client http3.HttpClient   // http客户端
}

// ServeHTTP http代理服务器 处理代理请求
func (h *HttpProxyImport) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 根据请求方法分发处理函数
	if request.Method == http.MethodConnect {
		h.handleHTTPS(writer, request)
	} else {
		h.handleHTTP(writer, request)
	}
}

// handleHTTP http的处理
func (h *HttpProxyImport) handleHTTP(w http.ResponseWriter, r *http.Request) {
	//hijacker, _ := w.(http.Hijacker)
	//hijack, _, err := hijacker.Hijack() // 获取底层连接
	//if err != nil {
	//	return
	//}
	//r.Header.Get("Set-Cookie2")

}

// handleHTTPS https的处理
func (h *HttpProxyImport) handleHTTPS(w http.ResponseWriter, r *http.Request) {

}

// Go 启动http代理导入
func (h *HttpProxyImport) Go() {

}
