package indicatif

//#include "indicatif.h"
import "C"

import "unsafe"

//##################
// ProgressStyle
//##################
// Controls the rendering style of progress bars.
type ProgressStyle uintptr

// Returns the default progress bar style for bars.
func NewProgressStyleWithDefaultBar() ProgressStyle {
	return ProgressStyle(C.progressstyle_default_bar())
}

// Sets the three progress characters `(filled, current, to do)`.
func (ps ProgressStyle) SetProgressChars(str string) {
	C.progressstyle_set_progress_chars(C.ProgressStyle(ps), goStrToC_RustStr(str))
}

// Sets the template string for the progress bar.
func (ps ProgressStyle) SetTemplate(str string) {
	C.progressstyle_set_template(C.ProgressStyle(ps), goStrToC_RustStr(str))
}

//##################
// ProgressBar
//##################
// A progress bar or spinner.
type ProgressBar uintptr

// Creates a new progress bar with a given length.
//
// This progress bar by default draws directly to stdout.
func NewProgressBar(size uint64) ProgressBar {
	return ProgressBar(C.progressbar_new(C.uint64_t(size)))
}

// Overrides the stored style.
func (pb ProgressBar) SetStyle(style ProgressStyle) {
	C.progressbar_set_style(C.ProgressBar(pb), C.ProgressStyle(style))
}

// Sets the position of the progress bar.
func (pb ProgressBar) SetPosition(pos uint64) {
	C.progressbar_set_position(C.ProgressBar(pb), C.uint64_t(pos))
}

// Finishes the progress bar and sets a message.
func (pb ProgressBar) FinishWithMessage(msg string) {
	C.progressbar_finish_with_message(C.ProgressBar(pb), goStrToC_RustStr(msg))
}

// ------Utils
// convert go string to C.RustStr
func goStrToC_RustStr(str string) C.RustStr {
	return *(*C.RustStr)(unsafe.Pointer(&str))
}
