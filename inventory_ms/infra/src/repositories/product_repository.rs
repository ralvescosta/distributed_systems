use async_trait::async_trait;
use futures::stream::TryStreamExt;
use mongodb::bson::{doc, DateTime};
use std::error::Error;
use tracing::instrument;
use uuid::Uuid;

use tracing_futures::Instrument;

use application::interfaces::i_product_repository::IProductRepository;
use domain::{dtos::create_product_dto::CreateProductDto, entities::product_entity::ProductEntity};

use crate::database::{connection::DbConnection, documents::product_documents::ProductDocument};

#[derive(Debug)]
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
    #[instrument(name = "MONGO SELECT PRODUCT BY ID")]
    async fn get_product_by_id(&self, id: String) -> Result<Option<ProductEntity>, Box<dyn Error>> {
        let collection = self
            .connection
            .get_collection::<ProductDocument>(&self.connection.inventory_collection_name);

        let filter = doc! { "id": id };
        match collection
            .find_one(filter, None)
            .instrument(tracing::Span::current())
            .await?
        {
            None => Ok(None),
            Some(document) => Ok(Some(document.to_entity())),
        }
    }

    #[instrument(name = "MONGO SELECT PRODUCT BY TYPE")]
    async fn get_products_by_type(
        &self,
        product_type: String,
    ) -> Result<Vec<ProductEntity>, Box<dyn Error>> {
        let collection = self
            .connection
            .get_collection::<ProductDocument>(&self.connection.inventory_collection_name);

        let filter = doc! { "type": product_type };
        let mut cursor = collection
            .find(filter, None)
            .instrument(tracing::Span::current())
            .await?;

        let mut products: Vec<ProductEntity> = vec![];
        while let Some(product) = cursor.try_next().await? {
            products.push(product.to_entity());
        }

        Ok(products)
    }

    #[instrument(name = "MONGO CREATE PRODUCT")]
    async fn create(&self, dto: CreateProductDto) -> Result<ProductEntity, Box<dyn Error>> {
        let collection = self
            .connection
            .get_collection::<ProductDocument>(&self.connection.inventory_collection_name);

        let filter = doc! {
            "title": dto.title.clone(),
            "product_category": dto.product_category.clone()
        };
        match collection
            .find_one(filter, None)
            .instrument(tracing::Span::current())
            .await?
        {
            None => {
                let guid = Uuid::new_v4().to_hyphenated().to_string();
                let document = ProductDocument {
                    id: guid,
                    product_category: dto.product_category,
                    tag: dto.tag,
                    title: dto.title,
                    subtitle: dto.subtitle,
                    authors: dto.authors,
                    amount_in_stock: dto.amount_in_stock,
                    num_pages: dto.num_pages,
                    tags: dto.tags,
                    created_at: DateTime::now(),
                    updated_at: DateTime::now(),
                };

                collection
                    .insert_one(document.clone(), None)
                    .instrument(tracing::Span::current())
                    .await?;

                Ok(document.to_entity())
            }
            Some(document) => Ok(document.to_entity()),
        }
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
