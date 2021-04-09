extern crate bindgen;

fn main() {
    let go_path = get_gp_path();
    let path = std::fs::canonicalize("lib").expect("Could not find 'lib' directory");
    let mut cmd = std::process::Command::new(go_path)
        .arg("build")
        .arg("-buildmode=c-archive")
        .arg("-o")
        .arg("libscraper.a")
        .arg("libscraper.go")
        .env("CGO_ENABLED", "1")
        .current_dir(path)
        .spawn()
        .expect("Unable to run 'go build'");
    cmd.wait().expect("Could not wait for command");
    println!("cargo:rustc-link-search=lib/");
    println!("cargo:rustc-link-lib=static=scraper");
    println!("cargo:rerun-if-changed=libscraper.h");

    let bindings = bindgen::Builder::default()
        .header("lib/libscraper.h")
        .parse_callbacks(Box::new(bindgen::CargoCallbacks))
        .generate()
        .expect("Could not generate bindings");
    
    bindings
        .write_to_file("src/bindings.rs")
        .expect("Could not generate bindings");
}

// https://stackoverflow.com/a/35046243
fn get_gp_path() -> String {
    if let Ok(path) = std::env::var("PATH") {
        for p in path.split(":") {
            let p_str = format!("{}/{}", p, "go");
            if std::fs::metadata(&p_str).is_ok() {
                return p_str
            }
        }
    }
    "".to_string()
}