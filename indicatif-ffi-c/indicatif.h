#ifndef INDICATIF_FFI_INDICATIF_H
#define INDICATIF_FFI_INDICATIF_H

#include <stdint.h>

#include "ruststr.h"

#ifdef __cplusplus
extern "C" {
#endif

//##################
// ProgressStyle
//##################
// Controls the rendering style of progress bars.
typedef void *ProgressStyle;

// Returns the default progress bar style for bars.
ProgressStyle progressstyle_default_bar();

// Sets the three progress characters `(filled, current, to do)`.
void progressstyle_set_progress_chars(ProgressStyle ps, RustStr str);

// Sets the template string for the progress bar.
void progressstyle_set_template(ProgressStyle ps, RustStr str);

//##################
// ProgressBar
//##################
// A progress bar or spinner.
typedef void *ProgressBar;

// Creates a new progress bar with a given length.
//
// This progress bar by default draws directly to stdout.
ProgressBar progressbar_new(uint64_t size);

// Overrides the stored style.
void progressbar_set_style(ProgressBar pb, ProgressStyle style);

// Sets the position of the progress bar.
void progressbar_set_position(ProgressBar pb, uint64_t pos);

// Finishes the progress bar and sets a message.
void progressbar_finish_with_message(ProgressBar pb, RustStr msg);

#ifdef __cplusplus
}
#endif

#endif // INDICATIF_FFI_INDICATIF_H
