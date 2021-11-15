use async_trait::async_trait;
use core::fmt::Debug;
use std::error::Error;

use crate::entities::product_entity::ProductEntity;

#[async_trait]
pub trait IGetProductByIdUseCase: Send + Sync {
    async fn perform(&self, id: String) -> Result<Option<ProductEntity>, Box<dyn Error>>;
}
impl Debug for dyn IGetProductByIdUseCase {
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        write!(f, "IGetProductByIdUseCase")
    }
}
