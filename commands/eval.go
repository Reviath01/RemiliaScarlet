package commands

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"github.com/google/uuid"
)

// EvalCommand is a handler for eval command
func EvalCommand(ctx CommandHandler.Context, _ []string) error {
	uuidvar := uuid.New()
	_, err := os.Stat("./evals")
	if os.IsNotExist(err) {
		err = os.Mkdir("./evals", os.ModeDevice)
		if err != nil {
			ctx.Reply("Error while creating evals folder")
			time.Sleep(1 * time.Second)
			ctx.Reply(err.Error())
			return nil
		}
	}

	file, err := os.Create(fmt.Sprintf("./evals/%s.go", uuidvar))
	if err != nil {
		ctx.Reply("Error while creating file..")
		time.Sleep(1 * time.Second)
		ctx.Reply(err.Error())
		return nil
	}
	defer file.Close()
	ctx.Reply(fmt.Sprintf("Created file %s.go", uuidvar))
	err = ioutil.WriteFile(fmt.Sprintf("./evals/%s.go", uuidvar), []byte(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")), fs.ModeDevice)
	if err != nil {
		ctx.Reply("Cannot write to file.")
		time.Sleep(1 * time.Second)
		ctx.Reply(err.Error())
		return nil
	}
	ctx.Reply(fmt.Sprintf("Wrote code to %s.go, trying to run...", uuidvar))
	cmd := exec.Command("go", "run", fmt.Sprintf("./evals/%s.go", uuidvar))
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err = cmd.Run()
	if err != nil {
		ctx.Reply(err.Error())
		return nil
	}
	ctx.Reply("Output:")
	time.Sleep(1 * time.Second)
	ctx.Reply(outb.String())
	return err
}
