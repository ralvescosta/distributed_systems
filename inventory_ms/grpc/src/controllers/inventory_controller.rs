use tonic::{Request, Response, Status};

use crate::inventory::{
    inventory_server::Inventory, CreateProductRequest, Empty, GetByIdRequest, ProductResponse,
    ProductsResponse, UpdateProductRequest,
};

#[derive(Default)]
pub struct InventoryController;

#[tonic::async_trait]
impl Inventory for InventoryController {
    async fn get_product_by_id(
        &self,
        request: Request<GetByIdRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        todo!()
    }

    async fn get_products(
        &self,
        request: Request<Empty>,
    ) -> Result<Response<ProductsResponse>, Status> {
        todo!()
    }

    async fn create_product(
        &self,
        request: Request<CreateProductRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        todo!()
    }

    async fn update_product(
        &self,
        request: Request<UpdateProductRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        todo!()
    }
}
