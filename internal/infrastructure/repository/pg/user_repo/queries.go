package user_repo

import _ "embed"

var (
	//go:embed sql/create_user.sql
	queryCreateUser string

	//go:embed sql/update_user.sql
	queryUpdateUser string

	//go:embed sql/fetch_users.sql
	queryFetchUsers string

	//go:embed sql/get_user_by_id.sql
	queryGetUserByID string

	//go:embed sql/lock_user_for_update.sql
	queryLockUserForUpdate string
)
