package nyb

import (
	"errors"
	"testing"
	"time"
)

func TestLoopTimeZones(t *testing.T) {
	nye := New("", []string{""}, "hny", "", false, "", "")
	nye.decodeZones(Zones)
	timeNow = func() time.Time {
		return target.Add(11*time.Hour + time.Minute*59 + time.Second*59)
	}
	nye.loopTimeZones()
	//Test stopper
	close(nye.Stopper)
	nye.loopTimeZones()
	//Test fatal error
	nye.Stopper = make(chan bool)
	nye.zones[0].Offset = "aeiiaoeii"
	nye.loopTimeZones()
}

func TestIrcControl(t *testing.T) {
	reconnectInterval = time.Millisecond * 250
	pingInterval = time.Millisecond * 25
	nye := New("", []string{""}, "hny", "", false, "", "")
	close(nye.Stopper)
	nye.Add(1)
	nye.ircControl()
	nye.Wait()
	nye.Stopper = make(chan bool)
	go func() {
		nye.Bot.Errchan <- errors.New("test error")
		time.Sleep(time.Second / 2)
		nye.pp <- true
		time.Sleep(time.Second / 2)
		close(nye.Stopper)
	}()
	nye.Add(1)
	nye.ircControl()
	nye.Wait()
}

func TestStop(t *testing.T) {
	nye := New("", []string{""}, "hny", "", false, "", "")
	nye.Stop()
	nye.Stop()
}

func TestDecodeZones(t *testing.T) {
	nye := New("", []string{""}, "hny", "", false, "", "")
	nye.decodeZones(Zones)
	zones := "*" + Zones
	nye.decodeZones(zones)
}

func TestStart(t *testing.T) {
	timeNow = func() time.Time {
		return target.Add(11*time.Hour + time.Minute*59 + time.Second*59 + time.Millisecond*500)
	}
	nye := New("", []string{""}, "hny", "", false, "", "")
	nye.Bot.DebugFakeConn = true
	close(nye.Stopper)
	nye.Start()
	nye.Stopper = make(chan bool)
	close(nye.start)
	go func() {
		time.Sleep(time.Second * 4)
		close(nye.Stopper)
	}()
	nye.Start()
}
