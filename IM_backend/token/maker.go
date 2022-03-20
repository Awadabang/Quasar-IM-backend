/*
 * @Author: your name
 * @Date: 2022-03-09 20:57:16
 * @LastEditTime: 2022-03-20 15:56:06
 * @LastEditors: your name
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Quasar-IM-backend\IM_backend\token\maker.go
 */
package token

import "time"

// Maker is
type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(id int64, username string, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
