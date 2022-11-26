export const up = async (db, client) => {
    const session = client.startSession();
    try {
        await session.withTransaction(async () => {
            await db.createCollection("users", {
                validator: {
                    $jsonSchema: {
                        bsonType: "object",
                        title: "User Object Validation",
                        required: ["user_id", "email", "name", "access_token"],
                        properties: {
                            user_id: {
                                bsonType: "string",
                                uniqueItems: true,
                                description:
                                    "'user_id' must be a string and is required",
                            },
                            email: {
                                bsonType: "string",
                                pattern:
                                    "^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$",
                                description:
                                    "'email' must be a string and is required",
                            },
                            name: {
                                bsonType: "string",
                                description:
                                    "'name' must be a string and is required",
                            },
                            access_token: {
                                bsonType: "string",
                                description:
                                    "'access_token' must be a string and is required",
                            },
                        },
                    },
                },
            });
        });
    } finally {
        await session.endSession();
    }
};

export const down = async (db, client) => {
    const session = client.startSession();
    try {
        await session.withTransaction(async () => {
            await db.collection("users").drop();
        });
    } finally {
        await session.endSession();
    }
};
