use async_trait::async_trait;
use std::error::Error;

use domain::{dtos::create_product_dto::CreateProductDto, entities::product_entity::ProductEntity};

#[async_trait]
pub trait IProductRepository: Send + Sync {
    async fn create(&self, dto: CreateProductDto) -> Result<ProductEntity, Box<dyn Error>>;
    async fn get_product_by_id(&self, id: String) -> Result<Option<ProductEntity>, Box<dyn Error>>;
    async fn get_products_by_type(
        &self,
        product_type: String,
    ) -> Result<Vec<ProductEntity>, Box<dyn Error>>;
    async fn get_products(
        &self,
        limit: u32,
        offset: u32,
    ) -> Result<Vec<ProductEntity>, Box<dyn Error>>;
}
