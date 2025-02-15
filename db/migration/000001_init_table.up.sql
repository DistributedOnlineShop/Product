CREATE TABLE "products" (
  "product_id" VARCHAR(12) PRIMARY KEY NOT NULL,
  "vendor_id" UUID NOT NULL,
  "name" VARCHAR NOT NULL,
  "description" TEXT NOT NULL,
  "price" DECIMAL(10,2) NOT NULL,
  "discount" DECIMAL(10,2) NOT NULL,
  "stock" INT NOT NULL,
  "status" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

CREATE TABLE "product_variants" (
  "pv_id" VARCHAR(12) PRIMARY KEY NOT NULL,
  "product_id" VARCHAR(12) NOT NULL,
  "sku" VARCHAR(50),
  "price" NUMERIC(7,2),
  "stock" INT NOT NULL,
  "attributes" JSON,
  "status" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

CREATE TABLE "product_images" (
  "pi_id" VARCHAR(12) PRIMARY KEY NOT NULL,
  "product_id" VARCHAR(12) NOT NULL,
  "pv_id" VARCHAR(12) NOT NULL,
  "image_url" VARCHAR NOT NULL,
  "position" INT NOT NULL,
  "is_primary" BOOLEAN DEFAULT true,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

CREATE TABLE "inventory_levels" (
  "inventory_id" UUID PRIMARY KEY,
  "product_id" VARCHAR(12) NOT NULL,
  "pv_id" VARCHAR(12),
  "stock" INT NOT NULL,
  "updated_at" TIMESTAMP(0) NOT NULL DEFAULT NOW()
);

CREATE TABLE "inventory_adjustments" (
  "adjustment_id" UUID PRIMARY KEY NOT NULL,
  "product_id" VARCHAR(12) NOT NULL,
  "pv_id" VARCHAR(12),
  "adjustment_type" VARCHAR NOT NULL,
  "quantity" INT NOT NULL,
  "reason" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW()
);

CREATE TABLE "vendors" (
  "vendor_id" UUID PRIMARY KEY NOT NULL,
  "vendor_name" VARCHAR NOT NULL,
  "contact_name" VARCHAR NOT NULL,
  "product_type" VARCHAR[] NOT NULL,
  "email" VARCHAR,
  "phone" VARCHAR(50),
  "status" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

ALTER TABLE "products" ADD FOREIGN KEY ("vendor_id") REFERENCES "vendors" ("vendor_id");

ALTER TABLE "product_variants" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id") ON DELETE CASCADE;

ALTER TABLE "product_images" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id") ON DELETE CASCADE;

ALTER TABLE "inventory_levels" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id") ON DELETE CASCADE;

ALTER TABLE "inventory_levels" ADD FOREIGN KEY ("pv_id") REFERENCES "product_variants" ("pv_id") ON DELETE CASCADE;

ALTER TABLE "inventory_adjustments" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id") ON DELETE CASCADE;

ALTER TABLE "inventory_adjustments" ADD FOREIGN KEY ("pv_id") REFERENCES "product_variants" ("pv_id") ON DELETE CASCADE;

CREATE SEQUENCE product_id_seq START 1;
CREATE SEQUENCE product_variant_id_seq START 1;
CREATE SEQUENCE product_image_id_seq START 1;

CREATE OR REPLACE FUNCTION generate_product_id() 
RETURNS TRIGGER AS $$
BEGIN
  NEW.product_id = 'P' || LPAD(nextval('product_id_seq')::TEXT, 6, '0');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION generate_product_variant_id() 
RETURNS TRIGGER AS $$
BEGIN
  NEW.pv_id = 'PV' || LPAD(nextval('product_variant_id_seq')::TEXT, 6, '0');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION generate_product_image_id()
RETURNS TRIGGER AS $$
BEGIN
  NEW.pi_id = 'PI' || LPAD(nextval('product_image_id_seq')::TEXT, 6, '0');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER generate_product_id_trigger
BEFORE INSERT ON "products"
FOR EACH ROW
EXECUTE FUNCTION generate_product_id();

CREATE TRIGGER generate_product_variant_id_trigger
BEFORE INSERT ON "product_variants"
FOR EACH ROW
EXECUTE FUNCTION generate_product_variant_id();

CREATE TRIGGER generate_product_image_id_trigger
BEFORE INSERT ON "product_images"
FOR EACH ROW
EXECUTE FUNCTION generate_product_image_id();
