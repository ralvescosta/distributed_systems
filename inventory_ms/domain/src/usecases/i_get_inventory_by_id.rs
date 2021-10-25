pub trait IGetInventoryByIdUseCase {
    fn perform(&self);
}

impl IGetInventoryByIdUseCase for dyn Send {
    fn perform(&self) {
        todo!()
    }
}

impl IGetInventoryByIdUseCase for dyn Sync {
    fn perform(&self) {
        todo!()
    }
}
