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
func TestGetAudioType(t *testing.T) {
	getAudioHex(t, "/Users/michael/Downloads/6b7e_5c9f_2214_f3bb176386e9ef79fff62f3f69dca2f7.mp3")
	getAudioHex(t, "/Users/michael/Downloads/Ensoniq-ESQ-1-Sympy-C4.wav")
}
func getAudioHex(t *testing.T, fp string) {
	fSrc, _ := readFile(t, fp)
	fileCode := bytesToHexString(fSrc[:20])
	t.Log(fileCode)
	t.Log(GetFileType(fSrc[:20]))
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

func readFile(t *testing.T, fp string) ([]byte, error) {
	f, err := os.Open(fp)
	if err != nil {
		t.Logf("open error: %v", err)
	}

	return ioutil.ReadAll(f)
}

func TestFileIsWav(t *testing.T) {
	wavFilePath := "/Users/michael/Downloads/Ensoniq-ESQ-1-Sympy-C4.wav"
	fSrc, _ := readFile(t, wavFilePath)
	t.Log(FileIsWav(fSrc[:20]))
}

func TestFileIsMp3(t *testing.T) {
	wavFilePath := "/Users/michael/Downloads/6b7e_5c9f_2214_f3bb176386e9ef79fff62f3f69dca2f7.mp3"
	fSrc, _ := readFile(t, wavFilePath)
	t.Log(FileIsMp3(fSrc[:20]))
}
