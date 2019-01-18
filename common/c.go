package common

import (
	"fmt"
	"os"
	"os/exec"
)

func CC() {
	fmt.Println("in the common")
}

type RunData struct {
	Cluster       bool
	ClusterStatus string
}

func (r *RunData) Test() {
	fmt.Println("test")
}

type Util struct {
	Cluster       bool
	ClusterStatus string
}

var Test string

func init() {
	//infoLog.Println("init Util running---------")
	Test = "initial test"
}

func (u *Util) GetCluster() {
	u.ClusterStatus = Test
}

func (u *Util) RunCommand(commandStr string) (string, error) {
	cmdstr := commandStr
	//out, _ := exec.Command("sh", "-c", cmdstr).Output()
	out, res := exec.Command("sh", "-c", cmdstr).CombinedOutput()
	//out, res := exec.Command("sh", "-c", cmdstr).Output()
	if res != nil {
		fmt.Println("show command error res ", res, string(out))
		return "", res
	}
	//fmt.Println("show command res %s", res)
	strout := string(out)
	return strout, nil
}

//background process
func (u *Util) RunCommandB(commandStr string) error {
	//  cmd := exec.Command("./script.sh")
	cmd := exec.Command("sh", "-c", commandStr)
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
	return nil
}
func (u *Util) RunScriptsPath(cmdStr string, chDir string) (string, error) {
	//os.Chdir("/opt/script")
	//commandStr := "bash /opt/script/run.sh"
	commandStr := cmdStr
	cmd := exec.Command("sh", "-c", commandStr)
	//cmd.Dir = "/opt/script"
	cmd.Dir = chDir

	out, err := cmd.Output()
	//      err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	//infoLog.Println(out)
	fmt.Println("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
	return string(out), nil
}
