#include "ruststr.h"

// convert c string to rust string
RustStr c_str_to_rust_str(char *str) {
    return (RustStr){data : str, .len = strlen(str)};
}
