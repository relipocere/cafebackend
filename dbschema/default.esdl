module default {
    scalar type Affordability extending enum<Cheap, Affordable, Expensive>;
    scalar type Cuisine extending enum<American, Asian, European>;

    type Store{
        required property title -> str;
        required property affordability -> Affordability;
        required property cuisine_type -> Cuisine;
        required link owner -> User{
            on target delete delete source;
        }
        multi link reviews -> StoreReview{
            constraint exclusive;
            on source delete delete target;
        }
        required property image_id -> str;
        required property created_at -> datetime;
        required property updated_at -> datetime;

        index on (.title);
        index on (.affordability);
        index on (.cuisine_type);
    }

    type StoreReview{
        required property content -> str;
        required link author -> User{
            on target delete delete source;
        };
    }

    scalar type UserKind extending enum<Consumer, Business>;
    type User{
        required property username -> str{
            constraint exclusive;
        }
        required property full_name -> str;
        required property kind -> UserKind;
        required property password_hash -> str;
        required property salt -> str;
        required property created_at -> datetime;
        required property updated_at -> datetime;
        link session -> AuthSession{
            constraint exclusive;
            on source delete delete target;
        }

        index on (.username);
    }

    type AuthSession{
        required property session_id -> str{
            constraint exclusive;
        };
        required property expires_at -> datetime;
    }

    type Product{
        required property name -> str;
        required link store -> Store{
            on target delete delete source;
        }
        property ingredients -> array<str>;
        required property calories -> int64;
        required property image_id -> str;
        required property created_at -> datetime;
        required property updated_at -> datetime;
    }
}
