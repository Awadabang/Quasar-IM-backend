/*
 * @Author: your name
 * @Date: 2022-03-20 15:40:36
 * @LastEditTime: 2022-03-20 16:15:04
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \Quasar-IM-backend\IM_backend\api\friend_about_model.go
 */
package api

type addfriendRequest struct {
	Friendid int64 `json:"friendid"`
}

type getfriendsRequest struct {
	Page_id   int32 `form:"page_id" binding:"required, min=0"`
	Page_size int32 `form:"page_size" binding:"required, min=1"`
}
