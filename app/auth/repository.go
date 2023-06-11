package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zlilemon/gin_auto/pkg/database"
	"github.com/zlilemon/gin_auto/pkg/log"
)

type IRepository interface {
}

type Repository struct {
	dbProxy *gorm.DB
}

var AuthorRepository = new(Repository)

func (r *Repository) GetNumOfAccessToken(c *gin.Context) (num int64, err error) {
	log.Infof("GetNumOfAccessToken ")

	var accessTokenModel []AccessToken

	//num = r.dbProxy.Table("access_token").Find(&accessTokenModel).RowsAffected
	num = database.DB.Table("access_token").Find(&accessTokenModel).RowsAffected

	log.Infof("find rowAffects : %d", num)

	return num, nil
}

func (r *Repository) GetAccessToken(c *gin.Context) (accessToken string, expiresIn int, err error) {
	log.Infof("GetAccessToken ")

	var accessTokenModel []AccessToken

	//num = r.dbProxy.Table("access_token").Find(&accessTokenModel).RowsAffected
	result := database.DB.Table("access_token").Find(&accessTokenModel)
	if result.Error != nil {
		log.Errorf("GetAccessToken error, errMsg:%s", result.Error)
		err = result.Error
		return
	}

	if result.RowsAffected == 0 {
		log.Errorf("GetAccessToken result is NULL")
		err = errors.New("GetAccessToken result is NULL")
		return
	}

	accessToken = accessTokenModel[0].AccessToken
	expiresIn = accessTokenModel[0].ExpiresIn
	log.Infof("GetAccessToken, accessToken: %s, expiresIn:%d", accessToken, expiresIn)

	return
}

func (r *Repository) InsertAccessToken(c *gin.Context, accessToken string, expiresIn int) (err error) {
	log.Infof("enter InsertAccessToken")

	// db中存在记录，更新原有的access_token
	insertSql := "insert into wxapp.access_token set access_token=?, expires_in=?"
	log.Infof("insert accessToken sql : %s", insertSql)

	if error := database.DB.Exec(insertSql, accessToken, expiresIn); error.Error != nil {
		log.Errorf("AccessToken insert to db error, err_smg:%s, error", error.Error)
		return error.Error
	} else {
		log.Infof("insert accessToken success, accessToken:%s, expiresIn:%d", accessToken, expiresIn)
	}

	return nil
}
func (r *Repository) UpdateAccessToken(c *gin.Context, accessToken string, expiresIn int) (err error) {
	log.Infof("enter UpdateAccessToken")

	// db中存在记录，更新原有的access_token
	updateSql := "update wxapp.access_token set access_token=?, expires_in=?"
	log.Infof("update accessToken sql : %s", updateSql)

	if error := database.DB.Exec(updateSql, accessToken, expiresIn); error.Error != nil {
		log.Errorf("AccessToken update to db error, err_smg:%s, error")
		return error.Error
	} else {
		log.Infof("Update accessToken success, accessToken:%s, expiresIn:%d", accessToken, expiresIn)
	}

	return nil
}
