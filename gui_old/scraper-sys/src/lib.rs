#![allow(improper_ctypes)]
#![allow(non_upper_case_globals)]
#![allow(non_camel_case_types)]
#![allow(non_snake_case)]


include!("bindings.rs");

pub fn get_study_set_c(ptr: *mut i8) -> GetStudySetC_return {
    unsafe { return GetStudySetC(ptr) }
}