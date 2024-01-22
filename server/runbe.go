package server

import (
	"io"
	"os"
	"os/exec"
	"runtime"
)

var SERVER *Server

type Server struct {
	Cmd    *exec.Cmd
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
}

func NewServer(path string) *Server {
	s := new(Server)

	SERVER = s

	if runtime.GOOS == "windows" {
		s.Cmd = exec.Command("Cmd.exe", "/c", "chcp 65001 &&", path)
	} else {
		s.Cmd = exec.Command("sh", "-c", path)
	}

	// s.Cmd.Stdout = os.Stdout
	s.Cmd.Stderr = os.Stderr

	stdin, err := s.Cmd.StdinPipe()
	if err != nil {
		return nil
	}
	s.Stdin = stdin

	stdout, err := s.Cmd.StdoutPipe()
	if err != nil {
		return nil
	}
	s.Stdout = stdout

	err = s.Cmd.Start()
	if err != nil {
		return nil
	}

	return s
}
