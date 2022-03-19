/*
 * @Author: your name
 * @Date: 2022-03-19 22:24:53
 * @LastEditTime: 2022-03-19 22:25:08
 * @LastEditors: your name
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Quasar-IM-backend\IM_backend\service\model.go
 */
package service

type Trainer struct {
	Content   string `bson:"content"`   // 内容
	StartTime int64  `bson:"startTime"` // 创建时间
	EndTime   int64  `bson:"endTime"`   // 过期时间
	Read      uint   `bson:"read"`      // 已读
}

type Result struct {
	StartTime int64
	Msg       string
	Content   interface{}
	From      string
}
