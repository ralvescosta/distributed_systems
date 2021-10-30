use std::sync::Arc;
use tonic::{Request, Response, Status};

use crate::{
    inventory::{
        inventory_server::Inventory, CreateProductRequest, Empty, GetByIdRequest, ProductResponse,
        ProductsResponse, UpdateProductRequest,
    },
    models::product_model::ProductModel,
};
use domain::usecases::{
    i_create_product::ICreateProductUseCase, i_get_inventory_by_id::IGetInventoryByIdUseCase,
};

pub struct InventoryController {
    get_product_by_id_use_case: Arc<dyn IGetInventoryByIdUseCase>,
    create_product_use_case: Arc<dyn ICreateProductUseCase>,
}

impl InventoryController {
    pub fn new(
        get_product_by_id_use_case: Arc<dyn IGetInventoryByIdUseCase>,
        create_product_use_case: Arc<dyn ICreateProductUseCase>,
    ) -> InventoryController {
        InventoryController {
            get_product_by_id_use_case,
            create_product_use_case,
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
            Ok(Some(product)) => Ok(Response::new(ProductModel::entity_to_response(product))),
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
        request: Request<CreateProductRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        match self
            .create_product_use_case
            .perform(ProductModel::create_request_to_dto(request.into_inner()))
            .await
        {
            Ok(response) => Ok(Response::new(ProductModel::entity_to_response(response))),
            Err(err) => Err(Status::internal(format!("{:?}", err))),
        }
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
    use domain::{
        dtos::create_product_dto::CreateProductDto, entities::product_entity::ProductEntity,
    };
    use mockall::*;

    #[tokio::test]
    async fn perform() {
        let mut get_by_id_use_case = MockGetInventoryByIdUseCase::new();
        let mut create_product_use_case = MockCreateProductUseCase::new();
        let request = Request::<GetByIdRequest>::new(GetByIdRequest {
            ..Default::default()
        });
        get_by_id_use_case
            .expect_perform()
            .returning(|_id| Ok(Some(ProductEntity::default())))
            .times(1);

        let sut = InventoryController::new(
            Arc::new(get_by_id_use_case),
            Arc::new(create_product_use_case),
        );

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
    mock! {
        pub CreateProductUseCase {}
        #[tonic::async_trait]
        impl ICreateProductUseCase for CreateProductUseCase{
            async fn perform(&self, dto: CreateProductDto) -> Result<ProductEntity, Box<dyn std::error::Error>> {
                todo!()
            }
        }
    }
}
