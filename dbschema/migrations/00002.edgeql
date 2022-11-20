CREATE MIGRATION m1vdrdgdbetqxpbbymcn3zfz6wan4wueg7wqsdcknk76nntavqfy4a
    ONTO m1b67kksepmhtqso6hoznfro4oq53g4ze7jsr3w6reyeo7qmlbxkha
{
  ALTER TYPE default::Product {
      ALTER LINK store {
          ON TARGET DELETE DELETE SOURCE;
      };
  };
};
