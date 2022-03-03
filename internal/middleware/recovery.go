package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/oddminng/go-blog-service/global"
	"github.com/oddminng/go-blog-service/pkg/app"
	"github.com/oddminng/go-blog-service/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	//defailtMailer := email.NewEmail(&email.SMTPInfo{
	//	Host:     global.EmailSetting.Host,
	//	Port:     global.EmailSetting.Port,
	//	IsSSL:    global.EmailSetting.IsSSL,
	//	UserName: global.EmailSetting.UserName,
	//	Password: global.EmailSetting.Password,
	//	From:     global.EmailSetting.From,
	//})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)

				//err := defailtMailer.SendMail(
				//	global.EmailSetting.To,
				//	fmt.Sprintf("异常抛出, 发生时间: %d", time.Now().Unix()),
				//	fmt.Sprintf("错误信息: %v", err),
				//)
				//if err != nil {
				//	global.Logger.Panicf("mail.SendMail err: %v", err)
				//}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
