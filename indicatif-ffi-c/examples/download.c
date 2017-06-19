#include <stdint.h>
#include <stdio.h>

#include "indicatif.h"

#ifndef min
#define min(a, b) (((a) < (b)) ? (a) : (b))
#endif

int main() {
    uint64_t downloaded = 0;
    uint64_t total_size = 231231231;

    ProgressStyle ps = progressstyle_default_bar();
    progressstyle_set_progress_chars(ps, c_str_to_rust_str("#>-"));
    progressstyle_set_template(
        ps, c_str_to_rust_str("{spinner:.green} [{elapsed_precise}] "
                              "[{bar:40.cyan/blue}] {bytes}/{total_bytes} "
                              "({eta})"));

    ProgressBar pb = progressbar_new(total_size);
    progressbar_set_style(pb, ps);

    while (downloaded < total_size) {
        uint64_t new = min(downloaded + 223211, total_size);
        downloaded = new;

        progressbar_set_position(pb, new);
        usleep(12 * 1000);
    }

    progressbar_finish_with_message(pb, c_str_to_rust_str("downloaded"));
}
