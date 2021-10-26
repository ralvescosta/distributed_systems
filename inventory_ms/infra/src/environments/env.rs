use dotenv;
use std::env::{self, VarError};

pub fn register_env() -> Result<(), VarError> {
    let rust_env = env::var("RUST_ENV")?;
    let env_file = format!(".env.{}", rust_env);

    match dotenv::from_filename(env_file).ok() {
        Some(_) => Ok(()),
        None => Err(VarError::NotPresent),
    }
}
