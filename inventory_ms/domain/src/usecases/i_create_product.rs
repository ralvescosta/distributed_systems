use async_trait::async_trait;
use std::error::Error;

use crate::{dtos::create_product_dto::CreateProductDto, entities::product_entity::ProductEntity};

#[async_trait]
pub trait ICreateProductUseCase: Send + Sync {
    async fn perform(&self, dto: CreateProductDto) -> Result<ProductEntity, Box<dyn Error>>;
}
