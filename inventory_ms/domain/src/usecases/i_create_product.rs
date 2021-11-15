use async_trait::async_trait;
use core::fmt::Debug;
use std::error::Error;

use crate::{dtos::create_product_dto::CreateProductDto, entities::product_entity::ProductEntity};

#[async_trait]
pub trait ICreateProductUseCase: Send + Sync {
    async fn perform(&self, dto: CreateProductDto) -> Result<ProductEntity, Box<dyn Error>>;
}

impl Debug for dyn ICreateProductUseCase {
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        write!(f, "ICreateProductUseCase")
    }
}
