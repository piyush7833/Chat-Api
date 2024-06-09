CREATE TABLE IF NOT EXISTS "FriendRequest" (
    "id" VARCHAR(255) NOT NULL DEFAULT uuid_generate_v4(),
    "senderId" VARCHAR(255),
    "receiverId" VARCHAR(255),
    "createdAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updatedAt" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT "unique_sender_receiver_pair" UNIQUE ("senderId", "receiverId"),
    PRIMARY KEY ("id")
);