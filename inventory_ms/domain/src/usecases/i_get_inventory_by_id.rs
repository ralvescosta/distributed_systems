use async_trait::async_trait;

#[async_trait]
pub trait IGetInventoryByIdUseCase: Send + Sync {
    async fn perform(&self);
}
