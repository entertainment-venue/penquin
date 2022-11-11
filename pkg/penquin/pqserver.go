package penquin

import (
	"github.com/gin-gonic/gin"
)

type PQServer struct {
	ginEngine *gin.Engine
	pq        *PenQuin
}

func NewPQServer(pq *PenQuin) PQServer {
	r := gin.New()

	// Ping
	r.GET("/ping", pq.pingHandler)

	// Config
	r.GET("/config", pq.configHandler)

	// Metadata
	// Topic
	tr := r.Group("/topic/:topic", pq.topicSearchHandler)
	{
		tr.POST("/create", pq.topicCreateHandler)
		tr.POST("/delete", pq.topicDeleteHandler)
		tr.POST("/update", pq.topicUpdateHandler)
	}
	// Shard
	sr := r.Group("/shard/:topic/:shard", pq.shardSearchHandler)
	{
		sr.POST("/create", pq.shardCreateHandler)
		sr.POST("/delete", pq.shardDeleteHandler)
		sr.POST("/update", pq.shardUpdateHandler)
	}

	// Produce
	r.POST("/produce", pq.produceHandler)
	r.POST("/mproduce", pq.mproduceHandler)

	return PQServer{
		ginEngine: r,
		pq:        pq,
	}
}

func (pq *PQServer) serve() error {
	if err := pq.ginEngine.Run(); err != nil {
		return err
	}
	return nil
}
