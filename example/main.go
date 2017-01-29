package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"unsafe"

	"github.com/lian/snowboy"
	"github.com/xlab/portaudio-go/portaudio"
)

var (
	samplesPerChannel uint    = 512
	sampleRate        float64 = 16000
	channels          int32   = 1
	sampleFormat              = portaudio.PaInt16
)

func main() {
	if err := portaudio.Initialize(); paError(err) {
		log.Fatalln("PortAudio init error:", paErrorText(err))
	}
	defer func() {
		if err := portaudio.Terminate(); paError(err) {
			log.Println("PortAudio term error:", paErrorText(err))
		}
	}()

	l := &Listener{}

	sb := snowboy.NewSnowboyDetect("resources/common.res", "resources/alexa.umdl")
	sb.SetSensitivity("0.5")
	sb.SetAudioGain(1)
	fmt.Println("loaded snowboy", sb)
	defer snowboy.DeleteSnowboyDetect(sb)

	channels = int32(sb.NumChannels())
	sampleRate = float64(sb.SampleRate())
	if sb.BitsPerSample() == 16 {
		sampleFormat = portaudio.PaInt16
	}

	l.Snowboy = sb

	var stream *portaudio.Stream
	if err := portaudio.OpenDefaultStream(&stream, channels, 0, sampleFormat, sampleRate,
		samplesPerChannel, l.paCallback, nil); paError(err) {
		log.Fatalln("PortAudio error:", paErrorText(err))
	}
	defer func() {
		if err := portaudio.CloseStream(stream); paError(err) {
			log.Println("[WARN] PortAudio error:", paErrorText(err))
		}
	}()

	if err := portaudio.StartStream(stream); paError(err) {
		log.Fatalln("PortAudio error:", paErrorText(err))
	}
	defer func() {
		if err := portaudio.StopStream(stream); paError(err) {
			log.Fatalln("[WARN] PortAudio error:", paErrorText(err))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for sig := range c {
		log.Println("exiting", sig)
		os.Exit(0)
	}
}

type Listener struct {
	Snowboy snowboy.SnowboyDetect
}

// paCallback: for simplicity reasons we process raw audio with sphinx in the this stream callback,
// never do that for any serious applications, use a buffered channel instead.
func (l *Listener) paCallback(input unsafe.Pointer, _ unsafe.Pointer, sampleCount uint, _ *portaudio.StreamCallbackTimeInfo, _ portaudio.StreamCallbackFlags, _ unsafe.Pointer) int32 {
	length := int(sampleCount) * int(channels)
	//in := (*(*[1 << 24]int16)(input))[:length]
	//log.Println("data", in)
	res := l.Snowboy.RunDetection(snowboy.SwigcptrInt16_t(input), length)
	if res > 0 {
		fmt.Println("found", res)
	}

	//return int32(portaudio.PaAbort)
	return int32(portaudio.PaContinue)
}

func paError(err portaudio.Error) bool {
	return portaudio.ErrorCode(err) != portaudio.PaNoError
}

func paErrorText(err portaudio.Error) string {
	return portaudio.GetErrorText(err)
}
