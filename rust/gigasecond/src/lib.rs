extern crate chrono;
use chrono::*;

// Returns a Utc DateTime one billion seconds after start.
pub fn after(start: DateTime<Utc>) -> DateTime<Utc> {
    //unimplemented!()
    return start + Duration::seconds(1_000_000_000);
}
