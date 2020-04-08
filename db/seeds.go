package db

// FillSeedsInformation <function>
// is used to fill or delete information that needed before usage of news endpoint
func (s Service) FillSeedsInformation() {
	s.DeleteOldData()
}
