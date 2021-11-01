use std::{borrow::Cow, env, error::Error};
use tracing::subscriber::set_global_default;
use tracing_bunyan_formatter::{BunyanFormattingLayer, JsonStorageLayer};
use tracing_log::LogTracer;
use tracing_subscriber::{layer::SubscriberExt, EnvFilter};

pub struct Telemetry {}

impl Telemetry {
    pub fn new() -> Result<(), Box<dyn Error>> {
        let sentry_api = env::var("SENTRY_API")?;
        let app_name = env::var("APP_NAME")?;

        let _guard = sentry::init((
            sentry_api,
            sentry::ClientOptions {
                default_integrations: false,
                release: Some(Cow::from(app_name.clone())),
                server_name: Some(Cow::from(app_name)),
                ..Default::default()
            },
        ));

        let env_filter = EnvFilter::try_from_default_env()
            .unwrap_or_else(|_| EnvFilter::new(String::from("debug")));
        let formatting_layer = BunyanFormattingLayer::new(String::from("debug"), std::io::stdout);
        let subscriber = tracing_subscriber::registry()
            .with(env_filter)
            .with(JsonStorageLayer)
            .with(formatting_layer)
            .with(sentry_tracing::layer());

        LogTracer::init().expect("logger");
        set_global_default(subscriber)?;

        Ok(())
    }

    pub fn span(&self) {}
}
