// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestBasicJobs(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	out, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("Failed to read stdout: %v", err)
	}
	got := string(out)

	want := "(?s).*Company created:.*Job created:.*Job retrieved:.*Job updated:.*Job updated:.*Job deleted.*Company deleted"
	if ok, err := regexp.MatchString(want, got); !ok {
		t.Errorf("stdout returned %s, wanted to contain %s, err: %v", got, want, err)
	}
}
