#[link(wasm_import_module = "env")]
extern "C" {
  fn ws_log(log_level: i32, ptr: *const u8, size: i32) -> i32;
  fn ws_get_data(event_id: i32, payload_ptr: *const *mut u8, payload_size: &i32) -> i32;
}

#[no_mangle]
// This handler will be matched by the default Project event strategy in W3bstream
pub extern "C" fn start(event_id: i32) -> i32 {
    log_info(
        &format!("Start handler called with event_id: {}", event_id));

    let payload = get_data_as_str(event_id).unwrap();
    log_info(&format!("event data as string: {}", payload));
    return 0;
}

// Returns the event payload as a string
fn get_data_as_str(event_id: i32) -> Option<String> {
    let data_ptr = &mut (0 as i32) as *const _ as *const *mut u8;
    let data_size = &(0 as i32);
    match unsafe { ws_get_data(event_id, data_ptr, data_size) } {
        0 => Some(unsafe { String::from_raw_parts(*data_ptr, *data_size as _, *data_size as _) }),
        _ => None,
    }
}

// Logs an info message string to the W3bstream console
fn log_info(str: &str) {
    unsafe { ws_log(3, str.as_bytes().as_ptr(), str.len() as _) };
}