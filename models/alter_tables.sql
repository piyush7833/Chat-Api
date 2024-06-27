-- Adding foreign keys
ALTER TABLE "notificationId"
ADD FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "userRelation"
ADD FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE CASCADE,
ADD FOREIGN KEY ("relatedUserId") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "groups"
ADD FOREIGN KEY ("ownerId") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "groupUsers"
ADD FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE CASCADE,
ADD FOREIGN KEY ("groupId") REFERENCES "groups" ("id") ON DELETE CASCADE;

ALTER TABLE "messages"
ADD FOREIGN KEY ("senderId") REFERENCES "users" ("id") ON DELETE CASCADE,
ADD FOREIGN KEY ("receiverId") REFERENCES "users" ("id") ON DELETE CASCADE,
ADD FOREIGN KEY ("threadId") REFERENCES "messageThreads" ("id") ON DELETE CASCADE;

ALTER TABLE "messageThreads"
ADD FOREIGN KEY ("messageId") REFERENCES "messages" ("id") ON DELETE CASCADE,
ADD FOREIGN KEY ("createdBy") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "reminders"
ADD FOREIGN KEY ("senderId") REFERENCES "users" ("id") ON DELETE CASCADE,
ADD FOREIGN KEY ("receiverId") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "tags"
ADD FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE CASCADE,
ADD FOREIGN KEY ("messageId") REFERENCES "messages" ("id") ON DELETE CASCADE,
ADD FOREIGN KEY ("reminderId") REFERENCES "reminders" ("id") ON DELETE CASCADE;

ALTER TABLE "statuses"
ADD FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "visibility"
ADD FOREIGN KEY ("userId") REFERENCES "users" ("id") ON DELETE CASCADE,
ADD FOREIGN KEY ("friendIds") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "calls"
ADD FOREIGN KEY ("callerId") REFERENCES "users" ("id") ON DELETE CASCADE,
ADD FOREIGN KEY ("receiverId") REFERENCES "users" ("id") ON DELETE CASCADE;
