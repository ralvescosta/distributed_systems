use infra::telemetry::telemetry::Telemetry;
use std::sync::Arc;
use tonic::{Request, Response, Status};
use tracing::instrument;
use tracing_futures::Instrument;

use crate::{
    inventory::{
        inventory_server::Inventory, CreateProductRequest, GetByIdRequest, GetByTypeRequest,
        GetProductsRequest, ProductResponse, ProductsResponse, UpdateProductRequest,
    },
    models::product_model::ProductModel,
};
use domain::usecases::{
    i_create_product::ICreateProductUseCase, i_get_product_by_id::IGetProductByIdUseCase,
    i_get_products::IGetProductsUseCase, i_get_products_by_type::IGetProductsByTypeUseCase,
};
#[derive(Debug)]
pub struct ProductController {
    get_product_by_id_use_case: Arc<dyn IGetProductByIdUseCase>,
    create_product_use_case: Arc<dyn ICreateProductUseCase>,
    get_products_by_tag_use_case: Arc<dyn IGetProductsByTypeUseCase>,
    get_products_use_case: Arc<dyn IGetProductsUseCase>,
    telemetry: Arc<Telemetry>,
}

impl ProductController {
    pub fn new(
        get_product_by_id_use_case: Arc<dyn IGetProductByIdUseCase>,
        create_product_use_case: Arc<dyn ICreateProductUseCase>,
        get_products_by_tag_use_case: Arc<dyn IGetProductsByTypeUseCase>,
        get_products_use_case: Arc<dyn IGetProductsUseCase>,
        telemetry: Arc<Telemetry>,
    ) -> ProductController {
        ProductController {
            get_product_by_id_use_case,
            create_product_use_case,
            get_products_by_tag_use_case,
            get_products_use_case,
            telemetry,
        }
    }
}

#[tonic::async_trait]
impl Inventory for ProductController {
    #[instrument(name = "gRPC getProductById")]
    async fn get_product_by_id(
        &self,
        request: Request<GetByIdRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        self.telemetry.grpc_set_span_parent(&request);

        let result = self
            .get_product_by_id_use_case
            .perform(request.into_inner().id)
            .instrument(tracing::Span::current())
            .await;
        match result {
            Ok(Some(product)) => Ok(Response::new(ProductModel::entity_to_response(product))),
            Ok(None) => Err(Status::not_found("Not Found")),
            Err(err) => Err(Status::internal(format!("{:?}", err))),
        }
    }

    #[instrument(name = "gRPC getProductByType")]
    async fn get_products_by_type(
        &self,
        request: Request<GetByTypeRequest>,
    ) -> Result<Response<ProductsResponse>, Status> {
        self.telemetry.grpc_set_span_parent(&request);

        let result = self
            .get_products_by_tag_use_case
            .perform(request.into_inner().product_type)
            .instrument(tracing::Span::current())
            .await;
        match result {
            Ok(product) => {
                if product.len() == 0 {
                    todo!()
                }

                todo!()
            }
            Err(err) => Err(Status::internal(format!("{:?}", err))),
        }
    }

    #[instrument(name = "gRPC getProducts")]
    async fn get_products(
        &self,
        request: Request<GetProductsRequest>,
    ) -> Result<Response<ProductsResponse>, Status> {
        self.telemetry.grpc_set_span_parent(&request);

        let request_data = request.into_inner();

        let result = self
            .get_products_use_case
            .perform(request_data.limit, request_data.offset)
            .instrument(tracing::Span::current())
            .await;
        match result {
            Ok(product) => {
                if product.len() == 0 {
                    todo!()
                }

                todo!()
            }
            Err(err) => Err(Status::internal(format!("{:?}", err))),
        }
    }

    #[instrument(name = "gRPC createProduct")]
    async fn create_product(
        &self,
        request: Request<CreateProductRequest>,
    ) -> Result<Response<ProductResponse>, Status> {
        self.telemetry.grpc_set_span_parent(&request);

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
        // let mut get_by_id_use_case = MockGetProductByIdUseCase::new();
        // let mut create_product_use_case = MockCreateProductUseCase::new();
        // let request = Request::<GetByIdRequest>::new(GetByIdRequest {
        //     ..Default::default()
        // });
        // get_by_id_use_case
        //     .expect_perform()
        //     .returning(|_id| Ok(Some(ProductEntity::default())))
        //     .times(1);

        // let sut = ProductController::new(
        //     Arc::new(get_by_id_use_case),
        //     Arc::new(create_product_use_case),
        // );

        // match sut.get_product_by_id(request).await {
        //     Ok(result) => assert_eq!(result.get_ref().title, String::from("")),
        //     _ => assert!(false),
        // }
    }

    mock! {
        pub GetProductByIdUseCase {}
        #[tonic::async_trait]
        impl IGetProductByIdUseCase for GetProductByIdUseCase {
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
