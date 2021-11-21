use async_trait::async_trait;
use log::debug;
use std::{error::Error, sync::Arc};

use crate::interfaces::i_product_repository::IProductRepository;
use domain::{
    entities::product_entity::ProductEntity, usecases::i_get_product_by_id::IGetProductByIdUseCase,
};

pub struct GetProductByIdUseCase {
    repo: Arc<dyn IProductRepository>,
}

impl GetProductByIdUseCase {
    pub fn new(repo: Arc<dyn IProductRepository>) -> GetProductByIdUseCase {
        GetProductByIdUseCase { repo }
    }
}

#[async_trait]
impl IGetProductByIdUseCase for GetProductByIdUseCase {
    async fn perform(&self, id: String) -> Result<Option<ProductEntity>, Box<dyn Error>> {
        debug!("Request");
        self.repo.get_product_by_id(id).await
    }
}

#[cfg(test)]
mod test {
    use super::*;
    use domain::dtos::create_product_dto::CreateProductDto;
    use mockall::*;
    use std::io::{Error, ErrorKind};

    #[tokio::test]
    async fn perform() {
        let mut mocked_repository = MockProductRepository::new();

        mocked_repository
            .expect_get_product_by_id()
            .returning(|_id| Ok(Some(ProductEntity::default())))
            .times(1);

        let sut = GetProductByIdUseCase::new(Arc::new(mocked_repository));

        match sut.perform(String::from("some")).await {
            Ok(Some(result)) => assert_eq!(result.title, ProductEntity::default().title),
            _ => assert!(false),
        }
    }

    #[tokio::test]
    async fn should_return_none_with_books_no_exist() {
        let mut mocked_repository = MockProductRepository::new();

        mocked_repository
            .expect_get_product_by_id()
            .returning(|_id| Ok(None))
            .times(1);

        let sut = GetProductByIdUseCase::new(Arc::new(mocked_repository));
        match sut.perform(String::from("some")).await {
            Ok(None) => assert!(true),
            _ => assert!(false),
        }
    }

    #[tokio::test]
    async fn should_return_error_if_repository_returns_error() {
        let mut mocked_repository = MockProductRepository::new();

        mocked_repository
            .expect_get_product_by_id()
            .returning(|_id| Err(Box::new(Error::from(ErrorKind::Unsupported))))
            .times(1);

        let sut = GetProductByIdUseCase::new(Arc::new(mocked_repository));
        match sut.perform(String::from("some")).await {
            Err(_err) => assert!(true),
            _ => assert!(false),
        }
    }

    mock! {
        pub ProductRepository{}
        #[async_trait]
        impl IProductRepository for ProductRepository {
            async fn get_product_by_id(
                &self,
                _id: String,
            ) -> Result<Option<ProductEntity>, Box<dyn std::error::Error>> {
                todo!()
            }
            async fn create(&self, dto: CreateProductDto) -> Result<ProductEntity, Box<dyn std::error::Error>> {
                todo!()
            }
            async fn get_products_by_type(&self, product_type: String) -> Result<Vec<ProductEntity>, Box<dyn std::error::Error>> {
                todo!()
            }
            async fn get_products(
                &self,
                limit: u32,
                offset: u32,
            ) -> Result<Vec<ProductEntity>, Box<dyn std::error::Error>> {
                todo!()
            }
        }
    }
}
