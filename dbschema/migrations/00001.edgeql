CREATE MIGRATION m1ofebno3ebi3hscb5hhqfztp6l7f4t7qowna4ffjxsnubz2dleupa
    ONTO initial
{
  CREATE SCALAR TYPE default::UserKind EXTENDING enum<Consumer, Business>;
  CREATE TYPE default::User {
      CREATE REQUIRED PROPERTY username -> std::str {
          CREATE CONSTRAINT std::exclusive;
      };
      CREATE INDEX ON (.username);
      CREATE REQUIRED PROPERTY created_at -> std::datetime;
      CREATE REQUIRED PROPERTY full_name -> std::str;
      CREATE REQUIRED PROPERTY kind -> default::UserKind;
      CREATE REQUIRED PROPERTY password_hash -> std::str;
      CREATE REQUIRED PROPERTY salt -> std::str;
      CREATE REQUIRED PROPERTY updated_at -> std::datetime;
  };
  CREATE TYPE default::StoreReview {
      CREATE REQUIRED LINK author -> default::User {
          ON TARGET DELETE DELETE SOURCE;
      };
      CREATE REQUIRED PROPERTY content -> std::str;
  };
  CREATE SCALAR TYPE default::Affordability EXTENDING enum<Cheap, Affordable, Expensive>;
  CREATE SCALAR TYPE default::Cuisine EXTENDING enum<American, Asian, European>;
  CREATE TYPE default::Store {
      CREATE REQUIRED PROPERTY cuisine_type -> default::Cuisine;
      CREATE INDEX ON (.cuisine_type);
      CREATE REQUIRED PROPERTY affordability -> default::Affordability;
      CREATE INDEX ON (.affordability);
      CREATE REQUIRED PROPERTY title -> std::str;
      CREATE INDEX ON (.title);
      CREATE REQUIRED LINK owner -> default::User {
          ON TARGET DELETE DELETE SOURCE;
      };
      CREATE MULTI LINK reviews -> default::StoreReview {
          ON SOURCE DELETE DELETE TARGET;
          CREATE CONSTRAINT std::exclusive;
      };
      CREATE REQUIRED PROPERTY created_at -> std::datetime;
      CREATE REQUIRED PROPERTY image_id -> std::str;
      CREATE REQUIRED PROPERTY updated_at -> std::datetime;
  };
  CREATE TYPE default::Product {
      CREATE REQUIRED LINK store -> default::Store;
      CREATE REQUIRED PROPERTY calories -> std::int64;
      CREATE REQUIRED PROPERTY created_at -> std::datetime;
      CREATE REQUIRED PROPERTY image_id -> std::str;
      CREATE PROPERTY ingredients -> array<std::str>;
      CREATE REQUIRED PROPERTY name -> std::str;
      CREATE REQUIRED PROPERTY updated_at -> std::datetime;
  };
};
