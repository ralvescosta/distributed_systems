use std::sync::Arc;
use tonic::{Request, Response, Status};

use crate::inventory::{
    inventory_server::Inventory, CreateProductRequest, Empty, GetByIdRequest, ProductResponse,
    ProductsResponse, UpdateProductRequest,
};
use domain::usecases::i_get_inventory_by_id::IGetInventoryByIdUseCase;

pub struct InventoryController {
    get_product_by_id_use_case: Arc<dyn IGetInventoryByIdUseCase>,
}

unsafe impl Sync for InventoryController {}
unsafe impl Send for InventoryController {}

impl InventoryController {
    pub fn new(
        get_product_by_id_use_case: Arc<dyn IGetInventoryByIdUseCase>,
    ) -> InventoryController {
        InventoryController {
            get_product_by_id_use_case,
        }
    }
}

#[tonic::async_trait]
impl Inventory for InventoryController {
    async fn get_product_by_id(
        &self,
        _request: Request<GetByIdRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        self.get_product_by_id_use_case.perform();

        Ok(Response::new(ProductResponse {
            id: 1,
            tag: String::new(),
            title: String::new(),
            subtitle: String::new(),
            authors: vec![],
            amount_in_stock: 10,
            created_at: String::new(),
            updated_at: String::new(),
        }))
    }

    async fn get_products(
        &self,
        _request: Request<Empty>,
    ) -> Result<Response<ProductsResponse>, Status> {
        Ok(Response::new(ProductsResponse { value: vec![] }))
    }

    async fn create_product(
        &self,
        _request: Request<CreateProductRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        Ok(Response::new(ProductResponse {
            id: 1,
            tag: String::new(),
            title: String::new(),
            subtitle: String::new(),
            authors: vec![],
            amount_in_stock: 10,
            created_at: String::new(),
            updated_at: String::new(),
        }))
    }

    async fn update_product(
        &self,
        _request: Request<UpdateProductRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        Ok(Response::new(ProductResponse {
            id: 1,
            tag: String::new(),
            title: String::new(),
            subtitle: String::new(),
            authors: vec![],
            amount_in_stock: 10,
            created_at: String::new(),
            updated_at: String::new(),
        }))
    }
}
