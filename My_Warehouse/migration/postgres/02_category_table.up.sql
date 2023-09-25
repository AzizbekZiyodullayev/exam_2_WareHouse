
CREATE TABLE "category" (
    "id" UUID PRIMARY KEY,
    "title" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);