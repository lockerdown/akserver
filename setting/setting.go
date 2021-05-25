package setting

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

var (
	ServerMode    string
	AkServerAddr  string
	AkServerPort  string
	JwtKey        string
	PrivateKey    string
	PublicKey     string
	TlsPublicKey  string
	TlsPrivateKey string
)

func init() {
	file, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadPem(file)
	GenerateTlsCert(2048)
	GenerateRSAKey(2048)
}

func LoadServer(file *ini.File) {
	ServerMode = file.Section("server").Key("AppMode").MustString("debug")
	AkServerAddr = file.Section("server").Key("AkServerAddr").MustString("0.0.0.0")
	AkServerPort = file.Section("server").Key("AkServerPort").MustString("30892")
	JwtKey = file.Section("server").Key("JwtKey").MustString("8gjs8s5s72")
}

func LoadPem(file *ini.File) {
	TlsPublicKey = file.Section("Tls").Key("PublicKey").MustString("tlsPublic.pem")
	TlsPrivateKey = file.Section("Tls").Key("PrivateKey").MustString("tlsPrivate.pem")
}

//GenerateTlsCert 生成默认证书
func GenerateTlsCert(bits int) {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	// 定义：引用IETF的安全领域的公钥基础实施（PKIX）工作组的标准实例化内容
	subject := pkix.Name{
		CommonName: "akserver",
	}

	// 设置 SSL证书的属性用途
	certificate509 := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(100 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	// 生成指定位数密匙
	pk, _ := rsa.GenerateKey(rand.Reader, bits)

	// 生成 SSL公匙
	derBytes, _ := x509.CreateCertificate(rand.Reader, &certificate509, &certificate509, &pk.PublicKey, pk)
	certOut, _ := os.Create(TlsPublicKey)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// 生成 SSL私匙
	keyOut, _ := os.Create(TlsPrivateKey)
	pem.Encode(keyOut, &pem.Block{Type: "RAS PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}

//GenerateRSAKey rsa加密使用
func GenerateRSAKey(bits int) {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	var privateKey, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile, &privateBlock)

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	pem.Encode(publicFile, &publicBlock)
}
