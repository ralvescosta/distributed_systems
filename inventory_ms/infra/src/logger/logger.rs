use std::fmt::Debug;

use application::interfaces::i_logger::ILogger;
use env_logger;
use log::{debug, error, info, trace, warn};

#[derive(Debug, Clone, Copy)]
pub struct Logger;

impl Logger {
    pub fn init() {
        env_logger::init();
    }
    pub fn new() -> impl ILogger {
        Logger {}
    }

    pub fn test<T: Debug>() {
        println!("{:?}", std::any::type_name::<T>());
    }
}

impl ILogger for Logger {
    fn trace(&self, target: &str, msg: &str) {
        trace!(target: target, "{}", msg)
    }

    fn debug(&self, target: &str, msg: &str) {
        debug!(target: target, "{}", msg)
    }

    fn info(&self, target: &str, msg: &str) {
        info!(target: target, "{}", msg)
    }

    fn warn(&self, target: &str, msg: &str) {
        warn!(target: target, "{}", msg)
    }

    fn error(&self, target: &str, msg: &str) {
        error!(target: target, "{}", msg)
    }
}
