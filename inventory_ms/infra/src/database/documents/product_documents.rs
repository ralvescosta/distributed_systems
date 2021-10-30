use domain::entities::product_entity::ProductEntity;
use mongodb::bson::DateTime;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone)]
pub struct ProductDocument {
    pub id: String,
    pub product_category: String,
    pub tag: String,
    pub title: String,
    pub subtitle: String,
    pub authors: Vec<String>,
    pub amount_in_stock: i64,
    pub created_at: DateTime,
    pub updated_at: DateTime,
    pub num_pages: i64,
    pub tags: Vec<String>,
}

impl ProductDocument {
    pub fn to_entity(self) -> ProductEntity {
        ProductEntity {
            id: self.id,
            product_category: self.product_category,
            tag: self.tag,
            title: self.title,
            subtitle: self.subtitle,
            authors: self.authors,
            amount_in_stock: self.amount_in_stock,
            created_at: self.created_at.to_string(),
            updated_at: self.updated_at.to_string(),
            num_pages: self.num_pages,
            tags: self.tags,
        }
    }
}
