package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/lian/snowboy"
	"github.com/lian/snowboy/openal"
)

func main() {
	fmt.Println("openal test")

	sb := snowboy.NewSnowboyDetect("resources/common.res", "resources/alexa.umdl")
	sb.SetSensitivity("0.5")
	sb.SetAudioGain(1)
	fmt.Println("loaded snowboy", sb)
	defer snowboy.DeleteSnowboyDetect(sb)

	var format openal.Format
	frequency := uint32(sb.SampleRate())
	captureSize := 512 * 4
	captureSizeMax := captureSize * 2

	if sb.BitsPerSample() == 16 && sb.NumChannels() == 1 {
		format = openal.FormatMono16
	} else {
		panic("snowboy uses more than 1 channel")
	}

	mic := openal.CaptureOpenDevice("", frequency, format, uint32(captureSizeMax*format.SampleSize()))
	mic.CaptureStart()
	defer mic.CloseDevice()

	samples := make([]int16, captureSizeMax)
	openalSamplesPointer := unsafe.Pointer(&samples[0])
	snowboySamplesPointer := snowboy.SwigcptrInt16_t(openalSamplesPointer)

	ticker := time.NewTicker(time.Millisecond * 100)

	for _ = range ticker.C {
		numSamples := int(mic.CapturedSamples())
		if numSamples >= captureSize {
			mic.CaptureToInt16Pointer(openalSamplesPointer, numSamples)
			res := sb.RunDetection(snowboySamplesPointer, numSamples)
			if res > 0 {
				fmt.Println("found", res)
			}
		}
	}
}
