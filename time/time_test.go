package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Log(int64(time.Second / time.Millisecond))
}

var c chan int

func handel(int) {}
func TestAfter(t *testing.T) {
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("time out")
	}
	//select {
	//case m := <-c:
	//	handel(m)
	//case <-time.After(10 * time.Second):
	//	fmt.Println("time out")
	//}
}

func TestSleep(t *testing.T) {
	t.Log("time sleep start")
	time.Sleep(1 * time.Second)
	t.Log("time sleep end")
}

//func TestTick(t *testing.T) {
//	t.Log("time tick start")
//	c := time.Tick(5 * time.Second)
//	for next := range c {
//		t.Logf("%v", next)
//	}
//	t.Log("time tick end")
//}

func TestParseDuration(t *testing.T) {
	hours, _ := time.ParseDuration("10h")
	t.Log(hours, hours.Hours())

	complex, _ := time.ParseDuration("1h10m10s")
	t.Log(complex, complex.Hours(), complex.Minutes(), complex.Seconds(), complex.String())

}

func TestParseInLocation(t *testing.T) {

}

//https://studygolang.com/articles/7134
func TestSince(t *testing.T) {
	t1 := time.Now() // get current time
	//logic handlers
	for i := 0; i < 1000; i++ {
		fmt.Print("*")
	}
	elapsed := time.Since(t1)
	t.Log("App elapsed: ", elapsed) //App elapsed:  753µs
}

func TestLocation(t *testing.T) {
	// China doesn't have daylight saving. It uses a fixed 8 hour offset from UTC.
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// If the system has a timezone database present, it's possible to load a location
	// from that, e.g.:
	//    newYork, err := time.LoadLocation("America/New_York")

	// Creating a time requires a location. Common locations are time.Local and time.UTC.
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)

	// Although the UTC clock time is 1200 and the Beijing clock time is 2000, Beijing is
	// 8 hours ahead so the two dates actually represent the same instant.
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

}

func TestDate(t *testing.T) {
	ti := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	t.Logf("Go launched at %s\n", ti.Local())

	ti = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.Local)
	t.Logf("Go launched at %s\n", ti.Local())

	t.Log(time.Now())
}

//定时器1 每1秒一次
func TestNewTimer(t *testing.T) {
	t.Log("time  new timer start")
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()
	for {
		<-timer.C
		t.Log("1 秒 到了")
		timer.Reset(time.Second) //设置下一个循环
	}
}

//定时器2 每1秒一次
func TestNewTicker(t *testing.T) {
	//1
	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()
	for {
		<-ticker.C
		t.Log("1 秒 到了")
	}
	//2
	ticker2 := time.NewTicker(time.Second * 1)
	defer ticker2.Stop()
	for range ticker2.C {
		t.Log("1 秒 到了")
	}
}

//定时器3 每1秒一次 配合 time.After
