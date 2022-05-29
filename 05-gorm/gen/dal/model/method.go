package model

type Method interface {
	// sql(SELECT COUNT(version) FROM user_user GROUP BY version ORDER BY version DESC LIMIT 1)
	GetMaxVersionCount() (int, error)
}
