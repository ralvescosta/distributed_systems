use newrelic::{App, AppConfig};
use std::{env, error::Error, thread, time::Duration};

pub struct Telemetry {
    application: App,
}

impl Telemetry {
    pub fn new() -> Result<Telemetry, Box<dyn Error>> {
        let key = env::var("NEW_RELIC_LICENSE_KEY")?;
        let app_name = env::var("APP_NAME")?;

        let app_config = AppConfig::new(&app_name, &key)?;
        let app = App::with_timeout(app_config, 9999)?;

        thread::sleep(Duration::from_secs(4));

        Ok(Telemetry { application: app })
    }

    pub fn span(&self) {}
}
