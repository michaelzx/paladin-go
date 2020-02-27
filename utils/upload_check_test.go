package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetFileType(t *testing.T) {

	getHex(t, "/Users/michael/workspace/personal/201911=到店易/图片类型识别问题/9.jpg")
	getHex(t, "/Users/michael/workspace/personal/201911=到店易/图片类型识别问题/WechatIMG17713.jpeg")
	getHex(t, "/Users/michael/workspace/personal/201911=到店易/图片类型识别问题/大图片/儿童版大黄蜂.jpg")
	getHex(t, "/Users/michael/workspace/personal/201911=到店易/图片类型识别问题/gps.jpeg")
	getHex(t, "/Users/michael/workspace/personal/201911=到店易/图片类型识别问题/bt3.gif")
}
func TestGetVideoType(t *testing.T) {
	getVideoHex(t, "/Users/michael/Downloads/QQ20200222-202435-HD.mp4")
	getVideoHex(t, "/Users/michael/Downloads/QQ20200222-202520-HD.mp4")
}
func getHex(t *testing.T, fp string) {
	// f, err := os.Open("C:\\Users\\Administrator\\Desktop\\api.html")
	f, err := os.Open(fp)
	if err != nil {
		t.Logf("open error: %v", err)
	}

	fSrc, err := ioutil.ReadAll(f)
	fileCode := bytesToHexString(fSrc[:10])
	t.Log(fileCode)
	t.Log(GetFileType(fSrc[:10]))
}
func getVideoHex(t *testing.T, fp string) {
	f, err := os.Open(fp)
	if err != nil {
		t.Logf("open error: %v", err)
	}

	fSrc, err := ioutil.ReadAll(f)
	t.Log(FileIsMp4(fSrc[:20]))
}
