use std::{env, error::Error};

use opentelemetry::{
    global,
    propagation::{Extractor, Injector},
    sdk::propagation::TraceContextPropagator,
    Context,
};
use tracing_bunyan_formatter::{BunyanFormattingLayer, JsonStorageLayer};
use tracing_opentelemetry::OpenTelemetrySpanExt;
use tracing_subscriber::prelude::*;

#[derive(Debug)]
pub struct Telemetry {}

impl Telemetry {
    pub fn new() -> Result<Telemetry, Box<dyn Error>> {
        let app_name = env::var("APP_NAME")?;
        let log_level = env::var("RUST_LOG")?;

        global::set_text_map_propagator(TraceContextPropagator::new());
        let tracer = opentelemetry_jaeger::new_pipeline()
            .with_service_name(app_name)
            .install_batch(opentelemetry::runtime::Tokio)?;

        let formatting_layer =
            BunyanFormattingLayer::new(String::from(log_level.as_str()), std::io::stdout);
        tracing_subscriber::registry()
            .with(tracing_subscriber::EnvFilter::new(log_level.as_str()))
            .with(tracing_opentelemetry::layer().with_tracer(tracer))
            .with(JsonStorageLayer)
            .with(formatting_layer)
            .try_init()?;

        Ok(Telemetry {})
    }

    pub fn grpc_set_span_parent<T>(&self, request: &tonic::Request<T>) -> Context {
        let ctx = global::get_text_map_propagator(|prop| {
            prop.extract(&GrpcExtractor(request.metadata()))
        });
        tracing::Span::current().set_parent(ctx.clone());
        ctx
    }

    pub fn inject_grpc_ctx<T>(mut request: tonic::Request<T>) {
        global::get_text_map_propagator(|propagator| {
            propagator.inject_context(
                &tracing::Span::current().context(),
                &mut GrpcInjector(request.metadata_mut()),
            )
        });
    }

    pub fn span(&self) {}
}

struct GrpcExtractor<'a>(&'a tonic::metadata::MetadataMap);
impl<'a> Extractor for GrpcExtractor<'a> {
    /// Get a value for a key from the MetadataMap.  If the value can't be converted to &str, returns None
    fn get(&self, key: &str) -> Option<&str> {
        self.0.get(key).and_then(|metadata| metadata.to_str().ok())
    }

    /// Collect all the keys from the MetadataMap.
    fn keys(&self) -> Vec<&str> {
        self.0
            .keys()
            .map(|key| match key {
                tonic::metadata::KeyRef::Ascii(v) => v.as_str(),
                tonic::metadata::KeyRef::Binary(v) => v.as_str(),
            })
            .collect::<Vec<_>>()
    }
}

struct GrpcInjector<'a>(&'a mut tonic::metadata::MetadataMap);
impl<'a> Injector for GrpcInjector<'a> {
    /// Set a key and value in the MetadataMap.  Does nothing if the key or value are not valid inputs
    fn set(&mut self, key: &str, value: String) {
        if let Ok(key) = tonic::metadata::MetadataKey::from_bytes(key.as_bytes()) {
            if let Ok(val) = tonic::metadata::MetadataValue::from_str(&value) {
                self.0.insert(key, val);
            }
        }
    }
}
