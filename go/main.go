/*
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-04 13:30:57
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 14:27:55
 * @FilePath: \checkIn\go\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"checkin/config"
	"checkin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.UserRouterInit(r)
	r.Run(":" + config.Application.Server.Port)
}
