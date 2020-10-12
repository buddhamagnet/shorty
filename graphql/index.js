import { ApolloServer, gql } from "apollo-server";
import { ShortenAPI } from "./datasource";

const typeDefs = gql`
  type URL {
    url: String
    id: String
  }
  type Query {
    check: URL
  }
  type Mutation {
    shorten(url: String!): URL!
  }
`;

const resolvers = {
  Mutation: {
    shorten: async(_root, args, {dataSources}) => {
      const data = await dataSources.shortenAPI.shorten(args.url);
      return data;
    }
  }
};

const server = new ApolloServer({
  typeDefs,
  resolvers,
  dataSources: () => ({
    shortenAPI: new ShortenAPI(),
  }),
});

server.listen().then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});
