package video_streaming

import (
	"context"
	"gocv.io/x/gocv"
	"log"
	"sync"
)

type BaseHLSHandler struct {
	config           *HLSConfig
	streamDirManager HLSRepoManager
	frameQ           *[]gocv.Mat
	frameCnt         int
	mutex            sync.Mutex
}

func (h *BaseHLSHandler) HandleFrame(ctx context.Context, frame gocv.Mat) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	(*h.frameQ)[h.frameCnt] = frame
	h.frameCnt++

	if h.frameCnt == h.config.frameNumPerShard {
		currentBatch := *h.frameQ
		go func() {
			log.Println("New HLS shard enqueued")
			err := h.streamDirManager.AddBatch(ctx, currentBatch)
			if err != nil {
				panic(err)
			}
		}()
		newFrameQ := make([]gocv.Mat, h.config.frameNumPerShard)
		h.frameQ = &newFrameQ
		h.frameCnt = 0
	}
}

func NewBaseHLSHandler(manager HLSRepoManager) HLS {
	data := make([]gocv.Mat, manager.GetConfig().frameNumPerShard)
	return &BaseHLSHandler{
		config:           manager.GetConfig(),
		streamDirManager: manager,
		frameQ:           &data,
	}
}
