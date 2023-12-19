package mysql

import (
	"bluebell1/models"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id,community_name from community`
	if err = db.Select(&communityList, sqlStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetail(id int64) (communitydetail *models.CommunityDetail, err error) {
	communitydetail = new(models.CommunityDetail)
	sqlStr := `select community_id,community_name,introduction,creat_time from community where community_id=?`
	if err = db.Get(communitydetail, sqlStr, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = InvalidID
		}
	}
	return
}
