package trans

import (
	"git.yitum.com/saas/shop-admin/model/mysql"
)

// ReqWechatAutoReplyKeywordsList 你可以把ReqWechatAutoReplyKeywordsList嵌套到需要自行修改的结构体中
type ReqWechatAutoReplyKeywordsList struct {
	ReqPage
	mysql.WechatAutoReplyKeywords
}

// ReqWechatAutoReplyKeywordsCreate 你可以把ReqWechatAutoReplyKeywordsCreate或mysql.WechatAutoReplyKeywords嵌套到需要自行修改的结构体中
type ReqWechatAutoReplyKeywordsCreate = mysql.WechatAutoReplyKeywords

// ReqWechatAutoReplyKeywordsUpdate 你可以把ReqWechatAutoReplyKeywordsUpdate或mysql.WechatAutoReplyKeywords嵌套到需要自行修改的结构体中
type ReqWechatAutoReplyKeywordsUpdate = mysql.WechatAutoReplyKeywords

// ReqWechatAutoReplyKeywordsDelete 你可以把ReqWechatAutoReplyKeywordsDelete嵌套到需要自行修改的结构体中
type ReqWechatAutoReplyKeywordsDelete struct {
}
