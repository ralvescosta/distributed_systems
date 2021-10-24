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
    let inventory_controller = InventoryController::default();
    println!("Server listening on {}", addr);
    Server::builder()
        .add_service(InventoryServer::new(inventory_controller))
        .serve(addr)
        .await?;
    Ok(())
}
