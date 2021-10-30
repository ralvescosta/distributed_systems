use std::sync::Arc;
use tonic::{Request, Response, Status};

use crate::inventory::{
    inventory_server::Inventory, CreateProductRequest, Empty, GetByIdRequest, ProductResponse,
    ProductsResponse, UpdateProductRequest,
};
use domain::usecases::i_get_inventory_by_id::IGetInventoryByIdUseCase;

pub struct InventoryController {
    get_product_by_id_use_case: Arc<dyn IGetInventoryByIdUseCase>,
}

impl InventoryController {
    pub fn new(
        get_product_by_id_use_case: Arc<dyn IGetInventoryByIdUseCase>,
    ) -> InventoryController {
        InventoryController {
            get_product_by_id_use_case,
        }
    }
}

#[tonic::async_trait]
impl Inventory for InventoryController {
    async fn get_product_by_id(
        &self,
        request: Request<GetByIdRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        let result = self
            .get_product_by_id_use_case
            .perform(request.into_inner().id)
            .await;
        match result {
            Ok(Some(product)) => Ok(Response::new(ProductResponse {
                id: product.id,
                product_category: product.product_category,
                tag: product.tag,
                title: product.title,
                subtitle: product.subtitle,
                authors: product.authors,
                amount_in_stock: product.amount_in_stock,
                created_at: product.created_at,
                updated_at: product.updated_at,
                num_pages: product.num_pages,
                tags: product.tags,
            })),
            Ok(None) => Err(Status::not_found("Not Found")),
            Err(err) => Err(Status::internal(format!("{:?}", err))),
        }
    }

    async fn get_products(
        &self,
        _request: Request<Empty>,
    ) -> Result<Response<ProductsResponse>, Status> {
        Ok(Response::new(ProductsResponse { value: vec![] }))
    }

    async fn create_product(
        &self,
        _request: Request<CreateProductRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        Ok(Response::new(ProductResponse {
            id: String::new(),
            product_category: String::new(),
            tag: String::new(),
            title: String::new(),
            subtitle: String::new(),
            authors: vec![],
            amount_in_stock: 10,
            created_at: String::new(),
            updated_at: String::new(),
            num_pages: 10,
            tags: vec![],
        }))
    }

    async fn update_product(
        &self,
        _request: Request<UpdateProductRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        Ok(Response::new(ProductResponse {
            id: String::new(),
            product_category: String::new(),
            tag: String::new(),
            title: String::new(),
            subtitle: String::new(),
            authors: vec![],
            amount_in_stock: 10,
            created_at: String::new(),
            updated_at: String::new(),
            num_pages: 10,
            tags: vec![],
        }))
    }
}

#[cfg(test)]
mod test {

    use super::*;
    use domain::entities::product_entity::ProductEntity;
    use mockall::*;

    #[tokio::test]
    async fn perform() {
        let mut use_case_mocked = MockGetInventoryByIdUseCase::new();
        let request = Request::<GetByIdRequest>::new(GetByIdRequest {
            ..Default::default()
        });
        use_case_mocked
            .expect_perform()
            .returning(|_id| Ok(Some(ProductEntity::default())))
            .times(1);

        let sut = InventoryController::new(Arc::new(use_case_mocked));

        match sut.get_product_by_id(request).await {
            Ok(result) => assert_eq!(result.get_ref().title, String::from("")),
            _ => assert!(false),
        }
    }

    mock! {
        pub GetInventoryByIdUseCase {}
        #[tonic::async_trait]
        impl IGetInventoryByIdUseCase for GetInventoryByIdUseCase {
            async fn perform(
                &self,
                id: String,
            ) -> Result<
                Option<domain::entities::product_entity::ProductEntity>,
                Box<dyn std::error::Error>,
            > {
                todo!()
            }
        }
    }
}
