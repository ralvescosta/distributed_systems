use async_trait::async_trait;
use std::{error::Error, sync::Arc};

use crate::interfaces::i_product_repository::IProductRepository;
use domain::{
    entities::product_entity::ProductEntity, usecases::i_get_products::IGetProductsUseCase,
};

pub struct GetProductsUseCase {
    repo: Arc<dyn IProductRepository>,
}

impl GetProductsUseCase {
    pub fn new(repo: Arc<dyn IProductRepository>) -> impl IGetProductsUseCase {
        GetProductsUseCase { repo }
    }
}

#[async_trait]
impl IGetProductsUseCase for GetProductsUseCase {
    async fn perform(&self, limit: u32, offset: u32) -> Result<Vec<ProductEntity>, Box<dyn Error>> {
        self.repo.get_products(limit, offset).await
    }
}
