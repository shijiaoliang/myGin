/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by my-gin.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		status := c.Writer.Status()
		clientIP := c.ClientIP()

		//clientUserAgent := c.Request.UserAgent()
		//referer := c.Request.Referer()
		//hostname, err := os.Hostname()
		//if err != nil {
		//	hostname = "unknow"
		//}

		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		entry := logrus.NewEntry(log).WithFields(logrus.Fields{
			//"hostname":   hostname,
			//"status": status,
			//"latency":    latency, // time to process
			//"clientIP":   clientIP,
			//"method":     c.Request.Method,
			//"path":       path,
			//"referer":    referer,
			//"dataLength": dataLength,
			//"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			//entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			//msg := fmt.Sprintf("%s - %s [%s] \"%s %s\" %d %d \"%s\" \"%s\" (%dms)", clientIP, hostname, time.Now().Format(time.RFC3339), c.Request.Method, path, status, dataLength, referer, clientUserAgent, latency)
			msg := fmt.Sprintf("[access][latency:%v][ip:%s] [method:%s] [path:%s] [status:%d]",
				latency,
				clientIP,
				c.Request.Method,
				path,
				status,
			)

			if status > 499 {
				entry.Error(msg)
			} else if status > 399 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
