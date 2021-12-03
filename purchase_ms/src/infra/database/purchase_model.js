const { Schema, model } = require("mongoose");

const PurchaseSchema = new Schema({
  order_id: {
    type: String,
    required: true,
  },
  user_id: {
    type: String,
    required: true,
  },
  product_id: {
    type: String,
    required: true,
  },
  created_at: {
    type: Date,
    default: new Date(),
  },
  deleted_at: Date,
});

const PurchaseModel = model("Purchase", PurchaseSchema);

module.exports = { PurchaseModel };