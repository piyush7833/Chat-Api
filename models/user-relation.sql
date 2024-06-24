CREATE TYPE relation_status AS ENUM ('friends', 'blocked','pending');
CREATE TABLE IF NOT EXISTS "userRelation" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "userId" VARCHAR(255) NOT NULL,
    "relatedUserId" VARCHAR(255) NOT NULL,
    "status" relation_status NOT NULL DEFAULT 'pending',
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id"),
    CONSTRAINT unique_user_related_user UNIQUE ("userId", "relatedUserId")
);
