use infra::logger::logger::Logger;
use std::sync::Arc;
use tonic::transport::Server;

use application::usecases::get_inventory_by_id::GetInventoryByIdUseCase;
use infra::environments;

use crate::controllers::inventory_controller::InventoryController;
use crate::inventory::inventory_server::InventoryServer;

mod controllers;
mod inventory;
mod middlewares;
mod models;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    environments::env::register_env()?;
    Logger::init();

    let addr = "127.0.0.1:50051".parse().unwrap();

    let logger = Arc::new(Logger::new());
    let get_inventory_by_id_use_case = Arc::new(GetInventoryByIdUseCase::new(logger));
    let inventory_controller = InventoryController::new(get_inventory_by_id_use_case);

    println!("Server listening on {}", addr);
    Server::builder()
        .add_service(InventoryServer::new(inventory_controller))
        .serve(addr)
        .await?;
    Ok(())
}