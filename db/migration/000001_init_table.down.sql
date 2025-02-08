ALTER TABLE "product_variants" DROP CONSTRAINT IF EXISTS product_variants_product_id_fkey;
ALTER TABLE "product_images" DROP CONSTRAINT IF EXISTS product_images_product_id_fkey;
ALTER TABLE "inventory_levels" DROP CONSTRAINT IF EXISTS inventory_levels_product_id_fkey;
ALTER TABLE "inventory_levels" DROP CONSTRAINT IF EXISTS inventory_levels_pv_id_fkey;
ALTER TABLE "inventory_adjustments" DROP CONSTRAINT IF EXISTS inventory_adjustments_product_id_fkey;
ALTER TABLE "inventory_adjustments" DROP CONSTRAINT IF EXISTS inventory_adjustments_pv_id_fkey;

DROP TRIGGER IF EXISTS generate_product_id_trigger ON products;
DROP TRIGGER IF EXISTS generate_product_variant_id_trigger ON product_variants;
DROP TRIGGER IF EXISTS generate_product_image_id_trigger ON product_images;

DROP FUNCTION IF EXISTS generate_product_id();
DROP FUNCTION IF EXISTS generate_product_variant_id();
DROP FUNCTION IF EXISTS generate_product_image_id();

DROP SEQUENCE IF EXISTS product_id_seq;
DROP SEQUENCE IF EXISTS product_variant_id_seq;
DROP SEQUENCE IF EXISTS product_image_id_seq;

DROP TABLE IF EXISTS "products";
DROP TABLE IF EXISTS "product_variants";
DROP TABLE IF EXISTS "product_images";
DROP TABLE IF EXISTS "inventory_levels";
DROP TABLE IF EXISTS "inventory_adjustments";
DROP TABLE IF EXISTS "vendors";