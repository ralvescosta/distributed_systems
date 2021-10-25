use std::sync::Arc;

use application::usecases::get_inventory_by_id::GetInventoryByIdUseCase;
use inventory::inventory_server::InventoryServer;
use tonic::transport::Server;

use crate::controllers::inventory_controller::InventoryController;

mod inventory;

mod controllers;
mod middlewares;
mod models;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse().unwrap();

    let get_inventory_by_id_use_case = GetInventoryByIdUseCase::new();
    let inventory_controller = InventoryController::new(Arc::new(get_inventory_by_id_use_case));

    println!("Server listening on {}", addr);
    Server::builder()
        .add_service(InventoryServer::new(inventory_controller))
        .serve(addr)
        .await?;
    Ok(())
}
