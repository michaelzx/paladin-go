package types

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type TimeTest struct {
	T Time
}
type TimeNullTest struct {
	T NullTime
}

func TestTime(t *testing.T) {
	fmt.Printf("%#v", TimeNow())
	fmt.Printf("%#v", Time(time.Now()))
}

func TestTimeFromJson(t *testing.T) {
	str := `{"T":"2019-12-12 17:24:54"}`
	timeTest := TimeTest{}
	err := json.Unmarshal([]byte(str), &timeTest)
	if err != nil {
		println(err)
	}
	fmt.Printf("%#v\n", timeTest)
	fmt.Printf("%#v\n", timeTest.T.Format("2006-01-02 15:04:05"))
}
func TestTimeToJson(t *testing.T) {
	timeTest := TimeTest{
		T: TimeNow(),
	}
	jsonStr, err := json.Marshal(timeTest)
	if err != nil {
		println(err)
	}
	fmt.Println(string(jsonStr))
}
func TestTimeNullFromJson(t *testing.T) {
	str := `{"T":"2019-12-12 17:47:29"}`
	timeTest := TimeNullTest{}
	err := json.Unmarshal([]byte(str), &timeTest)
	if err != nil {
		println(err)
	}
	fmt.Printf("%#v\n", timeTest.T)
}
func TestTimeNullToJson(t *testing.T) {
	timeTest := TimeNullTest{
		T: NewNullTime(time.Now()),
	}
	jsonStr, err := json.Marshal(timeTest)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(jsonStr))
}

func TestZone(t *testing.T) {
	var cstZone = time.FixedZone("CST", 8*3600)
	stdTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2019-12-24 16:48:35", cstZone)
	if err != nil {
		panic(err)
	}
	fmt.Printf(stdTime.String())
}
