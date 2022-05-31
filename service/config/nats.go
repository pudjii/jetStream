package config

import (
	"os"
)

var (
	ChannelOrder   = "ORDER.*"
	NATSURL        = os.Getenv("NATSURL")
	StreamName     = "FORDB"
	StreamSubjects = "FORDB.*"
	SubjectName    = "FORDB.subject1"
)
