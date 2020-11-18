package utils

import (
	"os/exec"
	"bytes"
	"os"
	"fmt"
	"io"
	"sync"
	"log"
	"encoding/json"
	"path/filepath"
	"strings"
)

func NiceBuffRunner(cmd *exec.Cmd, workdir string) (string, string) {
	var stdoutBuf, stderrBuf bytes.Buffer
	//newEnv := append(os.Environ(), ENV...)
	//cmd.Env = newEnv
	cmd.Dir = workdir
	pipe, _ := cmd.StdoutPipe()
	errpipe, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		return fmt.Sprintf("%s", os.Stdout), fmt.Sprintf("%s", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, pipe)
		wg.Done()
	}()
	go func() {
		_, errStderr = io.Copy(stderr, errpipe)
		wg.Wait()
	}()
	err = cmd.Wait()
	if err != nil {
		return fmt.Sprintf("%s", os.Stdout), fmt.Sprintf("%s", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("Command runninng error: failed to capture stdout or stderr\n")
	}
	return stdoutBuf.String(), stderrBuf.String()
}

func SimpleQuietRunner(cmdtorun []string, workdir string) (string, error){
	cmd := exec.Command(cmdtorun[0], cmdtorun[1:]...)
	cmd.Dir = workdir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Error while executing the command:\n%+v\n%v", cmdtorun, err)
	}
	return fmt.Sprintf("%s", string(out)), nil
}

func (config *VilyaCfg) SetupConfig() error {
	var configfilepath string
	for _, val := range os.Environ() {
		if strings.Contains(val, "VILYAROOT=") {
			configfilepath = strings.Replace(val, "VILYAROOT=", "", 1)
			if configfilepath == "" {
				log.Fatalf("ERROR! Please check your os.Env[VILYAROOT=]\n")

			}
		}
	}
	workdir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error getting working dir: %v", err)
	}
	out1, err := exec.Command("whoami").CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error running whoami %v", err)
	}
	whoami := strings.Replace(fmt.Sprintf("%s", string(out1)), "\n", "", 100)
	homefoldslice := []string{workdir, filepath.Join("/home", whoami), filepath.Join("/home", whoami, "vilya"), filepath.Join("/home", whoami, "go/src/vilya"),
		filepath.Join("/home", whoami, "golang/src/vilya", "vilyaCfg.json")}
	for _, value := range homefoldslice {
		cmdargs := []string{"ls", "-alh", value}
		out2, err := exec.Command(cmdargs[0], cmdargs[1:]...).CombinedOutput()
		if err != nil {
			return fmt.Errorf("Error executing ls... %v", err)
		}
		if strings.Contains(fmt.Sprintf("%s", string(out2)), "vilyaCfg.json") {
			configfilepath = filepath.Join(value, "vilyaCfg.json")
			break
		}
	}
	if configfilepath == "" {
		return fmt.Errorf("ERROR!...\nPlease indicate the config file path")
	} else {
		log.Printf("configfilepath: %s\n", configfilepath)
	}
	f, err := os.Open(configfilepath)
	defer f.Close()
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	if err := json.NewDecoder(f).Decode(&config); err != nil {
		log.Printf("Error: %s\n", err)
		return  err
	}
	return nil
}