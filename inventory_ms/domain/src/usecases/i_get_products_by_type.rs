use async_trait::async_trait;
use core::fmt::Debug;
use std::error::Error;

use crate::entities::product_entity::ProductEntity;

#[async_trait]
pub trait IGetProductsByTypeUseCase: Send + Sync {
    async fn perform(&self, product_type: String) -> Result<Vec<ProductEntity>, Box<dyn Error>>;
}
impl Debug for dyn IGetProductsByTypeUseCase {
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        write!(f, "IGetProductsByTypeUseCase")
    }
}
