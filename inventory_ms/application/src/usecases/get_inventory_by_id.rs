use async_trait::async_trait;
use std::{error::Error, sync::Arc};

use crate::interfaces::{i_logger::ILogger, i_product_repository::IProductRepository};
use domain::{
    entities::product_entity::ProductEntity,
    usecases::i_get_inventory_by_id::IGetInventoryByIdUseCase,
};

pub struct GetInventoryByIdUseCase {
    logger: Arc<dyn ILogger>,
    repo: Arc<dyn IProductRepository>,
}

impl GetInventoryByIdUseCase {
    pub fn new(
        logger: Arc<dyn ILogger>,
        repo: Arc<dyn IProductRepository>,
    ) -> GetInventoryByIdUseCase {
        GetInventoryByIdUseCase { logger, repo }
    }
}

#[async_trait]
impl IGetInventoryByIdUseCase for GetInventoryByIdUseCase {
    async fn perform(&self, id: String) -> Result<Option<ProductEntity>, Box<dyn Error>> {
        self.logger
            .debug("GetInventoryByIdUseCase::perform", "Request");
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
        let mut mocked_logger = MockLogger::new();
        let mut mocked_repository = MockProductRepository::new();

        mocked_logger.expect_debug().returning(|_t, _a| {}).times(1);
        mocked_repository
            .expect_get_product_by_id()
            .returning(|_id| Ok(Some(ProductEntity::default())))
            .times(1);

        let sut =
            GetInventoryByIdUseCase::new(Arc::new(mocked_logger), Arc::new(mocked_repository));

        match sut.perform(String::from("some")).await {
            Ok(Some(result)) => assert_eq!(result.title, ProductEntity::default().title),
            _ => assert!(false),
        }
    }

    #[tokio::test]
    async fn should_return_none_with_books_no_exist() {
        let mut mocked_logger = MockLogger::new();
        let mut mocked_repository = MockProductRepository::new();

        mocked_logger.expect_debug().returning(|_t, _a| {}).times(1);
        mocked_repository
            .expect_get_product_by_id()
            .returning(|_id| Ok(None))
            .times(1);

        let sut =
            GetInventoryByIdUseCase::new(Arc::new(mocked_logger), Arc::new(mocked_repository));
        match sut.perform(String::from("some")).await {
            Ok(None) => assert!(true),
            _ => assert!(false),
        }
    }

    #[tokio::test]
    async fn should_return_error_if_repository_returns_error() {
        let mut mocked_logger = MockLogger::new();
        let mut mocked_repository = MockProductRepository::new();

        mocked_logger.expect_debug().returning(|_t, _a| {}).times(1);
        mocked_repository
            .expect_get_product_by_id()
            .returning(|_id| Err(Box::new(Error::from(ErrorKind::Unsupported))))
            .times(1);

        let sut =
            GetInventoryByIdUseCase::new(Arc::new(mocked_logger), Arc::new(mocked_repository));
        match sut.perform(String::from("some")).await {
            Err(_err) => assert!(true),
            _ => assert!(false),
        }
    }

    mock! {
        pub Logger {}
        impl ILogger for Logger {
            fn trace(&self, _target: &str, _msg: &str) {}
            fn debug(&self, _target: &str, _msg: &str) {}
            fn info(&self, _target: &str, _msg: &str) {}
            fn warn(&self, _target: &str, _msg: &str) {}
            fn error(&self, _target: &str, _msg: &str) {}
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
        }
    }
}
