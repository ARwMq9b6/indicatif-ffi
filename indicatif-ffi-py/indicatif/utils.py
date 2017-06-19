from ._indicatif import lib


def py_str_to_rust_str(_str: str):
    return lib.c_str_to_rust_str(_str.encode())
