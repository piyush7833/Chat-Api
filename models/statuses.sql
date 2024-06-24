CREATE TYPE status_type AS ENUM ('video', 'image','text');
CREATE TABLE IF NOT EXISTS "statuses" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "userId" VARCHAR(255) NOT NULL,
    "content" TEXT,
    "caption" TEXT,
    "type"  status_type,
    "expiration" TIMESTAMPTZ,
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);