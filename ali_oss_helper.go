package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

const (
	_BUCKET = "app-build"
	_DOMAIN = ""
)

type FILEINFO struct {
	FILENAME string `json:"filename"`
	FILESIZE int64  `json:"filesize"`
	FILEURL  string `json:"fileurl"`
}

func getOSS() *oss.Bucket {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket所在地域对应的Endpoint。以华东1（杭州）为例，Endpoint填写为https://oss-cn-hangzhou.aliyuncs.com。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	aki, i_err := os.LookupEnv("AKI")
	aks, s_err := os.LookupEnv("AKS")
	fmt.Printf("err:", i_err, s_err)

	client, err := oss.New("oss-cn-hangzhou.aliyuncs.com", aki, aks)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写Bucket名称，例如examplebucket。
	bucket, err := client.Bucket("app-build")
	if err != nil {
		fmt.Printf("Error:", err)
		os.Exit(-1)
	}
	return bucket

}

func singURL(prefix string) string {
	// 依次填写Object完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径(例如D:\\localpath\\examplefile.txt)。Object完整路径中不能包含Bucket名称。
	singUrl, err := getOSS().SignURL(prefix, oss.HTTPGet, 6000)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("Sign Url:%s\n", singUrl)
	return singUrl

}

func GetFileList(prefixs string) []FILEINFO {
	var fileList []FILEINFO

	// 遍历文件。
	marker := oss.Marker("")
	prefix := oss.Prefix(prefixs)
	for {
		lor, err := getOSS().ListObjects(marker, prefix)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		for _, object := range lor.Objects {
			fmt.Println(singURL(object.Key), object.Type, object.Size)
			fileList = append(fileList, FILEINFO{object.Key, object.Size, singURL(object.Key)})
		}
		if lor.IsTruncated {
			prefix = oss.Prefix(lor.Prefix)
			marker = oss.Marker(lor.NextMarker)
		} else {
			break
		}
	}

	return fileList
}

func FileInfo() {

}
