use async_trait::async_trait;
use core::fmt::Debug;
use std::error::Error;

use crate::entities::product_entity::ProductEntity;

#[async_trait]
pub trait IGetProductsUseCase: Send + Sync {
    async fn perform(&self, limit: u32, offset: u32) -> Result<Vec<ProductEntity>, Box<dyn Error>>;
}
impl Debug for dyn IGetProductsUseCase {
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        write!(f, "IGetProductsUseCase")
    }
}
