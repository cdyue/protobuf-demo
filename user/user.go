package user

import (
	mu "github.com/cdyue/protobuf-demo/model/protobuf/user"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	log "github.com/inconshreveable/log15"
)

//Test test
func Test(c *gin.Context) {
	info := &mu.User{
		Id:         "001",
		Name:       "dengyue",
		Role:       []string{"GOD", "ADMIN"},
		UserType:   "User",
		UserStatus: "Normal",
	}
	log.Info("Start", "User", info)

	data, err := proto.Marshal(info)
	if err != nil {
		log.Error("marshal error", "msg", err.Error())
		c.JSON(500, nil)
		return
	}
	log.Info("print", "data", data)
	c.JSON(200, data)
}

//Test2 test
func Test2(c *gin.Context) {
	info := &mu.User{
		Id:         "001",
		Name:       "dengyue",
		Role:       []string{"GOD", "ADMIN"},
		UserType:   "User",
		UserStatus: "Normal",
	}
	log.Info("Start", "User", info)
	data, err := proto.Marshal(info)
	if err != nil {
		log.Error("marshal error", "msg", err.Error())
		c.JSON(500, nil)
		return
	}
	log.Info("print", "data", data)

	newdata := &mu.User{}
	if err = proto.Unmarshal(data, newdata); err != nil {
		log.Error("json marshal error", "msg", err.Error())
		c.JSON(500, nil)
		return
	}

	c.JSON(200, newdata)
}
