package middlewares

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"os"
)

func ConnectDB(dbSession *mgo.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		clonedDBSession := dbSession.Clone()
		c.Set("DB", clonedDBSession.DB(os.Getenv("MONGO_DB")))
		defer clonedDBSession.Close()

		c.Next()
	}
}
