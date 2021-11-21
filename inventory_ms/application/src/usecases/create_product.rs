use async_trait::async_trait;
use std::{error::Error, sync::Arc};

use crate::interfaces::i_product_repository::IProductRepository;
use domain::{
    dtos::create_product_dto::CreateProductDto, entities::product_entity::ProductEntity,
    usecases::i_create_product::ICreateProductUseCase,
};
pub struct CreateProductUseCase {
    repo: Arc<dyn IProductRepository>,
}

impl CreateProductUseCase {
    pub fn new(repo: Arc<dyn IProductRepository>) -> impl ICreateProductUseCase {
        CreateProductUseCase { repo }
    }
}

#[async_trait]
impl ICreateProductUseCase for CreateProductUseCase {
    async fn perform(&self, dto: CreateProductDto) -> Result<ProductEntity, Box<dyn Error>> {
        self.repo.create(dto).await
    }
}
