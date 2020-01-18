package ginkit

import (
	"fmt"
	"testing"
)

func TestUrlCutQuery(t *testing.T) {
	url := UrlCutQuery("http://1806022.shop.zx-io.cn/pages/asdfasdfasdf?xxx=1111&code=021FVn7201A3FG1yU4620KP4720FVn7v&state=login", "code", "state")
	fmt.Println(url)
}
