use log::info;
use std::sync::Arc;
use tonic::transport::Server;

use application::usecases::{
    create_product::CreateProductUseCase, get_product_by_id::GetProductByIdUseCase,
};
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

    let telemetry_app = Telemetry::new()?;

    let db_connection = database::connection::connect_to_database().await?;

    let addr = "127.0.0.1:50051".parse().unwrap();

    let product_repository = Arc::new(ProductRepository::new(db_connection));
    let get_product_by_id_use_case =
        Arc::new(GetProductByIdUseCase::new(product_repository.clone()));
    let create_product_use_case = Arc::new(CreateProductUseCase::new(product_repository));
    let product_controller = ProductController::new(
        get_product_by_id_use_case,
        create_product_use_case,
        Arc::new(telemetry_app),
    );

    info!("Server listening on {}", addr);

    Server::builder()
        .add_service(InventoryServer::new(product_controller))
        .serve(addr)
        .await?;
    Ok(())
}
