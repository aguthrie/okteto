// Copyright 2022 The Okteto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"

	sp "github.com/briandowns/spinner"
	oktetoLog "github.com/okteto/okteto/pkg/log"
	"github.com/okteto/okteto/pkg/model"
	"golang.org/x/term"
)

var spinnerSupport bool

//Spinner represents an okteto spinner
type Spinner struct {
	sp *sp.Spinner
}

//NewSpinner returns a new Spinner
func NewSpinner(suffix string) *Spinner {
	spinnerSupport = !LoadBoolean(model.OktetoDisableSpinnerEnvVar) && oktetoLog.IsInteractive()
	s := sp.New(sp.CharSets[14], 100*time.Millisecond)
	s.HideCursor = true
	s.Suffix = fmt.Sprintf(" %s", suffix)
	s.FinalMSG = s.Suffix
	s.PreUpdate = func(s *sp.Spinner) {
		width, _, _ := term.GetSize(int(os.Stdout.Fd()))
		if width > 4 && len(s.FinalMSG)+2 > width {
			s.Suffix = s.FinalMSG[:width-5] + "..."
		} else {
			s.Suffix = s.FinalMSG
		}
	}

	return &Spinner{
		sp: s,
	}
}

//Start starts the spinner
func (p *Spinner) Start() {
	if spinnerSupport {
		if p.sp.FinalMSG == "" {
			p.sp.FinalMSG = p.sp.Suffix
		}
		p.sp.Start()
	} else {
		oktetoLog.Println(strings.TrimSpace(p.sp.Suffix))
	}
}

//Stop stops the spinner
func (p *Spinner) Stop() {
	if p.sp.FinalMSG != "" {
		p.sp.FinalMSG = ""
	}
	if spinnerSupport {
		p.sp.Stop()
	}
}

//Update updates the spinner message
func (p *Spinner) Update(text string) {
	p.sp.Suffix = fmt.Sprintf(" %s", ucFirst(text))
	p.sp.FinalMSG = fmt.Sprintf(" %s", ucFirst(text))
	if !spinnerSupport {
		oktetoLog.Println(strings.TrimSpace(text))
	}
}

func ucFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
