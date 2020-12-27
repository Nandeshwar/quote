package repo

import (
	"context"
	"github.com/sirupsen/logrus"
	"quote/pkg/constants"
	"quote/pkg/model"
	"time"
)

type IEmailStatusRepo interface {
	CreateEmailStatus(ctx context.Context, emailStatus model.EmailStatusGORM) (int64, error)
	EmailSentForEvents(ctx context.Context, sentAt time.Time, typ string) bool
}

func (s SQLite3Repo) CreateEmailStatus(ctx context.Context, emailStatus model.EmailStatusGORM) (int64, error) {
	tx := s.GORMDB.WithContext(ctx).Create(&emailStatus).Debug()

	if tx.Error != nil {
		return 0, tx.Error
	}

	return emailStatus.ID, nil
}

func (s SQLite3Repo) EmailSentForEvents(ctx context.Context, sentAt time.Time, typ string) bool {
	var emailStatus model.EmailStatusGORM
	err := s.GORMDB.WithContext(ctx).
		Where("type = ? and status = ? and DATE(sent_at)",
			typ,
			"sent",
			sentAt.Format(constants.DATE_FORMAT_EVENT_DATE)).Debug().First(&emailStatus).Error

	if err != nil {
		logrus.Errorf("error=%v", err)
		return false
	}

	return true
}
