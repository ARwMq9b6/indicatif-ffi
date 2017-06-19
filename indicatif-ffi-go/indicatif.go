package indicatif

//##################
// ProgressStyle
//##################
// Controls the rendering style of progress bars.
type ProgressStyle uintptr

// Returns the default progress bar style for bars.
func NewProgressStyleWithDefaultBar() ProgressStyle {
	return progressstyle_default_bar()
}

//extern progressstyle_default_bar
func progressstyle_default_bar() ProgressStyle

// Sets the three progress characters `(filled, current, to do)`.
func (ps ProgressStyle) SetProgressChars(str string) {
	progressstyle_set_progress_chars(ps, str)
}

//extern progressstyle_set_progress_chars
func progressstyle_set_progress_chars(ps ProgressStyle, str string)

// Sets the template string for the progress bar.
func (ps ProgressStyle) SetTemplate(str string) {
	progressstyle_set_template(ps, str)
}

//extern progressstyle_set_template
func progressstyle_set_template(ps ProgressStyle, str string)

//##################
// ProgressBar
//##################
// A progress bar or spinner.
type ProgressBar uintptr

// Creates a new progress bar with a given length.
//
// This progress bar by default draws directly to stdout.
func NewProgressBar(size uint64) ProgressBar {
	return progressbar_new(size)
}

//extern progressbar_new
func progressbar_new(size uint64) ProgressBar

// Overrides the stored style.
func (pb ProgressBar) SetStyle(style ProgressStyle) {
	progressbar_set_style(pb, style)
}

//extern progressbar_set_style
func progressbar_set_style(pb ProgressBar, style ProgressStyle)

// Sets the position of the progress bar.
func (pb ProgressBar) SetPosition(pos uint64) {
	progressbar_set_position(pb, pos)
}

//extern progressbar_set_position
func progressbar_set_position(pb ProgressBar, pos uint64)

// Finishes the progress bar and sets a message.
func (pb ProgressBar) FinishWithMessage(msg string) {
	progressbar_finish_with_message(pb, msg)
}

//extern progressbar_finish_with_message
func progressbar_finish_with_message(pb ProgressBar, msg string)
