/*
 * @Author: dengyue.chen
 * @Date: 2017-03-31 11:37:10
 * @Last Modified by:   dengyue.chen
 * @Last Modified time: 2017-03-31 11:37:10
 */
package main

import (
	"github.com/cdyue/protobuf-demo/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", user.Test)
	r.GET("/test2", user.Test2)

	r.Run(":4000")
}
