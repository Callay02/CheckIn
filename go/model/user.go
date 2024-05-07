/*
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-04 13:43:37
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 22:18:06
 * @FilePath: \checkIn\go\entity\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	UpdateTime string `json:"updateTime"`
	CreateTime string `json:"createTime"`
	State      int    `json:"state"`
}

func (User) TableName() string {
	return "users"
}
