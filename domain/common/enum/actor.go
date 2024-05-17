package commonenum

type Actor string

var (
	ACTOR_SYSTEM Actor = "SYSTEM"
	ACTOR_ADMIN  Actor = "ADMIN"
	ACTOR_USER   Actor = "USER"
)

func (a Actor) String() string {
	return string(a)
}
