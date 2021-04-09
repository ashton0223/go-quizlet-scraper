extern crate scraper_sys;

use std::ffi::{CStr, CString};
use std::slice;
use std::str;

fn main() {
    println!("{:?}", get_study_set("aaa".to_string()));
}

fn get_study_set(url: String) -> Result<(Vec<String>, Vec<String>), String> {
    let url = CString::new(url).expect("Unable to create CString");
    let ptr = url.into_raw();
    let ret = scraper_sys::get_study_set_c(ptr);
    let mut rust_term_arr = Vec::new();
    let mut rust_def_arr = Vec::new();
    unsafe { 
        let len = ret.r2 as usize;
        let term_arr = slice::from_raw_parts(ret.r0, len);
        let def_arr = slice::from_raw_parts(ret.r1, len);
        for (i, _) in term_arr.iter().enumerate() {
            let c_str_term: &CStr = CStr::from_ptr(term_arr[i]);
            let term_slice: &str = c_str_term.to_str().unwrap();
            let c_str_def: &CStr = CStr::from_ptr(def_arr[i]);
            let def_slice: &str = c_str_def.to_str().unwrap();
            rust_term_arr.push(term_slice.to_owned().to_string());
            rust_def_arr.push(def_slice.to_owned().to_string());
        }
        CString::from_raw(ptr);

        let err = CStr::from_ptr(ret.r3)
            .to_str()
            .unwrap()
            .to_owned();
        
        if err != "" {
            return Err(err);
        }
    }
    Ok((rust_term_arr, rust_def_arr))
}