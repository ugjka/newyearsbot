package nyb

const logChanLen = 100

//LogChan is a channel that sends log messages
type LogChan chan string

func (l LogChan) Write(p []byte) (n int, err error) {
	if len(l) < logChanLen {
		l <- string(p)
	}
	return len(p), nil
}

//NewLogChan make new log channel
func NewLogChan() LogChan {
	return make(chan string, logChanLen)
}
