use std::ffi::CString;
use std::os::raw::c_char;
use std::mem::{swap, forget, transmute_copy};

use indicatif::{ProgressBar, ProgressStyle};

//##################
// ProgressBar
//##################
#[no_mangle]
pub extern "C" fn progressbar_new(size: u64) -> *const ProgressBar {
    let pb = Box::new(ProgressBar::new(size));
    let ptr = Box::into_raw(pb);

    ptr
}

#[no_mangle]
pub extern "C" fn progressbar_set_style(pb: *const ProgressBar, style: *const ProgressStyle) {
    let style: ProgressStyle = unsafe { transmute_copy(&*style) };
    let pb = unsafe { &*pb };
    pb.set_style(style)
}

#[no_mangle]
pub extern "C" fn progressbar_set_position(pb: *const ProgressBar, pos: u64) {
    let pb = unsafe { &*pb };
    pb.set_position(pos)
}

#[no_mangle]
pub extern "C" fn progressbar_finish_with_message(pb: *const ProgressBar, msg: &str) {
    unsafe {
        let pb = &*pb;
        pb.finish_with_message(msg)
    }
}

//##################
// ProgressStyle
//##################
#[no_mangle]
pub extern "C" fn progressstyle_default_bar() -> *const ProgressStyle {
    let ps = Box::new(ProgressStyle::default_bar());
    let ptr = Box::into_raw(ps);

    ptr
}

#[no_mangle]
pub extern "C" fn progressstyle_set_progress_chars(ps: *mut ProgressStyle, str: &str) {
    unsafe {
        let mut style: ProgressStyle = transmute_copy(&*ps);
        let mut style = style.progress_chars(str);

        swap(&mut *ps, &mut style);
        forget(style);
    }
}

#[no_mangle]
pub extern "C" fn progressstyle_set_template(ps: *mut ProgressStyle, str: &str) {
    unsafe {
        let mut style: ProgressStyle = transmute_copy(&*ps);
        let mut style = style.template(str);
        swap(&mut *ps, &mut style);

        forget(style);
    }
}
