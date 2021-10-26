use std::sync::Arc;

use domain::usecases::i_get_inventory_by_id::IGetInventoryByIdUseCase;

use crate::interfaces::i_logger::ILogger;

pub struct GetInventoryByIdUseCase {
    logger: Arc<dyn ILogger>,
}

impl GetInventoryByIdUseCase {
    pub fn new(logger: Arc<dyn ILogger>) -> GetInventoryByIdUseCase {
        GetInventoryByIdUseCase { logger }
    }
}

impl IGetInventoryByIdUseCase for GetInventoryByIdUseCase {
    fn perform(&self) {
        self.logger
            .debug("[GetInventoryByIdUseCase::perform]", "Request");
    }
}
