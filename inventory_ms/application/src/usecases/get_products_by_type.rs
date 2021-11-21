use async_trait::async_trait;
use std::{error::Error, sync::Arc};

use crate::interfaces::i_product_repository::IProductRepository;
use domain::{
    entities::product_entity::ProductEntity,
    usecases::i_get_products_by_type::IGetProductsByTypeUseCase,
};

pub struct GetProductsByTypeUseCase {
    repo: Arc<dyn IProductRepository>,
}

impl GetProductsByTypeUseCase {
    pub fn new(repo: Arc<dyn IProductRepository>) -> impl IGetProductsByTypeUseCase {
        GetProductsByTypeUseCase { repo }
    }
}

#[async_trait]
impl IGetProductsByTypeUseCase for GetProductsByTypeUseCase {
    async fn perform(&self, product_type: String) -> Result<Vec<ProductEntity>, Box<dyn Error>> {
        self.repo.get_products_by_type(product_type).await
    }
}
