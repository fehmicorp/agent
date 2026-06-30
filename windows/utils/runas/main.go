package runas

func RequireAdministrator() error {
	if IsAdmin() {
		return nil
	}
	return RequestElevation()
}
