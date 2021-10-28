use async_trait::async_trait;
use std::sync::Arc;

use domain::usecases::i_get_inventory_by_id::IGetInventoryByIdUseCase;

use crate::interfaces::{i_logger::ILogger, i_product_repository::IProductRepository};

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
    async fn perform(&self) {
        let _product = self.repo.get_product_by_id("1234".to_string()).await;
        self.logger
            .debug("GetInventoryByIdUseCase::perform", "Request");
    }
}
