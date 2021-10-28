use async_trait::async_trait;
use domain::entities::product_entity::ProductEntity;

#[async_trait]
pub trait IProductRepository: Send + Sync {
    async fn get_product_by_id(
        &self,
        id: String,
    ) -> Result<Option<ProductEntity>, Box<dyn std::error::Error>>;
}
