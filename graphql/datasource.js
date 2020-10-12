import { RESTDataSource } from "apollo-datasource-rest";

export class ShortenAPI extends RESTDataSource {
  constructor() {
    super();
    this.baseURL = "http://localhost:8080/";
  }

  async shorten(url) {
    console.dir({url})
    return this.post(
      `shorten`,
      {url}
    );
  }
}
