package webserver

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/govice/golinks-daemon/pkg/authentication"
	"github.com/govice/golinks/block"
)

func (w *Webserver) externalAuthenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAuth := &authentication.ExternalUserAuth{
			Token: c.Query("token"),
			Email: c.Query("email"),
		}

		ok, err := w.servicer.AuthenticationService().Valid(userAuth)
		if (err == nil) && ok {
			c.Next()
			return
		}

		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "not authorized",
		})

		c.Abort()
	}
}

func (w *Webserver) postBlockEndpoint(c *gin.Context) {
	body, _ := c.GetRawData()
	var data blockData
	if err := json.Unmarshal(body, &data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "recieved invalid block data",
		})
		return
	}

	block, err := w.servicer.BlockchainService().AddBlock([]byte(data.Data))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error adding block to chain",
		})
		return
	}

	c.PureJSON(http.StatusOK, block)
}

func (w *Webserver) findBlockEndpoint(c *gin.Context) {
	//todo find a way to pass raw bytes in parameter or migrate to body request
	body, _ := c.GetRawData()
	var finder blockchainSearch
	if err := json.Unmarshal(body, &finder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}

	block := &block.Block{}
	switch finder.Format {
	case "index":
		index, _ := strconv.Atoi(finder.Key)
		block, _ = w.servicer.BlockchainService().FindBlockByIndex(index)
		break
	case "hash":
		hash, _ := base64.StdEncoding.DecodeString(finder.Key)
		block, _ = w.servicer.BlockchainService().FindBlockByHash(hash)
		break
	case "parent_hash":
		parentHash, _ := base64.StdEncoding.DecodeString(finder.Key)
		block, _ = w.servicer.BlockchainService().FindBlockByParentHash(parentHash)
		break
	case "timestamp":
		timestamp, _ := strconv.ParseInt(finder.Key, 10, 64)
		block, _ = w.servicer.BlockchainService().FindBlockByTimestamp(timestamp)
		break
	}

	if block != nil {
		c.PureJSON(http.StatusOK, block)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Block not found",
		})
	}
}

func (w *Webserver) getChainEndpoint(c *gin.Context) {
	c.PureJSON(http.StatusOK, w.servicer.BlockchainService().Chain())
}

type blockchainSearch struct {
	Format string `json:"format"`
	Key    string `json:"key"`
}

type blockData struct {
	Data []byte `json:"data"`
}
