package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/joho/godotenv"
)

type Env struct {
	DbUrl string
}

type Args struct {
	up   *bool
	down *bool
	path *string
}

func RunCommand(cmd *exec.Cmd) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error creating StdoutPipe: %v\n", err)
		os.Exit(1)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Error creating StderrPipe: %v\n", err)
		os.Exit(1)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error starting command: %v\n", err)
		os.Exit(1)
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Error running goose command: %v\n", err.Error())
		os.Exit(1)
		return
	}
}

func LoadStdout(cmd *exec.Cmd) *io.ReadCloser {
	var stdout, err = cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error creating StdoutPipe: %v\n", err)
		os.Exit(1)
		return nil
	}
	return &stdout
}

func LoadStderr(cmd *exec.Cmd) *io.ReadCloser {
	var stderr, err = cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Error creating StderrPipe: %v\n", err)
		os.Exit(1)
		return nil
	}
	return &stderr
}

func LoadEnv() *Env {
	var err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file. Please check if you have one")
		os.Exit(1)
		return nil
	}
	return &Env{DbUrl: os.Getenv("DB_URL")}
}

func ParseArgs() *Args {
	var parser = argparse.NewParser("print", "Prints provided string to stdout")
	var args = Args{
		up:   parser.Flag("u", "up", nil),
		down: parser.Flag("d", "down", nil),
		path: parser.StringPositional(&argparse.Options{Required: true, Help: "You must provide the path to the migrations folder"}),
	}

	var err = parser.Parse(os.Args)
	if err != nil {
		fmt.Printf("Error parsing the algs: %v\n", err.Error())
		os.Exit(1)
		return nil
	}

	return &args
}

func main() {
	var env = LoadEnv()
	var args = ParseArgs()

	var cmd *exec.Cmd
	var commandPath = []string{"goose", "-dir", *args.path, "postgres", (*env).DbUrl}
	if *args.up {
		commandPath = append(commandPath, "up")
		cmd = exec.Command(commandPath[0], commandPath[1:]...)
	} else {
		commandPath = append(commandPath, "down")
		cmd = exec.Command(commandPath[0], commandPath[1:]...)
	}
	fmt.Printf("Running the command: %v\n", strings.Join(commandPath, " "))

	RunCommand(cmd)
	fmt.Println("Command ran successfully")
}
