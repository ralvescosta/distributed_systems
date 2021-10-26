pub trait ILogger {
    fn trace(&self, target: &str, msg: &str);
    fn debug(&self, target: &str, msg: &str);
    fn info(&self, target: &str, msg: &str);
    fn warn(&self, target: &str, msg: &str);
    fn error(&self, target: &str, msg: &str);
}
