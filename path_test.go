package ctxgrp

import (
	"testing"
)

func TestPanicWhenBothEmpty(t *testing.T) {
	shouldPanic(t, "", "")
}
func TestPanicWhenNoLeading(t *testing.T) {
	shouldPanic(t, "foo", "bar")
}
func TestEmptyPost(t *testing.T) {
	v(t, "/pre", "/pre", "")
}
func TestEmptyPre(t *testing.T) {
	v(t, "/post", "", "/post")
}
func TestEmptyPreWithLeadingPost(t *testing.T) {
	v(t, "/pre/post", "/pre", "/post")
}
func TestTrailingPreWithLeadingPost(t *testing.T) {
	v(t, "/pre/post", "/pre/", "/post")
}

func v(t *testing.T, expected, pre, post string) {
	if got := mkpath(pre, post); got != expected {
		t.Errorf(`Expected "%s", but mkpath("%s", "%s") returned: "%s"`, expected, pre, post, got)
	}
}
func shouldPanic(t *testing.T, pre string, post string) {
	defer func() {
		if err := recover(); err != nil {
			// we got an error, as expected
		}
	}()
	got := mkpath(pre, post)
	t.Errorf(`Expected panic, but mkpath("%s", "%s") returned: "%s"`, pre, post, got)
}
