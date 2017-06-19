#ifndef INDICATIF_FFI_RUSTSTR_H
#define INDICATIF_FFI_RUSTSTR_H

#include <stddef.h>
#include <string.h>

#ifdef __cplusplus
extern "C" {
#endif

// rust string header
typedef struct {
    char *data;
    size_t len;
} RustStr;

// convert c string to rust string
RustStr c_str_to_rust_str(char *str);

#ifdef __cplusplus
}
#endif

#endif // INDICATIF_FFI_RUSTSTR_H
