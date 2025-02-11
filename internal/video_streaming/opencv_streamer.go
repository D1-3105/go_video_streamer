package video_streaming

import (
	"errors"
	"go_video_streamer/internal/opencv_global_capture"
	"gocv.io/x/gocv"
	"sync"
)

type CaptureError struct {
	Err error
}

func (e CaptureError) Error() string {
	return e.Err.Error()
}

type CaptureStreamer struct {
	capture      *opencv_global_capture.VideoCapture
	captureMutex sync.RWMutex
}

func NewCaptureStreamer(capture *opencv_global_capture.VideoCapture) CaptureStreamer {
	return CaptureStreamer{capture: capture}
}

func (streamer *CaptureStreamer) GetFrame() (gocv.Mat, error) {
	streamer.captureMutex.RLock()
	defer streamer.captureMutex.RUnlock()

	img := gocv.NewMat()
	err := !(*streamer.capture).Read(&img)
	if err {
		return img, CaptureError{Err: errors.New("failed to read frame")}
	}
	return img, nil
}

func (streamer *CaptureStreamer) GetVideoFPS() float64 {
	return (*streamer.capture).Get(gocv.VideoCaptureFPS)
}
