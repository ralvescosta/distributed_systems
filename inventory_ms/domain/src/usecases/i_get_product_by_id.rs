use async_trait::async_trait;
use std::error::Error;

use crate::entities::product_entity::ProductEntity;

#[async_trait]
pub trait IGetProductByIdUseCase: Send + Sync {
    async fn perform(&self, id: String) -> Result<Option<ProductEntity>, Box<dyn Error>>;
}
