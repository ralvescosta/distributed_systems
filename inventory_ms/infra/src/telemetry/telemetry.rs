use std::{env, error::Error};

use opentelemetry::{global, sdk::propagation::TraceContextPropagator};
use tracing_bunyan_formatter::{BunyanFormattingLayer, JsonStorageLayer};
use tracing_log::LogTracer;
use tracing_subscriber::prelude::*;

pub struct Telemetry {}

impl Telemetry {
    pub fn new() -> Result<(), Box<dyn Error>> {
        let app_name = env::var("APP_NAME")?;

        global::set_text_map_propagator(TraceContextPropagator::new());
        let tracer = opentelemetry_jaeger::new_pipeline()
            .with_service_name(app_name)
            .install_batch(opentelemetry::runtime::Tokio)?;

        let formatting_layer = BunyanFormattingLayer::new(String::from("debug"), std::io::stdout);
        tracing_subscriber::registry()
            .with(tracing_subscriber::EnvFilter::new("debug"))
            .with(tracing_opentelemetry::layer().with_tracer(tracer))
            .with(JsonStorageLayer)
            .with(formatting_layer)
            .try_init()?;

        // match LogTracer::init() {
        //     Err(err) => {
        //         println!("ERROR AQUI");
        //         println!("{}", err.to_string())
        //     }
        //     _ => (),
        // }

        Ok(())
    }

    pub fn span(&self) {}
}
