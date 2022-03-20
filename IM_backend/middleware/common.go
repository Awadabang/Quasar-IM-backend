/*
 * @Author: your name
 * @Date: 2022-03-20 15:07:37
 * @LastEditTime: 2022-03-20 15:07:39
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Quasar-IM-backend\IM_backend\middleware\common.go
 */
//返回错误信息 ErrorResponse
package middleware

import "github.com/gin-gonic/gin"

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
