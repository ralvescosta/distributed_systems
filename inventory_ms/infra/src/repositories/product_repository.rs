use async_trait::async_trait;
use chrono::Utc;
use mongodb::bson::doc;
use std::error::Error;
use uuid::Uuid;

use application::interfaces::i_product_repository::IProductRepository;
use domain::{dtos::create_product_dto::CreateProductDto, entities::product_entity::ProductEntity};

use crate::database::{connection::DbConnection, documents::product_documents::ProductDocument};

pub struct ProductRepository {
    connection: DbConnection,
}

impl ProductRepository {
    pub fn new(connection: DbConnection) -> ProductRepository {
        ProductRepository { connection }
    }
}

#[async_trait]
impl IProductRepository for ProductRepository {
    async fn get_product_by_id(&self, id: String) -> Result<Option<ProductEntity>, Box<dyn Error>> {
        let collection = self
            .connection
            .get_collection::<ProductDocument>(&self.connection.inventory_collection_name);

        let filter = doc! { "id": id };
        match collection.find_one(filter, None).await? {
            None => Ok(None),
            Some(document) => Ok(Some(document.to_entity())),
        }
    }

    async fn create(&self, dto: CreateProductDto) -> Result<ProductEntity, Box<dyn Error>> {
        let collection = self
            .connection
            .get_collection(&self.connection.inventory_collection_name);

        let guid = Uuid::new_v4().to_hyphenated().to_string();

        let doc = doc! {
            "id": guid,
            "product_category": dto.product_category,
            "tag": dto.tag,
            "title": dto.title,
            "subtitle": dto.subtitle,
            "authors": dto.authors,
            "amount_in_stock": dto.amount_in_stock,
            "num_pages": dto.num_pages,
            "tags": dto.tags,
            "created_at": Utc::now().to_rfc3339(),
            "updated_at": Utc::now().to_rfc3339(),
        };

        let result = collection.insert_one(doc, None).await?;

        Ok(ProductEntity::default())
    }
}

#[cfg(test)]
mod test {
    use super::*;
    use mongodb::Client;

    #[tokio::test]
    async fn should_execute_method_with_connection_error() -> mongodb::error::Result<()> {
        let client = Client::with_uri_str("mongodb://example.com").await?;
        let db_connection = DbConnection {
            client,
            app_name: String::from(""),
            db_name: String::from(""),
            inventory_collection_name: String::from(""),
        };

        let sut = ProductRepository::new(db_connection);

        match sut.get_product_by_id(String::from("some")).await {
            Err(_err) => assert!(true),
            _ => assert!(false),
        }

        Ok(())
    }
}
