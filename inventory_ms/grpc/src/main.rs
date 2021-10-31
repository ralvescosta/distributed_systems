use application::usecases::create_product::CreateProductUseCase;
use infra::logger::logger::Logger;
use std::sync::Arc;
use tonic::transport::Server;

use application::usecases::get_product_by_id::GetProductByIdUseCase;
use infra::{
    database, environments, repositories::product_repository::ProductRepository,
    telemetry::telemetry::Telemetry,
};

use crate::controllers::product_controller::ProductController;
use crate::inventory::inventory_server::InventoryServer;

mod controllers;
mod inventory;
mod middlewares;
mod models;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    environments::env::register_env()?;
    Logger::init();

    Telemetry::new();

    let db_connection = database::connection::connect_to_database().await?;

    let addr = "127.0.0.1:50051".parse().unwrap();

    let logger = Arc::new(Logger::new());
    let product_repository = Arc::new(ProductRepository::new(db_connection));
    let get_product_by_id_use_case = Arc::new(GetProductByIdUseCase::new(
        logger,
        product_repository.clone(),
    ));
    let create_product_use_case = Arc::new(CreateProductUseCase::new(product_repository));
    let product_controller =
        ProductController::new(get_product_by_id_use_case, create_product_use_case);

    println!("Server listening on {}", addr);
    Server::builder()
        .add_service(InventoryServer::new(product_controller))
        .serve(addr)
        .await?;
    Ok(())
}
