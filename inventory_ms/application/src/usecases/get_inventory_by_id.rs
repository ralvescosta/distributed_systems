use domain::usecases::i_get_inventory_by_id::IGetInventoryByIdUseCase;

pub struct GetInventoryByIdUseCase {}

impl GetInventoryByIdUseCase {
    pub fn new() -> GetInventoryByIdUseCase {
        GetInventoryByIdUseCase {}
    }
}

impl IGetInventoryByIdUseCase for GetInventoryByIdUseCase {
    fn perform(&self) {
        println!("[GetInventoryByIdUseCase::perform]")
    }
}
