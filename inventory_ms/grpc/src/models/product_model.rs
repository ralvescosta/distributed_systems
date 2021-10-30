use domain::{dtos::create_product_dto::CreateProductDto, entities::product_entity::ProductEntity};

use crate::inventory::{CreateProductRequest, ProductResponse};
pub struct ProductModel;

impl ProductModel {
    pub fn create_request_to_dto(request: CreateProductRequest) -> CreateProductDto {
        CreateProductDto {
            authors: request.authors,
            amount_in_stock: request.amount_in_stock,
            num_pages: request.num_pages,
            product_category: String::from("book"),
            subtitle: request.subtitle,
            tag: request.tag,
            tags: request.tags,
            title: request.title,
        }
    }

    pub fn entity_to_response(entity: ProductEntity) -> ProductResponse {
        ProductResponse {
            id: entity.id,
            amount_in_stock: entity.amount_in_stock,
            authors: entity.authors,
            created_at: entity.created_at,
            num_pages: entity.num_pages,
            product_category: entity.product_category,
            subtitle: entity.subtitle,
            tag: entity.tag,
            tags: entity.tags,
            title: entity.title,
            updated_at: entity.updated_at,
        }
    }
}
