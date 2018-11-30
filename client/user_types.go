// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "computingProvider service APIs": Application User Types
//
// Command:
// $ goagen
// --design=github.com/ZJU-DistributedAI/ComputingProvider/design
// --out=$(GOPATH)/src/github.com/ZJU-DistributedAI/ComputingProvider
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
)

// filePayload user type.
type filePayload struct {
	// file
	File *string `form:"file,omitempty" json:"file,omitempty" yaml:"file,omitempty" xml:"file,omitempty"`
}

// Validate validates the filePayload type instance.
func (ut *filePayload) Validate() (err error) {
	if ut.File == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "file"))
	}
	return
}

// Publicize creates FilePayload from filePayload
func (ut *filePayload) Publicize() *FilePayload {
	var pub FilePayload
	if ut.File != nil {
		pub.File = *ut.File
	}
	return &pub
}

// FilePayload user type.
type FilePayload struct {
	// file
	File string `form:"file" json:"file" yaml:"file" xml:"file"`
}

// Validate validates the FilePayload type instance.
func (ut *FilePayload) Validate() (err error) {
	if ut.File == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "file"))
	}
	return
}
