package enum

type Topic string

const (
	TopicUserRegistrations  Topic = "user_registrations"
	TopicUserProfileUpdates Topic = "user_profile_updates"
)

func (t Topic) String() string {
	return string(t)
}
