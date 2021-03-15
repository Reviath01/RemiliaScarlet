import (
	_ "../commands"
	"../state"
)

func init() {
    state.Multiplexer.SessionRegisterHandlers(state.RawSession)
}
